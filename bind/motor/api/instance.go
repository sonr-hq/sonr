package api

import (
	"context"

	mt "github.com/sonr-hq/sonr/bind/motor/types/v1"
	"github.com/sonr-hq/sonr/pkg/node"
)

type MotorInstance struct {
	// Node is the libp2p host
	Node *node.Motor
	ctx  context.Context
	//
}

func NewMotorInstance(ctx context.Context) *MotorInstance {
	return &MotorInstance{
		ctx: ctx,
	}
}

func (mi *MotorInstance) Connect(req *mt.ConnectRequest, options ...node.NodeOption) (*mt.ConnectResponse, error) {
	n, err := node.New(mi.ctx, options...)
	if err != nil {
		return nil, err
	}
	mi.Node = n
	return &mt.ConnectResponse{}, nil
}
