package downloader

import (
	"fmt"
	"math/big"

	"github.com/openether/ethcore/common"
	"github.com/openether/ethcore/core/types"
)

// headerCheckFn is a callback type for verifying a header's presence in the local chain.
type headerCheckFn func(common.Hash) bool

// blockAndStateCheckFn is a callback type for verifying block and associated states' presence in the local chain.
type blockAndStateCheckFn func(common.Hash) bool

// headerRetrievalFn is a callback type for retrieving a header from the local chain.
type headerRetrievalFn func(common.Hash) *types.Header

// blockRetrievalFn is a callback type for retrieving a block from the local chain.
type blockRetrievalFn func(common.Hash) *types.Block

// headHeaderRetrievalFn is a callback type for retrieving the head header from the local chain.
type headHeaderRetrievalFn func() *types.Header

// headBlockRetrievalFn is a callback type for retrieving the head block from the local chain.
type headBlockRetrievalFn func() *types.Block

// headFastBlockRetrievalFn is a callback type for retrieving the head fast block from the local chain.
type headFastBlockRetrievalFn func() *types.Block

// headBlockCommitterFn is a callback for directly committing the head block to a certain entity.
type headBlockCommitterFn func(common.Hash) error

// tdRetrievalFn is a callback type for retrieving the total difficulty of a local block.
type tdRetrievalFn func(common.Hash) *big.Int

// headerChainInsertFn is a callback type to insert a batch of headers into the local chain.
type headerChainInsertFn func([]*types.Header, int) (int, error)

// blockChainInsertFn is a callback type to insert a batch of blocks into the local chain.
type blockChainInsertFn func(types.Blocks) (int, error)

// receiptChainInsertFn is a callback type to insert a batch of receipts into the local chain.
type receiptChainInsertFn func(types.Blocks, []types.Receipts) (int, error)

// chainRollbackFn is a callback type to remove a few recently added elements from the local chain.
type chainRollbackFn func([]common.Hash)

// peerDropFn is a callback type for dropping a peer detected as malicious.
type peerDropFn func(id string)

// dataPack is a data message returned by a peer for some query.
type dataPack interface {
	PeerId() string
	Items() int
	Stats() string
}

// headerPack is a batch of block headers returned by a peer.
type headerPack struct {
	peerId  string
	headers []*types.Header
}

func (p *headerPack) PeerId() string { return p.peerId }
func (p *headerPack) Items() int     { return len(p.headers) }
func (p *headerPack) Stats() string  { return fmt.Sprintf("%d", len(p.headers)) }

// bodyPack is a batch of block bodies returned by a peer.
type bodyPack struct {
	peerId       string
	transactions [][]*types.Transaction
	uncles       [][]*types.Header
}

func (p *bodyPack) PeerId() string { return p.peerId }
func (p *bodyPack) Items() int {
	if len(p.transactions) <= len(p.uncles) {
		return len(p.transactions)
	}
	return len(p.uncles)
}
func (p *bodyPack) Stats() string { return fmt.Sprintf("%d:%d", len(p.transactions), len(p.uncles)) }

// receiptPack is a batch of receipts returned by a peer.
type receiptPack struct {
	peerId   string
	receipts [][]*types.Receipt
}

func (p *receiptPack) PeerId() string { return p.peerId }
func (p *receiptPack) Items() int     { return len(p.receipts) }
func (p *receiptPack) Stats() string  { return fmt.Sprintf("%d", len(p.receipts)) }

// statePack is a batch of states returned by a peer.
type statePack struct {
	peerId string
	states [][]byte
}

func (p *statePack) PeerId() string { return p.peerId }
func (p *statePack) Items() int     { return len(p.states) }
func (p *statePack) Stats() string  { return fmt.Sprintf("%d", len(p.states)) }
