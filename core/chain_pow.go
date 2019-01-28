package core

import (
	"runtime"

	"github.com/openether/ethcore/core/types"
	"github.com/openether/ethcore/pow"
)

// nonceCheckResult contains the result of a nonce verification.
type nonceCheckResult struct {
	index int  // Index of the item verified from an input array
	valid bool // Result of the nonce verification
}

// verifyNoncesFromHeaders starts a concurrent header nonce verification,
// returning a quit channel to abort the operations and a results channel
// to retrieve the async verifications.
func verifyNoncesFromHeaders(checker pow.PoW, headers []*types.Header) (chan<- struct{}, <-chan nonceCheckResult) {
	items := make([]pow.Block, len(headers))
	for i, header := range headers {
		items[i] = types.NewBlockWithHeader(header)
	}
	return verifyNonces(checker, items)
}

// verifyNoncesFromBlocks starts a concurrent block nonce verification,
// returning a quit channel to abort the operations and a results channel
// to retrieve the async verifications.
func verifyNoncesFromBlocks(checker pow.PoW, blocks []*types.Block) (chan<- struct{}, <-chan nonceCheckResult) {
	items := make([]pow.Block, len(blocks))
	for i, block := range blocks {
		items[i] = block
	}
	return verifyNonces(checker, items)
}

// verifyNonces starts a concurrent nonce verification, returning a quit channel
// to abort the operations and a results channel to retrieve the async checks.
func verifyNonces(checker pow.PoW, items []pow.Block) (chan<- struct{}, <-chan nonceCheckResult) {
	// Spawn as many workers as allowed threads
	workers := runtime.GOMAXPROCS(0)
	if len(items) < workers {
		workers = len(items)
	}
	// Create a task channel and spawn the verifiers
	tasks := make(chan int, workers)
	results := make(chan nonceCheckResult, len(items)) // Buffered to make sure all workers stop
	for i := 0; i < workers; i++ {
		go func() {
			for index := range tasks {
				results <- nonceCheckResult{index: index, valid: checker.Verify(items[index])}
			}
		}()
	}
	// Feed item indices to the workers until done or aborted
	abort := make(chan struct{})
	go func() {
		defer close(tasks)

		for i := range items {
			select {
			case tasks <- i:
				continue
			case <-abort:
				return
			}
		}
	}()
	return abort, results
}
