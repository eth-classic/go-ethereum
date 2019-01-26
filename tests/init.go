package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ether-core/go-ethereum/core"
)

var (
	baseDir            = filepath.Join(".", "files")
	blockTestDir       = filepath.Join(baseDir, "BlockchainTests")
	stateTestDir       = filepath.Join(baseDir, "StateTests")
	transactionTestDir = filepath.Join(baseDir, "TransactionTests")
	vmTestDir          = filepath.Join(baseDir, "VMTests")
	rlpTestDir         = filepath.Join(baseDir, "RLPTests")

	BlockSkipTests = initBlockSkipTests()

	/* Go client does not support transaction (account) nonces above 2^64. This
	technically breaks consensus but is regarded as "reasonable
	engineering constraint" as accounts cannot easily reach such high
	nonce values in practice
	*/
	TransSkipTests = []string{"TransactionWithHihghNonce256"}
	StateSkipTests = []string{}
	VmSkipTests    = []string{}
)

func initBlockSkipTests() []string {
	if core.UseSputnikVM == "true" {
		return []string{
			// These tests are not valid, as they are out of scope for RLP and
			// the consensus protocol.
			"BLOCK__RandomByteAtTheEnd",
			"TRANSCT__RandomByteAtTheEnd",
			"BLOCK__ZeroByteAtTheEnd",
			"TRANSCT__ZeroByteAtTheEnd",

			"ChainAtoChainB_blockorder2",
			"ChainAtoChainB_blockorder1",
			"ChainAtoChainB_BlockHash",
			"CallingCanonicalContractFromFork_CALLCODE",
		}
	} else {
		return []string{
			// These tests are not valid, as they are out of scope for RLP and
			// the consensus protocol.
			"BLOCK__RandomByteAtTheEnd",
			"TRANSCT__RandomByteAtTheEnd",
			"BLOCK__ZeroByteAtTheEnd",
			"TRANSCT__ZeroByteAtTheEnd",

			"ChainAtoChainB_blockorder2",
			"ChainAtoChainB_blockorder1",
		}
	}
}

func readJson(reader io.Reader, value interface{}) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %v", err)
	}
	if err = json.Unmarshal(data, &value); err != nil {
		if syntaxerr, ok := err.(*json.SyntaxError); ok {
			line := findLine(data, syntaxerr.Offset)
			return fmt.Errorf("JSON syntax error at line %v: %v", line, err)
		}
		return fmt.Errorf("JSON unmarshal error: %v", err)
	}
	return nil
}

func readJsonFile(fn string, value interface{}) error {
	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	err = readJson(file, value)
	if err != nil {
		return fmt.Errorf("%s in file %s", err.Error(), fn)
	}
	return nil
}

// findLine returns the line number for the given offset into data.
func findLine(data []byte, offset int64) (line int) {
	line = 1
	for i, r := range string(data) {
		if int64(i) >= offset {
			return
		}
		if r == '\n' {
			line++
		}
	}
	return
}
