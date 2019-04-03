package fetcher

import (
	"github.com/openether/ethcore/core/types"
)

type FetcherInsertBlockEvent struct {
	Peer  string
	Block *types.Block
}
