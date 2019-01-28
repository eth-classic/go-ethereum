package types

import (
	"fmt"
	"math/big"

	"github.com/openether/ethcore/common"
	"github.com/openether/ethcore/core/vm"
	"github.com/openether/ethcore/crypto"
)

const bloomLength = 256

type Bloom [bloomLength]byte

func BytesToBloom(b []byte) Bloom {
	var bloom Bloom
	bloom.SetBytes(b)
	return bloom
}

func (b *Bloom) SetBytes(d []byte) {
	if len(b) < len(d) {
		panic(fmt.Sprintf("bloom bytes too big %d %d", len(b), len(d)))
	}

	copy(b[bloomLength-len(d):], d)
}

func (b *Bloom) Add(d *big.Int) {
	bin := new(big.Int).SetBytes(b[:])
	bin.Or(bin, bloom9(d.Bytes()))
	b.SetBytes(bin.Bytes())
}

func (b Bloom) Bytes() []byte {
	return b[:]
}

func (b Bloom) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%#x"`, b.Bytes())), nil
}

func CreateBloom(receipts Receipts) Bloom {
	bin := new(big.Int)
	for _, receipt := range receipts {
		bin.Or(bin, LogsBloom(receipt.Logs))
	}

	return BytesToBloom(bin.Bytes())
}

func LogsBloom(logs vm.Logs) *big.Int {
	bin := new(big.Int)
	for _, log := range logs {
		data := make([]common.Hash, len(log.Topics))
		bin.Or(bin, bloom9(log.Address.Bytes()))

		for i, topic := range log.Topics {
			data[i] = topic
		}

		for _, b := range data {
			bin.Or(bin, bloom9(b[:]))
		}
	}

	return bin
}

func bloom9(b []byte) *big.Int {
	b = crypto.Keccak256(b[:])

	r := new(big.Int)

	for i := 0; i < 6; i += 2 {
		t := big.NewInt(1)
		b := (uint(b[i+1]) + (uint(b[i]) << 8)) & 2047
		r.Or(r, t.Lsh(t, b))
	}

	return r
}

var Bloom9 = bloom9

func BloomLookup(bin Bloom, topic []byte) bool {
	bloom := new(big.Int).SetBytes(bin[:])
	cmp := bloom9(topic)
	return bloom.And(bloom, cmp).Cmp(cmp) == 0
}
