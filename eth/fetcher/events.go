package fetcher

import (
	"github.com/ether-core/go-ethereum/core/types"
)

type FetcherInsertBlockEvent struct {
	Peer  string
	Block *types.Block
}
