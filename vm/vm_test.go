package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH4, PUSH2, SUB)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(2))
	assert.Equal(t, vm.estack.Len(), 0)
}

func TestAdd(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH4, PUSH2, ADD)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(6))
	assert.Equal(t, vm.estack.Len(), 0)
}

func TestMul(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH4, PUSH2, MUL)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(8))
	assert.Equal(t, vm.estack.Len(), 0)
}

func TestDiv(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH8, PUSH4, DIV)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(2))
	assert.Equal(t, vm.estack.Len(), 0)
}

func TestMod(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH8, PUSH6, MOD)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(2))
	assert.Equal(t, vm.estack.Len(), 0)
}

func TestDup(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH2, DUP)
	vm.Run(script)
	assert.Equal(t, vm.estack.Len(), 2)
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(2))
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(2))
}

func TestSwap(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH2, PUSH5, SWAP)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(2))
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(5))
}

func TestXswap(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH2, PUSH5, PUSH8, PUSH2, XSWAP)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(2))
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(5))
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(8))
}

func TestTuck(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH2, PUSH5, PUSH8, PUSH3, TUCK)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(8))
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(5))
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(2))
	assert.Equal(t, vm.estack.Pop().BigInt().Int64(), int64(8))
}

func TestShl(t *testing.T) {
	// TODO
}

func TestShr(t *testing.T) {
	// TODO
}

func createScript(instructions ...Instruction) []byte {
	script := make([]byte, len(instructions))
	for i, instr := range instructions {
		script[i] = byte(instr)
	}
	return script
}
