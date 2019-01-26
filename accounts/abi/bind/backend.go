package bind

import (
	"errors"
	"math/big"

	"github.com/ether-core/go-ethereum/common"
	"github.com/ether-core/go-ethereum/core/types"
)

// ErrNoCode is returned by call and transact operations for which the requested
// recipient contract to operate on does not exist in the state db or does not
// have any code associated with it (i.e. suicided).
var ErrNoCode = errors.New("no contract code at given address")

// ContractCaller defines the methods needed to allow operating with contract on a read
// only basis.
type ContractCaller interface {
	// HasCode checks if the contract at the given address has any code associated
	// with it or not. This is needed to differentiate between contract internal
	// errors and the local chain being out of sync.
	HasCode(contract common.Address, pending bool) (bool, error)

	// ContractCall executes an Ethereum contract call with the specified data as
	// the input. The pending flag requests execution against the pending block, not
	// the stable head of the chain.
	ContractCall(contract common.Address, data []byte, pending bool) ([]byte, error)
}

// ContractTransactor defines the methods needed to allow operating with contract
// on a write only basis. Beside the transacting method, the remainder are helpers
// used when the user does not provide some needed values, but rather leaves it up
// to the transactor to decide.
type ContractTransactor interface {
	// PendingAccountNonce retrieves the current pending nonce associated with an
	// account.
	PendingAccountNonce(account common.Address) (uint64, error)

	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
	SuggestGasPrice() (*big.Int, error)

	// HasCode checks if the contract at the given address has any code associated
	// with it or not. This is needed to differentiate between contract internal
	// errors and the local chain being out of sync.
	HasCode(contract common.Address, pending bool) (bool, error)

	// EstimateGasLimit tries to estimate the gas needed to execute a specific
	// transaction based on the current pending state of the backend blockchain.
	// There is no guarantee that this is the true gas limit requirement as other
	// transactions may be added or removed by miners, but it should provide a basis
	// for setting a reasonable default.
	EstimateGasLimit(sender common.Address, contract *common.Address, value *big.Int, data []byte) (*big.Int, error)

	// SendTransaction injects the transaction into the pending pool for execution.
	SendTransaction(tx *types.Transaction) error
}

// ContractBackend defines the methods needed to allow operating with contract
// on a read-write basis.
//
// This interface is essentially the union of ContractCaller and ContractTransactor
// but due to a bug in the Go compiler (https://github.com/golang/go/issues/6977),
// we cannot simply list it as the two interfaces. The other solution is to add a
// third interface containing the common methods, but that convolutes the user API
// as it introduces yet another parameter to require for initialization.
type ContractBackend interface {
	// HasCode checks if the contract at the given address has any code associated
	// with it or not. This is needed to differentiate between contract internal
	// errors and the local chain being out of sync.
	HasCode(contract common.Address, pending bool) (bool, error)

	// ContractCall executes an Ethereum contract call with the specified data as
	// the input. The pending flag requests execution against the pending block, not
	// the stable head of the chain.
	ContractCall(contract common.Address, data []byte, pending bool) ([]byte, error)

	// PendingAccountNonce retrieves the current pending nonce associated with an
	// account.
	PendingAccountNonce(account common.Address) (uint64, error)

	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
	SuggestGasPrice() (*big.Int, error)

	// EstimateGasLimit tries to estimate the gas needed to execute a specific
	// transaction based on the current pending state of the backend blockchain.
	// There is no guarantee that this is the true gas limit requirement as other
	// transactions may be added or removed by miners, but it should provide a basis
	// for setting a reasonable default.
	EstimateGasLimit(sender common.Address, contract *common.Address, value *big.Int, data []byte) (*big.Int, error)

	// SendTransaction injects the transaction into the pending pool for execution.
	SendTransaction(tx *types.Transaction) error
}
