// Copyright 2015 The go-ethereum Authors
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
	"fmt"
	"math/big"
	"strconv"

	"github.com/eth-classic/go-ethereum/common"
	"github.com/eth-classic/go-ethereum/core"
)

// DifficultyTest is the structure of JSON from test files
type DifficultyTest struct {
	ParentTimestamp    string      `json:"parentTimestamp"`
	ParentDifficulty   string      `json:"parentDifficulty"`
	UncleHash          common.Hash `json:"parentUncles"`
	CurrentTimestamp   string      `json:"currentTimestamp"`
	CurrentBlockNumber string      `json:"currentBlockNumber"`
	CurrentDifficulty  string      `json:"currentDifficulty"`
}

func (test *DifficultyTest) runDifficulty(config *core.ChainConfig) error {
	currentNumber, _ := ParseBigInt(test.CurrentBlockNumber)
	parentNumber := new(big.Int).Sub(currentNumber, big.NewInt(1))
	parentTimestamp, _ := ParseUint64(test.ParentTimestamp)
	parentDifficulty, _ := ParseBigInt(test.ParentDifficulty)
	currentTimestamp, _ := ParseUint64(test.CurrentTimestamp)

	actual := core.CalcDifficulty(config, currentTimestamp, parentTimestamp, parentNumber, parentDifficulty)
	exp, _ := ParseBigInt(test.CurrentDifficulty)

	if actual.Cmp(exp) != 0 {
		return fmt.Errorf("parent[time %v diff %v unclehash:%x] child[time %v number %v] diff %v != expected %v",
			test.ParentTimestamp, test.ParentDifficulty, test.UncleHash,
			test.CurrentTimestamp, test.CurrentBlockNumber, actual, exp)
	}
	return nil

}

// ParseUint64 parses ambiguous string of hex/decimal into *big.Int
func ParseUint64(s string) (uint64, bool) {
	if s == "" {
		return 0, true
	}
	if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		v, err := strconv.ParseUint(s[2:], 16, 64)
		return v, err == nil
	}
	v, err := strconv.ParseUint(s, 10, 64)
	return v, err == nil
}

// ParseBigInt parses ambiguous string of hex/decimal into *big.Int
func ParseBigInt(s string) (*big.Int, bool) {
	if s == "" {
		return new(big.Int), true
	}
	var bigint *big.Int
	var ok bool
	if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		bigint, ok = new(big.Int).SetString(s[2:], 16)
	} else {
		bigint, ok = new(big.Int).SetString(s, 10)
	}
	if ok && bigint.BitLen() > 256 {
		bigint, ok = nil, false
	}
	return bigint, ok
}
