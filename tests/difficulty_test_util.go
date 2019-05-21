// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package tests

import (
	"math/big"
	"strconv"
	"testing"

	"github.com/eth-classic/go-ethereum/common"
	"github.com/eth-classic/go-ethereum/core"
	"github.com/eth-classic/go-ethereum/core/types"
	// "github.com/eth-classic/go-ethereum/core/"
)

//go:generate gencodec -type DifficultyTest -field-override difficultyTestMarshaling -out gen_difficultytest.go

// type DifficultyTest struct {
// 	ParentTimestamp    uint64      `json:"ParentTimestamp"`
// 	ParentDifficulty   uint64      `json:"parentDifficulty"`
// 	UncleHash          common.Hash `json:"parentUncles"`
// 	CurrentTimestamp   uint64      `json:"currentTimestamp"`
// 	CurrentBlockNumber uint64      `json:"currentBlockNumber"`
// 	CurrentDifficulty  uint64      `json:"currentDifficulty"`
// }

type DifficultyTest struct {
	ParentTimestamp    string
	ParentDifficulty   string
	UncleHash          common.Hash
	CurrentTimestamp   string
	CurrentBlockNumber string
	CurrentDifficulty  string
}

// type difficultyTestMarshaling struct {
// 	ParentTimestamp    math.HexOrDecimal64
// 	ParentDifficulty   *math.HexOrDecimal256
// 	CurrentTimestamp   math.HexOrDecimal64
// 	CurrentDifficulty  *math.HexOrDecimal256
// 	UncleHash          common.Hash
// 	CurrentBlockNumber math.HexOrDecimal64
// }

func RunDifficultyTests(path string, t *testing.T, config *core.ChainConfig) error {

	tests := make(map[string]DifficultyTest)

	//Ensure no problems w/ JSON file
	if err := readJsonFile(path, &tests); err != nil {
		return err
	}

	// return nil
	for name := range tests {

		test := tests[name]
		parentNumber, _ := new(big.Int).SetString(test.CurrentBlockNumber, 10)
		parentTimestamp, _ := new(big.Int).SetString(test.ParentTimestamp, 10)
		parentDifficulty, _ := new(big.Int).SetString(test.ParentDifficulty, 10)
		curTimestamp, _ := strconv.ParseUint(test.CurrentTimestamp, 10, 64)
		// parentDifficulty := big.NewInt(int64(test.ParentDifficulty))
		// parentDifficulty := big.NewInt(int64(test.ParentDifficulty))

		parent := &types.Header{
			Number:     parentNumber,
			Time:       parentTimestamp,
			Difficulty: parentDifficulty,
		}

		//How to define chainconfig?
		actual := core.CalcDifficulty(config, curTimestamp, parent)
		// actual := core.CalcDifficulty(core.DefaultConfigMainnet.ChainConfig, curTimestamp, parent)

		exp, _ := new(big.Int).SetString(test.CurrentDifficulty, 10)

		if actual.Cmp(exp) != 0 {
			// return fmt.Errorf("parent[time %v diff %v unclehash:%x] child[time %v number %v] diff %v != expected %v",
			// 	test.ParentTimestamp, test.ParentDifficulty, test.UncleHash,
			// 	test.CurrentTimestamp, test.CurrentBlockNumber, actual, exp)

			t.Error(name, "failed. Expected", test.CurrentDifficulty, "and calculated", actual)
		}

	}
	return nil

}

// func TestCalcDifficulty(t *testing.T) {
// 	file, err := os.Open(filepath.Join("..", "..", "tests", "testdata", "BasicTests", "difficulty.json"))
// 	if err != nil {
// 		t.Skip(err)
// 	}
// 	defer file.Close()

// 	tests := make(map[string]diffTest)
// 	err = json.NewDecoder(file).Decode(&tests)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	config := &params.ChainConfig{HomesteadBlock: big.NewInt(1150000)}

// 	for name, test := range tests {
// 		number := new(big.Int).Sub(test.CurrentBlocknumber, big.NewInt(1))
// 		diff := CalcDifficulty(config, test.CurrentTimestamp, &types.Header{
// 			Number:     number,
// 			Time:       test.ParentTimestamp,
// 			Difficulty: test.ParentDifficulty,
// 		})
// 		if diff.Cmp(test.CurrentDifficulty) != 0 {
// 			t.Error(name, "failed. Expected", test.CurrentDifficulty, "and calculated", diff)
// 		}
// 	}
// }
