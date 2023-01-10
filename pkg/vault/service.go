package vault

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/node/ipfs"

	"github.com/sonr-hq/sonr/pkg/vault/session"
	v1 "github.com/sonr-hq/sonr/third_party/types/highway/vault/v1"
	"github.com/sonr-hq/sonr/x/identity/types"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

// Default Variables
var (
	defaultRpOrigins = []string{
		"https://auth.sonr.io",
		"https://sonr.id",
		"https://sandbox.sonr.network",
		"localhost:3000",
	}
)

// `VaultService` is a type that implements the `v1.VaultServer` interface, and has a field called
// `highway` of type `*HighwayNode`.
// @property  - `v1.VaultServer`: This is the interface that the Vault service implements.
// @property highway - This is the HighwayNode that the VaultService is running on.
type VaultService struct {
	highway ipfs.IPFS
	rpName  string
	rpIcon  string
	cache   *gocache.Cache
}

// It creates a new VaultService and registers it with the gRPC server
func NewVaultService(ctx context.Context, mux *runtime.ServeMux, hway ipfs.IPFS, cache *gocache.Cache) (*VaultService, error) {
	srv := &VaultService{
		cache:   cache,
		highway: hway,
		// TODO: Make all Webauthn options configurable through cmd line flags
		rpName: "Sonr",
		rpIcon: "https://raw.githubusercontent.com/sonr-hq/sonr/master/docs/static/favicon.png",
	}
	err := v1.RegisterVaultHandlerServer(ctx, mux, srv)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

// Challeng returns a random challenge for the client to sign.
func (v *VaultService) Challenge(ctx context.Context, req *v1.ChallengeRequest) (*v1.ChallengeResponse, error) {
	entry, err := session.NewEntry(req.RpId)
	if err != nil {
		return nil, err
	}
	optsJson, err := entry.BeginRegistration()
	if err != nil {
		return nil, err
	}
	v.cache.Set(entry.ID, entry, -1)
	return &v1.ChallengeResponse{
		RpName:          v.rpName,
		CreationOptions: optsJson,
		SessionId:       entry.ID,
		RpIcon:          v.rpIcon,
	}, nil
}

// Register registers a new keypair and returns the public key.
func (v *VaultService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	// Get Session
	entry, err := session.GetEntry(req.SessionId, v.cache)
	if err != nil {
		return nil, err
	}
	cred, err := entry.FinishRegistration(req.CredentialResponse)
	if err != nil {
		return nil, err
	}
	vm, err := types.NewWebAuthnVM(cred)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create new verification method: %s", err))
	}
	return &v1.RegisterResponse{
		Success:            true,
		VerificationMethod: vm,
	}, nil

}

// Keygen generates a new keypair and returns the public key.
func (v *VaultService) Keygen(ctx context.Context, req *v1.KeygenRequest) (*v1.KeygenResponse, error) {
	// Create a new offline wallet
	wallet, err := NewWallet(ctx, req.Prefix, v.highway)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create new offline wallet using MPC: %s", err))
	}

	// Return Configuration Response
	return &v1.KeygenResponse{
		Id:      []byte(uuid.New().String()),
		Address: wallet.Address(),
		//		VaultCid:    cid,
		ShareConfig: wallet.Find("current").Share(),
	}, nil
}

// Refresh refreshes the keypair and returns the public key.
func (v *VaultService) Refresh(ctx context.Context, req *v1.RefreshRequest) (*v1.RefreshResponse, error) {
	self, wallet, err := v.assembleWalletFromShares(req.VaultCid, req.ShareConfig)
	if err != nil {
		return nil, err
	}

	newWallet, err := wallet.Refresh(self)
	if err != nil {
		return nil, err
	}
	return &v1.RefreshResponse{
		Id:      []byte(uuid.New().String()),
		Address: newWallet.Address(),
	}, nil
}

// Sign signs the data with the private key and returns the signature.
func (v *VaultService) Sign(ctx context.Context, req *v1.SignRequest) (*v1.SignResponse, error) {
	return nil, errors.New("Method is unimplemented")
}

// Derive derives a new key from the private key and returns the public key.
func (v *VaultService) Derive(ctx context.Context, req *v1.DeriveRequest) (*v1.DeriveResponse, error) {
	return nil, errors.New("Method is unimplemented")
}

//
// Helper functions
//

// assembleWalletFromShares takes a WalletShareConfig and CID to return a Offline Wallet
func (v *VaultService) assembleWalletFromShares(cid string, current *common.WalletShareConfig) (party.ID, common.Wallet, error) {
	// Initialize provided share
	shares := make([]*common.WalletShareConfig, 0)
	shares = append(shares, current)

	// Fetch Vault share from IPFS
	oldbz, err := v.highway.Get(cid)
	if err != nil {
		return "", nil, err
	}

	// Unmarshal share
	share := &common.WalletShareConfig{}
	err = share.Unmarshal(oldbz)
	if err != nil {
		return "", nil, err
	}

	// Load wallet
	wallet, err := LoadOfflineWallet(shares)
	if err != nil {
		return "", nil, err
	}
	return party.ID(current.SelfId), wallet, nil
}