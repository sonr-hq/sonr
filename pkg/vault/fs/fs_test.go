package fs

import (
	"context"
	"testing"

	"github.com/sonr-hq/sonr/pkg/common"
	ipfs "github.com/sonr-hq/sonr/pkg/node/ipfs/local"
)

func TestNew(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	hw, err := ipfs.New(ctx)
	if err != nil {
		t.Fatal(err)
	}
	v1, err := New(ctx, "test", hw)
	if err != nil {
		t.Fatal(err)
	}
	v2, err := New(ctx, "test2", hw, WithIPFSPath(v1.CID()))
	if err != nil {
		t.Fatal(err)
	}
	if v1.CID() != v2.CID() {
		t.Fatal("expected same CID")
	}
}

func TestStoreShare(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	hw, err := ipfs.New(ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := New(ctx, "test", hw)
	if err != nil {
		t.Fatal(err)
	}
	password := "abcdef123456"
	testShare := &common.WalletShareConfig{
		SelfId: "test",
	}
	bz, err := testShare.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	if err := v.StoreShare(bz, "test", password); err != nil {
		t.Fatal(err)
	}

	v2, err := New(ctx, "test2", hw, WithIPFSPath(v.CID()))
	if err != nil {
		t.Fatal(err)
	}
	shares, err := v2.LoadShares(password)
	if err != nil {
		t.Fatal(err)
	}
	if len(shares) != 1 {
		t.Fatal("expected 1 share")
	}
	if shares[0].SelfId != "test" {
		t.Fatal("expected test")
	}
}
