package vm

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/openether/ethcore/common"
	"github.com/openether/ethcore/rlp"
)

type Log struct {
	// Consensus fields
	Address common.Address
	Topics  []common.Hash
	Data    []byte

	// Derived fields (don't reorder!)
	BlockNumber uint64
	TxHash      common.Hash
	TxIndex     uint
	BlockHash   common.Hash
	Index       uint
}

func NewLog(address common.Address, topics []common.Hash, data []byte, number uint64) *Log {
	return &Log{Address: address, Topics: topics, Data: data, BlockNumber: number}
}

func (l *Log) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{l.Address, l.Topics, l.Data})
}

func (l *Log) DecodeRLP(s *rlp.Stream) error {
	var log struct {
		Address common.Address
		Topics  []common.Hash
		Data    []byte
	}
	if err := s.Decode(&log); err != nil {
		return err
	}
	l.Address, l.Topics, l.Data = log.Address, log.Topics, log.Data
	return nil
}

func (l *Log) String() string {
	return fmt.Sprintf(`log: %x %x %x %x %d %x %d`, l.Address, l.Topics, l.Data, l.TxHash, l.TxIndex, l.BlockHash, l.Index)
}

func (r *Log) MarshalJSON() ([]byte, error) {
	fields := map[string]interface{}{
		"address":          r.Address,
		"data":             fmt.Sprintf("%#x", r.Data),
		"blockNumber":      fmt.Sprintf("%#x", r.BlockNumber),
		"logIndex":         fmt.Sprintf("%#x", r.Index),
		"blockHash":        r.BlockHash,
		"transactionHash":  r.TxHash,
		"transactionIndex": fmt.Sprintf("%#x", r.TxIndex),
		"topics":           r.Topics,
	}

	return json.Marshal(fields)
}

type Logs []*Log

// LogForStorage is a wrapper around a Log that flattens and parses the entire
// content of a log, as opposed to only the consensus fields originally (by hiding
// the rlp interface methods).
type LogForStorage Log
