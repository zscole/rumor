package methods

import (
	"encoding/hex"
	"fmt"
	"github.com/protolambda/rumor/rpc/reqresp"
)

// instead of parsing the whole body, we can just leave it as bytes.
type BeaconBlockBodyRaw []byte

func (b *BeaconBlockBodyRaw) Limit() uint64 {
	// just cap block body size at 1 MB
	return 1 << 20
}

type BeaconBlock struct {
	Slot Slot
	ParentRoot Root
	StateRoot Root
	Body BeaconBlockBodyRaw
}

type SignedBeaconBlock struct {
	Message BeaconBlock
	Signature BLSSignature
}


type BlocksByRangeReqV1 struct {
	HeadBlockRoot Root // TO be removed in BlocksByRange v2
	StartSlot Slot
	Count uint64
	Step uint64
}

func (r *BlocksByRangeReqV1) String() string {
	return fmt.Sprintf("%v", *r)
}

var BlocksByRangeRPCv1 = reqresp.RPCMethod{
	Protocol:      "/eth2/beacon_chain/req/beacon_blocks_by_range/1/ssz",
	MaxChunkCount: 100, // 100 blocks default maximum
	RequestCodec: reqresp.NewSSZCodec((*BlocksByRangeReqV1)(nil)),
	ResponseChunkCodec: reqresp.NewSSZCodec((*SignedBeaconBlock)(nil)),
}

type BlocksByRangeReqV2 struct {
	StartSlot Slot
	Count uint64
	Step uint64
}

func (r *BlocksByRangeReqV2) String() string {
	return fmt.Sprintf("%v", *r)
}

var BlocksByRangeRPCv2 = reqresp.RPCMethod{
	Protocol:      "/eth2/beacon_chain/req/beacon_blocks_by_range/2/ssz",
	MaxChunkCount: 100, // 100 blocks default maximum
	RequestCodec: reqresp.NewSSZCodec((*BlocksByRangeReqV2)(nil)),
	ResponseChunkCodec: reqresp.NewSSZCodec((*SignedBeaconBlock)(nil)),
}

type BlocksByRootReq []Root

func (*BlocksByRootReq) Limit() uint64 {
	return 100
}

func (r BlocksByRootReq) String() string {
	if len(r) == 0 {
		return "empty blocks-by-root request"
	}
	out := make([]byte, 0, len(r) * 66)
	for i, root := range r {
		hex.Encode(out[i*66:], root[:])
		out[(i+1)*66-2] = ','
		out[(i+1)*66-1] = ' '
	}
	return "blocks-by-root requested: " + string(out[:len(out) - 1])
}

var BlocksByRootRPCv1 = reqresp.RPCMethod{
	Protocol:      "/eth2/beacon_chain/req/beacon_blocks_by_root/1/ssz",
	MaxChunkCount: 100, // 100 blocks default maximum
	RequestCodec: reqresp.NewSSZCodec((*BlocksByRootReq)(nil)),
	ResponseChunkCodec: reqresp.NewSSZCodec((*SignedBeaconBlock)(nil)),
}
