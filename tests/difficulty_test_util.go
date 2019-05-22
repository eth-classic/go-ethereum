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

type DifficultyTest struct {
	ParentTimestamp    string      `json:"parentTimestamp"`
	ParentDifficulty   string      `json:"parentDifficulty"`
	UncleHash          common.Hash `json:"parentUncles"`
	CurrentTimestamp   string      `json:"currentTimestamp"`
	CurrentBlockNumber string      `json:"currentBlockNumber"`
	CurrentDifficulty  string      `json:"currentDifficulty"`
}

func (test *DifficultyTest) runDifficulty(config *core.ChainConfig) error {
	parentNumber, _ := new(big.Int).SetString(test.CurrentBlockNumber, 10)
	parentTimestamp, _ := strconv.ParseUint(test.ParentTimestamp, 10, 64)
	parentDifficulty, _ := new(big.Int).SetString(test.ParentDifficulty, 10)
	currentTimestamp, _ := strconv.ParseUint(test.CurrentTimestamp, 10, 64)

	actual := core.CalcDifficulty(config, currentTimestamp, parentTimestamp, parentNumber, parentDifficulty)
	exp, _ := new(big.Int).SetString(test.CurrentDifficulty, 10)

	if actual.Cmp(exp) != 0 {
		return fmt.Errorf("parent[time %v diff %v unclehash:%x] child[time %v number %v] diff %v != expected %v",
			test.ParentTimestamp, test.ParentDifficulty, test.UncleHash,
			test.CurrentTimestamp, test.CurrentBlockNumber, actual, exp)
	}
	return nil

}
