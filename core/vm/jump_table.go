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

package vm

import "math/big"

type jumpPtr struct {
	fn    instrFn
	valid bool
}

type vmJumpTable [256]jumpPtr

func newJumpTable(ruleset RuleSet, blockNumber *big.Int) vmJumpTable {
	jumpTable := newFrontierInstructionSet()

	// when initialising a new VM execution we must first check the homestead
	// changes.
	if ruleset.IsHomestead(blockNumber) {
		jumpTable[DELEGATECALL] = jumpPtr{
			fn:    opDelegateCall,
			valid: true,
		}
	}

	jumpTable[ADD] = jumpPtr{opAdd, true}
	jumpTable[SUB] = jumpPtr{opSub, true}
	jumpTable[MUL] = jumpPtr{opMul, true}
	jumpTable[DIV] = jumpPtr{opDiv, true}
	jumpTable[SDIV] = jumpPtr{opSdiv, true}
	jumpTable[MOD] = jumpPtr{opMod, true}
	jumpTable[SMOD] = jumpPtr{opSmod, true}
	jumpTable[EXP] = jumpPtr{opExp, true}
	jumpTable[SIGNEXTEND] = jumpPtr{opSignExtend, true}
	jumpTable[NOT] = jumpPtr{opNot, true}
	jumpTable[LT] = jumpPtr{opLt, true}
	jumpTable[GT] = jumpPtr{opGt, true}
	jumpTable[SLT] = jumpPtr{opSlt, true}
	jumpTable[SGT] = jumpPtr{opSgt, true}
	jumpTable[EQ] = jumpPtr{opEq, true}
	jumpTable[ISZERO] = jumpPtr{opIszero, true}
	jumpTable[AND] = jumpPtr{opAnd, true}
	jumpTable[OR] = jumpPtr{opOr, true}
	jumpTable[XOR] = jumpPtr{opXor, true}
	jumpTable[BYTE] = jumpPtr{opByte, true}
	jumpTable[ADDMOD] = jumpPtr{opAddmod, true}
	jumpTable[MULMOD] = jumpPtr{opMulmod, true}
	jumpTable[SHA3] = jumpPtr{opSha3, true}
	jumpTable[ADDRESS] = jumpPtr{opAddress, true}
	jumpTable[BALANCE] = jumpPtr{opBalance, true}
	jumpTable[ORIGIN] = jumpPtr{opOrigin, true}
	jumpTable[CALLER] = jumpPtr{opCaller, true}
	jumpTable[CALLVALUE] = jumpPtr{opCallValue, true}
	jumpTable[CALLDATALOAD] = jumpPtr{opCalldataLoad, true}
	jumpTable[CALLDATASIZE] = jumpPtr{opCalldataSize, true}
	jumpTable[CALLDATACOPY] = jumpPtr{opCalldataCopy, true}
	jumpTable[CODESIZE] = jumpPtr{opCodeSize, true}
	jumpTable[EXTCODESIZE] = jumpPtr{opExtCodeSize, true}
	jumpTable[CODECOPY] = jumpPtr{opCodeCopy, true}
	jumpTable[EXTCODECOPY] = jumpPtr{opExtCodeCopy, true}
	jumpTable[GASPRICE] = jumpPtr{opGasprice, true}
	jumpTable[BLOCKHASH] = jumpPtr{opBlockhash, true}
	jumpTable[COINBASE] = jumpPtr{opCoinbase, true}
	jumpTable[TIMESTAMP] = jumpPtr{opTimestamp, true}
	jumpTable[NUMBER] = jumpPtr{opNumber, true}
	jumpTable[DIFFICULTY] = jumpPtr{opDifficulty, true}
	jumpTable[GASLIMIT] = jumpPtr{opGasLimit, true}
	jumpTable[POP] = jumpPtr{opPop, true}
	jumpTable[MLOAD] = jumpPtr{opMload, true}
	jumpTable[MSTORE] = jumpPtr{opMstore, true}
	jumpTable[MSTORE8] = jumpPtr{opMstore8, true}
	jumpTable[SLOAD] = jumpPtr{opSload, true}
	jumpTable[SSTORE] = jumpPtr{opSstore, true}
	jumpTable[JUMPDEST] = jumpPtr{opJumpdest, true}
	jumpTable[PC] = jumpPtr{nil, true}
	jumpTable[MSIZE] = jumpPtr{opMsize, true}
	jumpTable[GAS] = jumpPtr{opGas, true}
	jumpTable[CREATE] = jumpPtr{opCreate, true}
	jumpTable[CALL] = jumpPtr{opCall, true}
	jumpTable[CALLCODE] = jumpPtr{opCallCode, true}
	jumpTable[LOG0] = jumpPtr{makeLog(0), true}
	jumpTable[LOG1] = jumpPtr{makeLog(1), true}
	jumpTable[LOG2] = jumpPtr{makeLog(2), true}
	jumpTable[LOG3] = jumpPtr{makeLog(3), true}
	jumpTable[LOG4] = jumpPtr{makeLog(4), true}
	jumpTable[SWAP1] = jumpPtr{makeSwap(1), true}
	jumpTable[SWAP2] = jumpPtr{makeSwap(2), true}
	jumpTable[SWAP3] = jumpPtr{makeSwap(3), true}
	jumpTable[SWAP4] = jumpPtr{makeSwap(4), true}
	jumpTable[SWAP5] = jumpPtr{makeSwap(5), true}
	jumpTable[SWAP6] = jumpPtr{makeSwap(6), true}
	jumpTable[SWAP7] = jumpPtr{makeSwap(7), true}
	jumpTable[SWAP8] = jumpPtr{makeSwap(8), true}
	jumpTable[SWAP9] = jumpPtr{makeSwap(9), true}
	jumpTable[SWAP10] = jumpPtr{makeSwap(10), true}
	jumpTable[SWAP11] = jumpPtr{makeSwap(11), true}
	jumpTable[SWAP12] = jumpPtr{makeSwap(12), true}
	jumpTable[SWAP13] = jumpPtr{makeSwap(13), true}
	jumpTable[SWAP14] = jumpPtr{makeSwap(14), true}
	jumpTable[SWAP15] = jumpPtr{makeSwap(15), true}
	jumpTable[SWAP16] = jumpPtr{makeSwap(16), true}
	jumpTable[PUSH1] = jumpPtr{makePush(1, big.NewInt(1)), true}
	jumpTable[PUSH2] = jumpPtr{makePush(2, big.NewInt(2)), true}
	jumpTable[PUSH3] = jumpPtr{makePush(3, big.NewInt(3)), true}
	jumpTable[PUSH4] = jumpPtr{makePush(4, big.NewInt(4)), true}
	jumpTable[PUSH5] = jumpPtr{makePush(5, big.NewInt(5)), true}
	jumpTable[PUSH6] = jumpPtr{makePush(6, big.NewInt(6)), true}
	jumpTable[PUSH7] = jumpPtr{makePush(7, big.NewInt(7)), true}
	jumpTable[PUSH8] = jumpPtr{makePush(8, big.NewInt(8)), true}
	jumpTable[PUSH9] = jumpPtr{makePush(9, big.NewInt(9)), true}
	jumpTable[PUSH10] = jumpPtr{makePush(10, big.NewInt(10)), true}
	jumpTable[PUSH11] = jumpPtr{makePush(11, big.NewInt(11)), true}
	jumpTable[PUSH12] = jumpPtr{makePush(12, big.NewInt(12)), true}
	jumpTable[PUSH13] = jumpPtr{makePush(13, big.NewInt(13)), true}
	jumpTable[PUSH14] = jumpPtr{makePush(14, big.NewInt(14)), true}
	jumpTable[PUSH15] = jumpPtr{makePush(15, big.NewInt(15)), true}
	jumpTable[PUSH16] = jumpPtr{makePush(16, big.NewInt(16)), true}
	jumpTable[PUSH17] = jumpPtr{makePush(17, big.NewInt(17)), true}
	jumpTable[PUSH18] = jumpPtr{makePush(18, big.NewInt(18)), true}
	jumpTable[PUSH19] = jumpPtr{makePush(19, big.NewInt(19)), true}
	jumpTable[PUSH20] = jumpPtr{makePush(20, big.NewInt(20)), true}
	jumpTable[PUSH21] = jumpPtr{makePush(21, big.NewInt(21)), true}
	jumpTable[PUSH22] = jumpPtr{makePush(22, big.NewInt(22)), true}
	jumpTable[PUSH23] = jumpPtr{makePush(23, big.NewInt(23)), true}
	jumpTable[PUSH24] = jumpPtr{makePush(24, big.NewInt(24)), true}
	jumpTable[PUSH25] = jumpPtr{makePush(25, big.NewInt(25)), true}
	jumpTable[PUSH26] = jumpPtr{makePush(26, big.NewInt(26)), true}
	jumpTable[PUSH27] = jumpPtr{makePush(27, big.NewInt(27)), true}
	jumpTable[PUSH28] = jumpPtr{makePush(28, big.NewInt(28)), true}
	jumpTable[PUSH29] = jumpPtr{makePush(29, big.NewInt(29)), true}
	jumpTable[PUSH30] = jumpPtr{makePush(30, big.NewInt(30)), true}
	jumpTable[PUSH31] = jumpPtr{makePush(31, big.NewInt(31)), true}
	jumpTable[PUSH32] = jumpPtr{makePush(32, big.NewInt(32)), true}
	jumpTable[DUP1] = jumpPtr{makeDup(1), true}
	jumpTable[DUP2] = jumpPtr{makeDup(2), true}
	jumpTable[DUP3] = jumpPtr{makeDup(3), true}
	jumpTable[DUP4] = jumpPtr{makeDup(4), true}
	jumpTable[DUP5] = jumpPtr{makeDup(5), true}
	jumpTable[DUP6] = jumpPtr{makeDup(6), true}
	jumpTable[DUP7] = jumpPtr{makeDup(7), true}
	jumpTable[DUP8] = jumpPtr{makeDup(8), true}
	jumpTable[DUP9] = jumpPtr{makeDup(9), true}
	jumpTable[DUP10] = jumpPtr{makeDup(10), true}
	jumpTable[DUP11] = jumpPtr{makeDup(11), true}
	jumpTable[DUP12] = jumpPtr{makeDup(12), true}
	jumpTable[DUP13] = jumpPtr{makeDup(13), true}
	jumpTable[DUP14] = jumpPtr{makeDup(14), true}
	jumpTable[DUP15] = jumpPtr{makeDup(15), true}
	jumpTable[DUP16] = jumpPtr{makeDup(16), true}

	jumpTable[RETURN] = jumpPtr{nil, true}
	jumpTable[SUICIDE] = jumpPtr{nil, true}
	jumpTable[JUMP] = jumpPtr{nil, true}
	jumpTable[JUMPI] = jumpPtr{nil, true}
	jumpTable[STOP] = jumpPtr{nil, true}
	jumpTable[STATICCALL] = jumpPtr{opStaticCall, true}

	return jumpTable
}

func newFrontierInstructionSet() vmJumpTable {
	return vmJumpTable{
		ADD: {
			fn:    opAdd,
			valid: true,
		},
		SUB: {
			fn:    opSub,
			valid: true,
		},
		MUL: {
			fn:    opMul,
			valid: true,
		},
		DIV: {
			fn:    opDiv,
			valid: true,
		},
		SDIV: {
			fn:    opSdiv,
			valid: true,
		},
		MOD: {
			fn:    opMod,
			valid: true,
		},
		SMOD: {
			fn:    opSmod,
			valid: true,
		},
		EXP: {
			fn:    opExp,
			valid: true,
		},
		SIGNEXTEND: {
			fn:    opSignExtend,
			valid: true,
		},
		NOT: {
			fn:    opNot,
			valid: true,
		},
		LT: {
			fn:    opLt,
			valid: true,
		},
		GT: {
			fn:    opGt,
			valid: true,
		},
		SLT: {
			fn:    opSlt,
			valid: true,
		},
		SGT: {
			fn:    opSgt,
			valid: true,
		},
		EQ: {
			fn:    opEq,
			valid: true,
		},
		ISZERO: {
			fn:    opIszero,
			valid: true,
		},
		AND: {
			fn:    opAnd,
			valid: true,
		},
		OR: {
			fn:    opOr,
			valid: true,
		},
		XOR: {
			fn:    opXor,
			valid: true,
		},
		BYTE: {
			fn:    opByte,
			valid: true,
		},
		ADDMOD: {
			fn:    opAddmod,
			valid: true,
		},
		MULMOD: {
			fn:    opMulmod,
			valid: true,
		},
		SHA3: {
			fn:    opSha3,
			valid: true,
		},
		ADDRESS: {
			fn:    opAddress,
			valid: true,
		},
		BALANCE: {
			fn:    opBalance,
			valid: true,
		},
		ORIGIN: {
			fn:    opOrigin,
			valid: true,
		},
		CALLER: {
			fn:    opCaller,
			valid: true,
		},
		CALLVALUE: {
			fn:    opCallValue,
			valid: true,
		},
		CALLDATALOAD: {
			fn:    opCalldataLoad,
			valid: true,
		},
		CALLDATASIZE: {
			fn:    opCalldataSize,
			valid: true,
		},
		CALLDATACOPY: {
			fn:    opCalldataCopy,
			valid: true,
		},
		CODESIZE: {
			fn:    opCodeSize,
			valid: true,
		},
		EXTCODESIZE: {
			fn:    opExtCodeSize,
			valid: true,
		},
		CODECOPY: {
			fn:    opCodeCopy,
			valid: true,
		},
		EXTCODECOPY: {
			fn:    opExtCodeCopy,
			valid: true,
		},
		GASPRICE: {
			fn:    opGasprice,
			valid: true,
		},
		BLOCKHASH: {
			fn:    opBlockhash,
			valid: true,
		},
		COINBASE: {
			fn:    opCoinbase,
			valid: true,
		},
		TIMESTAMP: {
			fn:    opTimestamp,
			valid: true,
		},
		NUMBER: {
			fn:    opNumber,
			valid: true,
		},
		DIFFICULTY: {
			fn:    opDifficulty,
			valid: true,
		},
		GASLIMIT: {
			fn:    opGasLimit,
			valid: true,
		},
		POP: {
			fn:    opPop,
			valid: true,
		},
		MLOAD: {
			fn:    opMload,
			valid: true,
		},
		MSTORE: {
			fn:    opMstore,
			valid: true,
		},
		MSTORE8: {
			fn:    opMstore8,
			valid: true,
		},
		SLOAD: {
			fn:    opSload,
			valid: true,
		},
		SSTORE: {
			fn:    opSstore,
			valid: true,
		},
		JUMPDEST: {
			fn:    opJumpdest,
			valid: true,
		},
		PC: {
			fn:    nil,
			valid: true,
		},
		MSIZE: {
			fn:    opMsize,
			valid: true,
		},
		GAS: {
			fn:    opGas,
			valid: true,
		},
		CREATE: {
			fn:    opCreate,
			valid: true,
		},
		CALL: {
			fn:    opCall,
			valid: true,
		},
		CALLCODE: {
			fn:    opCallCode,
			valid: true,
		},
		LOG0: {
			fn:    makeLog(0),
			valid: true,
		},
		LOG1: {
			fn:    makeLog(1),
			valid: true,
		},
		LOG2: {
			fn:    makeLog(2),
			valid: true,
		},
		LOG3: {
			fn:    makeLog(3),
			valid: true,
		},
		LOG4: {
			fn:    makeLog(4),
			valid: true,
		},
		SWAP1: {
			fn:    makeSwap(1),
			valid: true,
		},
		SWAP2: {
			fn:    makeSwap(2),
			valid: true,
		},
		SWAP3: {
			fn:    makeSwap(3),
			valid: true,
		},
		SWAP4: {
			fn:    makeSwap(4),
			valid: true,
		},
		SWAP5: {
			fn:    makeSwap(5),
			valid: true,
		},
		SWAP6: {
			fn:    makeSwap(6),
			valid: true,
		},
		SWAP7: {
			fn:    makeSwap(7),
			valid: true,
		},
		SWAP8: {
			fn:    makeSwap(8),
			valid: true,
		},
		SWAP9: {
			fn:    makeSwap(9),
			valid: true,
		},
		SWAP10: {
			fn:    makeSwap(10),
			valid: true,
		},
		SWAP11: {
			fn:    makeSwap(11),
			valid: true,
		},
		SWAP12: {
			fn:    makeSwap(12),
			valid: true,
		},
		SWAP13: {
			fn:    makeSwap(13),
			valid: true,
		},
		SWAP14: {
			fn:    makeSwap(14),
			valid: true,
		},
		SWAP15: {
			fn:    makeSwap(15),
			valid: true,
		},
		SWAP16: {
			fn:    makeSwap(16),
			valid: true,
		},
		PUSH1: {
			fn:    makePush(1, big.NewInt(1)),
			valid: true,
		},
		PUSH2: {
			fn:    makePush(2, big.NewInt(2)),
			valid: true,
		},
		PUSH3: {
			fn:    makePush(3, big.NewInt(3)),
			valid: true,
		},
		PUSH4: {
			fn:    makePush(4, big.NewInt(4)),
			valid: true,
		},
		PUSH5: {
			fn:    makePush(5, big.NewInt(5)),
			valid: true,
		},
		PUSH6: {
			fn:    makePush(6, big.NewInt(6)),
			valid: true,
		},
		PUSH7: {
			fn:    makePush(7, big.NewInt(7)),
			valid: true,
		},
		PUSH8: {
			fn:    makePush(8, big.NewInt(8)),
			valid: true,
		},
		PUSH9: {
			fn:    makePush(9, big.NewInt(9)),
			valid: true,
		},
		PUSH10: {
			fn:    makePush(10, big.NewInt(10)),
			valid: true,
		},
		PUSH11: {
			fn:    makePush(11, big.NewInt(11)),
			valid: true,
		},
		PUSH12: {
			fn:    makePush(12, big.NewInt(12)),
			valid: true,
		},
		PUSH13: {
			fn:    makePush(13, big.NewInt(13)),
			valid: true,
		},
		PUSH14: {
			fn:    makePush(14, big.NewInt(14)),
			valid: true,
		},
		PUSH15: {
			fn:    makePush(15, big.NewInt(15)),
			valid: true,
		},
		PUSH16: {
			fn:    makePush(16, big.NewInt(16)),
			valid: true,
		},
		PUSH17: {
			fn:    makePush(17, big.NewInt(17)),
			valid: true,
		},
		PUSH18: {
			fn:    makePush(18, big.NewInt(18)),
			valid: true,
		},
		PUSH19: {
			fn:    makePush(19, big.NewInt(19)),
			valid: true,
		},
		PUSH20: {
			fn:    makePush(20, big.NewInt(20)),
			valid: true,
		},
		PUSH21: {
			fn:    makePush(21, big.NewInt(21)),
			valid: true,
		},
		PUSH22: {
			fn:    makePush(22, big.NewInt(22)),
			valid: true,
		},
		PUSH23: {
			fn:    makePush(23, big.NewInt(23)),
			valid: true,
		},
		PUSH24: {
			fn:    makePush(24, big.NewInt(24)),
			valid: true,
		},
		PUSH25: {
			fn:    makePush(25, big.NewInt(25)),
			valid: true,
		},
		PUSH26: {
			fn:    makePush(26, big.NewInt(26)),
			valid: true,
		},
		PUSH27: {
			fn:    makePush(27, big.NewInt(27)),
			valid: true,
		},
		PUSH28: {
			fn:    makePush(28, big.NewInt(28)),
			valid: true,
		},
		PUSH29: {
			fn:    makePush(29, big.NewInt(29)),
			valid: true,
		},
		PUSH30: {
			fn:    makePush(30, big.NewInt(30)),
			valid: true,
		},
		PUSH31: {
			fn:    makePush(31, big.NewInt(31)),
			valid: true,
		},
		PUSH32: {
			fn:    makePush(32, big.NewInt(32)),
			valid: true,
		},
		DUP1: {
			fn:    makeDup(1),
			valid: true,
		},
		DUP2: {
			fn:    makeDup(2),
			valid: true,
		},
		DUP3: {
			fn:    makeDup(3),
			valid: true,
		},
		DUP4: {
			fn:    makeDup(4),
			valid: true,
		},
		DUP5: {
			fn:    makeDup(5),
			valid: true,
		},
		DUP6: {
			fn:    makeDup(6),
			valid: true,
		},
		DUP7: {
			fn:    makeDup(7),
			valid: true,
		},
		DUP8: {
			fn:    makeDup(8),
			valid: true,
		},
		DUP9: {
			fn:    makeDup(9),
			valid: true,
		},
		DUP10: {
			fn:    makeDup(10),
			valid: true,
		},
		DUP11: {
			fn:    makeDup(11),
			valid: true,
		},
		DUP12: {
			fn:    makeDup(12),
			valid: true,
		},
		DUP13: {
			fn:    makeDup(13),
			valid: true,
		},
		DUP14: {
			fn:    makeDup(14),
			valid: true,
		},
		DUP15: {
			fn:    makeDup(15),
			valid: true,
		},
		DUP16: {
			fn:    makeDup(16),
			valid: true,
		},
		RETURN: {
			fn:    nil,
			valid: true,
		},
		SUICIDE: {
			fn:    nil,
			valid: true,
		},
		JUMP: {
			fn:    nil,
			valid: true,
		},
		JUMPI: {
			fn:    nil,
			valid: true,
		},
		STOP: {
			fn:    nil,
			valid: true,
		},
	}
}
