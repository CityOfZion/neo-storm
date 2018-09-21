package vm

type VMError int

const (
	OutOfGas VMError = iota
)
