package reqresp

import (
	"bytes"
	"context"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"io"
)

type NewStreamFn func(ctx context.Context, peerId peer.ID, protocolId ...protocol.ID) (network.Stream, error)

func (newStreamFn NewStreamFn) Request(ctx context.Context, peerId peer.ID, protocolId protocol.ID, r io.Reader, comp Compression, handle ResponseHandler) error {
	stream, err := newStreamFn(ctx, peerId, protocolId)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := EncodePayload(r, &buf, comp); err != nil {
		return err
	}
	if _, err := stream.Write(buf.Bytes()); err != nil {
		return err
	}
	return handle(ctx, stream, stream)
}