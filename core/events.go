package core

import (
	"math/big"
	"time"

	"github.com/openether/ethcore/common"
	"github.com/openether/ethcore/core/types"
	"github.com/openether/ethcore/core/vm"
)

// TxPreEvent is posted when a transaction enters the transaction pool.
type TxPreEvent struct{ Tx *types.Transaction }

// TxPostEvent is posted when a transaction has been processed.
type TxPostEvent struct{ Tx *types.Transaction }

// PendingLogsEvent is posted pre mining and notifies of pending logs.
type PendingLogsEvent struct {
	Logs vm.Logs
}

// PendingStateEvent is posted pre mining and notifies of pending state changes.
type PendingStateEvent struct{}

// NewBlockEvent is posted when a block has been imported.
type NewBlockEvent struct{ Block *types.Block }

// NewMinedBlockEvent is posted when a block has been imported.
type NewMinedBlockEvent struct{ Block *types.Block }

// RemovedTransactionEvent is posted when a reorg happens
type RemovedTransactionEvent struct{ Txs types.Transactions }

// RemovedLogEvent is posted when a reorg happens
type RemovedLogsEvent struct{ Logs vm.Logs }

// ChainSplit is posted when a new head is detected
type ChainSplitEvent struct {
	Block *types.Block
	Logs  vm.Logs
}

type ChainEvent struct {
	Block *types.Block
	Hash  common.Hash
	Logs  vm.Logs
}

type ChainSideEvent struct {
	Block *types.Block
	Logs  vm.Logs
}

// TODO: no usages found in project files
type PendingBlockEvent struct {
	Block *types.Block
	Logs  vm.Logs
}

type ChainInsertEvent struct {
	Processed       int
	Queued          int
	Ignored         int
	TxCount         int
	LastNumber      uint64
	LastHash        common.Hash
	Elasped         time.Duration
	LatestBlockTime time.Time
}

type ReceiptChainInsertEvent struct {
	Processed         int
	Ignored           int
	FirstNumber       uint64
	FirstHash         common.Hash
	LastNumber        uint64
	LastHash          common.Hash
	Elasped           time.Duration
	LatestReceiptTime time.Time
}

type HeaderChainInsertEvent struct {
	Processed  int
	Ignored    int
	LastNumber uint64
	LastHash   common.Hash
	Elasped    time.Duration
}

// TODO: no usages found in project files
type ChainUncleEvent struct {
	Block *types.Block
}

type ChainHeadEvent struct{ Block *types.Block }

type GasPriceChanged struct{ Price *big.Int }

// Mining operation events
type StartMining struct{}
type StopMining struct{}
