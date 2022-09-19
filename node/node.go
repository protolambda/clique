package node

import (
	"context"
	"github.com/ethereum/go-ethereum/core/beacon"
	"github.com/ethereum/go-ethereum/core/types"
)

// ExecutionPayload is a b
// TODO add difficulty field
type ExecutionPayload = beacon.ExecutableDataV1

type Engine interface {
	GetPayload(ctx context.Context, payloadId beacon.PayloadID) (*ExecutionPayload, error)
	ForkchoiceUpdate(ctx context.Context, state *beacon.ForkchoiceStateV1, attr *beacon.PayloadAttributesV1) (*beacon.ForkChoiceResponse, error)
	NewPayload(ctx context.Context, payload *ExecutionPayload) (*beacon.PayloadStatusV1, error)
}

type P2P interface {
	Dial()
}

type Metrics interface {
}

// TODO: clique api, requires header reader

type CliqueNode struct {
	eng     Engine
	p2p     P2P
	metrics Metrics

	blocks chan *types.Block
	quit   chan chan error
}

func (cn *CliqueNode) Start() error {
	for {
		select {
		case block := <-cn.blocks:
			// TODO insert payload

		case quitCh := <-cn.quit:
			// TOOD close rpc
			// TODO close peer connections
		}
	}
}

func (cn *CliqueNode) Close() error {
	ch := make(chan error)
	cn.quit <- ch
	err := <-ch
	return err
}
