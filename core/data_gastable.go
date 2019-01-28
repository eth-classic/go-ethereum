package core

import (
	"math/big"

	"github.com/openether/ethcore/core/vm"
)

var DefaultHomeSteadGasTable = &vm.GasTable{
	ExtcodeSize:     big.NewInt(20),
	ExtcodeCopy:     big.NewInt(20),
	Balance:         big.NewInt(20),
	SLoad:           big.NewInt(50),
	Calls:           big.NewInt(40),
	Suicide:         big.NewInt(0),
	ExpByte:         big.NewInt(10),
	CreateBySuicide: nil,
}

var DefaultGasRepriceGasTable = &vm.GasTable{
	ExtcodeSize:     big.NewInt(700),
	ExtcodeCopy:     big.NewInt(700),
	Balance:         big.NewInt(400),
	SLoad:           big.NewInt(200),
	Calls:           big.NewInt(700),
	Suicide:         big.NewInt(5000),
	ExpByte:         big.NewInt(10),
	CreateBySuicide: big.NewInt(25000),
}

var DefaultDiehardGasTable = &vm.GasTable{
	ExtcodeSize:     big.NewInt(700),
	ExtcodeCopy:     big.NewInt(700),
	Balance:         big.NewInt(400),
	SLoad:           big.NewInt(200),
	Calls:           big.NewInt(700),
	Suicide:         big.NewInt(5000),
	ExpByte:         big.NewInt(50),
	CreateBySuicide: big.NewInt(25000),
}
