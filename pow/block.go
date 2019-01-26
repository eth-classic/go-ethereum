package pow

import (
	"math/big"

	"github.com/ether-core/go-ethereum/common"
	"github.com/ether-core/go-ethereum/core/types"
)

type Block interface {
	Difficulty() *big.Int
	HashNoNonce() common.Hash
	Nonce() uint64
	MixDigest() common.Hash
	NumberU64() uint64
}

type ChainManager interface {
	GetBlockByNumber(uint64) *types.Block
	CurrentBlock() *types.Block
}
