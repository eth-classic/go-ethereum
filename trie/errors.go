package trie

import (
	"fmt"

	"github.com/ether-core/go-ethereum/common"
)

// MissingNodeError is returned by the trie functions (TryGet, TryUpdate, TryDelete)
// in the case where a trie node is not present in the local database. Contains
// information necessary for retrieving the missing node through an ODR service.
//
// NodeHash is the hash of the missing node
//
// RootHash is the original root of the trie that contains the node
//
// Key is a binary-encoded key that contains the prefix that leads to the first
// missing node and optionally a suffix that hints on which further nodes should
// also be retrieved
//
// PrefixLen is the nibble length of the key prefix that leads from the root to
// the missing node
//
// SuffixLen is the nibble length of the remaining part of the key that hints on
// which further nodes should also be retrieved (can be zero when there are no
// such hints in the error message)

// MissingNodeError is returned by the trie functions (TryGet, TryUpdate, TryDelete)
// in the case where a trie node is not present in the local database. It contains
// information necessary for retrieving the missing node.
type MissingNodeError struct {
	NodeHash common.Hash // hash of the missing node
	Path     []byte      // hex-encoded path to the missing node
}

func (err *MissingNodeError) Error() string {
	return fmt.Sprintf("missing trie node %x (path %x)", err.NodeHash, err.Path)
}
