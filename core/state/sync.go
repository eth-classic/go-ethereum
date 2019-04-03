package state

import (
	"bytes"
	"math/big"

	"github.com/openether/ethcore/common"
	"github.com/openether/ethcore/ethdb"
	"github.com/openether/ethcore/rlp"
	"github.com/openether/ethcore/trie"
)

// NewStateSync create a new state trie download scheduler.
func NewStateSync(root common.Hash, database ethdb.Database) *trie.Sync {
	var syncer *trie.Sync

	callback := func(leaf []byte, parent common.Hash) error {
		var obj struct {
			Nonce    uint64
			Balance  *big.Int
			Root     common.Hash
			CodeHash []byte
		}
		if err := rlp.Decode(bytes.NewReader(leaf), &obj); err != nil {
			return err
		}
		syncer.AddSubTrie(obj.Root, 64, parent, nil)
		syncer.AddRawEntry(common.BytesToHash(obj.CodeHash), 64, parent)

		return nil
	}
	syncer = trie.NewTrieSync(root, database, callback)
	return syncer
}
