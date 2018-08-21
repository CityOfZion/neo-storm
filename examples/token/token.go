package token

import (
	// "github.com/CityOfZion/neo-storm/interop/action"
	// "github.com/CityOfZion/neo-storm/interop/blockchain"
	// "github.com/CityOfZion/neo-storm/interop/contract"
	// "github.com/CityOfZion/neo-storm/interop/executionEngine"
	"github.com/CityOfZion/neo-storm/interop/runtime"
	"github.com/CityOfZion/neo-storm/interop/storage"
	"github.com/CityOfZion/neo-storm/interop/util"
)

const (
	decimals   = 8
	multiplier = 100000000
)

var owner = util.FromAddress("AK2nJJpJr6o664CWJKi1QRXjqeic2zRp8y")

// CreateToken initializes the Token Interface for the Smart Contract to operate with
func CreateToken() Token {
	return Token{
		Name:           "Awesome NEO Token",
		Symbol:         "ANT",
		Decimals:       decimals,
		Owner:          owner,
		TotalSupply:    11000000 * multiplier,
		CirculationKey: "TokenCirculation",
	}
}

// Main function = contract entry
func Main(operation string, args []interface{}) interface{} {
	token := CreateToken()

	if operation == "name" {
		return token.Name
	}
	if operation == "symbol" {
		return token.Symbol
	}
	if operation == "decimals" {
		return token.Decimals
	}

	// The following operations need ctx
	ctx := storage.GetContext()

	if operation == "totalSupply" {
		return token.GetSupply(ctx)
	}
	if operation == "balanceOf" {
		hodler := args[0].([]byte)
		return token.BalanceOf(ctx, hodler)
	}
	if operation == "transfer" && CheckArgs(args, 3) {
		from := args[0].([]byte)
		to := args[1].([]byte)
		amount := args[2].(int)
		return token.Transfer(ctx, from, to, amount)
	}

	return true
}

// CheckArgs checks args array against a length indicator
func CheckArgs(args []interface{}, length int) bool {
	if len(args) == length {
		return true
	}

	return false
}

// TODO: move everything below this line to its separate package

// Token holds all token info
type Token struct {
	// Token name
	Name string
	// Ticker symbol
	Symbol string
	// Amount of decimals
	Decimals int
	// Token owner address
	Owner []byte
	// Total tokens * multiplier
	TotalSupply int
	// Storage key for circulation value
	CirculationKey string
}

// TODO: Transfer event
// DoTransfer := action.RegisterAction("transfer", "from", "to", "amount")

// GetSupply gets the token totalSupply value from VM storage
func (t Token) GetSupply(ctx storage.Context) interface{} {
	return storage.Get(ctx, t.CirculationKey)
}

// BalanceOf gets the token balance of a specific address
func (t Token) BalanceOf(ctx storage.Context, hodler []byte) interface{} {
	return storage.Get(ctx, hodler)
}

// Transfer token from one user to another
func (t Token) Transfer(ctx storage.Context, from []byte, to []byte, amount int) bool {
	amountFrom := CanTransfer(ctx, from, to, amount)
	if amountFrom == -1 {
		return false
	}

	if amountFrom == 0 {
		storage.Delete(ctx, from)
	}

	if amountFrom > 0 {
		diff := amountFrom - amount
		storage.Put(ctx, from, diff)
	}

	amountTo := storage.Get(ctx, to).(int)
	totalAmountTo := amountTo + amount
	storage.Put(ctx, to, totalAmountTo)
	// DoTransfer(from, to, amount)
	return true
}

// CanTransfer returns the amount it can transfer
func CanTransfer(ctx storage.Context, from []byte, to []byte, amount int) int {
	if len(to) != 20 && !IsUsableAddress(from) {
		return -1
	}

	amountFrom := storage.Get(ctx, from).(int)
	if amountFrom < amount {
		return -1
	}

	// Tell Transfer the result is equal - special case since it uses Delete
	if amountFrom == amount {
		return 0
	}

	// return amountFrom value back to Transfer, reduces extra Get
	return amountFrom
}

// IsUsableAddress checks if the sender is either the correct NEO address or SC address
func IsUsableAddress(from []byte) bool {
	if runtime.CheckWitness(from) {
		return true
	}

	// This method isn't implemented yet
	// TODO: check if contract is calling scripthash
	// if (contract.Contract{}) != blockchain.GetContract(from) && from == executionEngine.GetCallingScriptHash()
	// 	return true
	// }

	return false
}
