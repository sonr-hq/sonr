package mpc

import (
	"errors"
	"sync"

	peer "github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

// Keygen Generates a new ECDSA private key shared among all the given participants.
func Keygen(current party.ID, threshold int, net common.Network, addrPrefix string) ([]common.WalletShare, error) {
	var mtx sync.Mutex
	configs := make(map[party.ID]*cmp.Config)
	var wg sync.WaitGroup
	for _, id := range net.Ls() {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			conf, err := CmpKeygen(id, net.Ls(), net, threshold, &wg, pl)
			if err != nil {
				return
			}
			mtx.Lock()
			configs[conf.ID] = conf
			mtx.Unlock()
		}(id)
	}
	wg.Wait()
	// conf := <-doneChan
	shares := make([]common.WalletShare, 0)
	for _, conf := range configs {
		shares = append(shares, NewWalletShare(addrPrefix, conf))
	}
	return shares, nil
}

//
// CMP methods
//

// It creates a new handler for the keygen protocol, runs the handler loop, and returns the result
func CmpKeygen(id party.ID, ids party.IDSlice, n common.Network, threshold int, wg *sync.WaitGroup, pl *pool.Pool) (*cmp.Config, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Keygen(curve.Secp256k1{}, id, ids, threshold, pl), nil)
	if err != nil {
		return nil, err
	}

	handlerLoop(id, h, n)
	r, err := h.Result()
	if err != nil {
		return nil, err
	}
	conf := r.(*cmp.Config)
	return conf, nil
}

// It creates a new handler for the refresh protocol, runs the handler loop, and returns the result
func CmpRefresh(c *cmp.Config, n common.Network, wg *sync.WaitGroup, pl *pool.Pool) (*cmp.Config, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Refresh(c, pl), nil)
	if err != nil {
		return nil, err
	}

	handlerLoop(c.ID, h, n)
	r, err := h.Result()
	if err != nil {
		return nil, err
	}
	conf := r.(*cmp.Config)
	return conf, nil
}

// It creates a new `protocol.MultiHandler` for the `cmp.Sign` protocol, and then runs the handler loop
func CmpSign(c *cmp.Config, m []byte, signers party.IDSlice, n common.Network, wg *sync.WaitGroup, pl *pool.Pool) (*ecdsa.Signature, error) {
	defer wg.Done()
	h, err := protocol.NewMultiHandler(cmp.Sign(c, signers, m, pl), nil)
	if err != nil {
		return nil, err
	}
	handlerLoop(c.ID, h, n)

	r, err := h.Result()
	if err != nil {
		return nil, err
	}
	sig := r.(*ecdsa.Signature)
	if !sig.Verify(c.PublicPoint(), m) {
		return nil, errors.New("failed to verify cmp signature")
	}
	return sig, nil
}

// handlerLoop is a helper function that loops over all the parties and calls the given handler.
func handlerLoop(id party.ID, h protocol.Handler, network common.Network) {
	for {
		select {

		// outgoing messages
		case msg, ok := <-h.Listen():
			if !ok {
				<-network.Done(id)
				// the channel was closed, indicating that the protocol is done executing.
				return
			}
			go network.Send(msg)

			// incoming messages
		case msg := <-network.Next(id):
			h.Accept(msg)
		}
	}
}

//
// Helper Functions
//

// It converts a peer.ID to a party.ID
func peerIdToPartyId(id peer.ID) party.ID {
	return party.ID(id)
}

// It converts a party ID to a peer ID
func partyIdToPeerId(id party.ID) peer.ID {
	return peer.ID(id)
}

// It converts a list of peer IDs to a list of party IDs
func peerIdListToPartyIdList(ids []peer.ID) []party.ID {
	partyIds := make([]party.ID, len(ids))
	for i, id := range ids {
		partyIds[i] = peerIdToPartyId(id)
	}
	return partyIds
}

// It converts a list of party IDs to a list of peer IDs
func partyIdListToPeerIdList(ids []party.ID) []peer.ID {
	peerIds := make([]peer.ID, len(ids))
	for i, id := range ids {
		peerIds[i] = partyIdToPeerId(id)
	}
	return peerIds
}
