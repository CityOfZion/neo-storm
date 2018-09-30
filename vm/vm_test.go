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

func TestAnd(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH4, PUSH4, AND)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop(), NewStackItem(4))
}

func TestOr(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH4, PUSH2, OR)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop(), NewStackItem(6))
}

func TestXor(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH7, PUSH2, XOR)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop(), NewStackItem(5))
}

func TestShl(t *testing.T) {
	// TODO
}

func TestShr(t *testing.T) {
	// TODO
}

func TestInc(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH7, INC)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop(), NewStackItem(8))
}

func TestDec(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH7, DEC)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop(), NewStackItem(6))
}

func TestNewArray(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH4, NEWARRAY)
	vm.Run(script)
	assert.Equal(t, vm.estack.Pop(), NewStackItem(make([]*StackItem, 4)))
}

func TestAppend(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH2, NEWARRAY, PUSH4, APPEND)
	vm.Run(script)
	arr := vm.estack.Pop().Array()
	assert.Equal(t, 3, len(arr))
	assert.Equal(t, arr[len(arr)-1], NewStackItem(4))
}

func TestAppendBytes(t *testing.T) {
	vm := NewVM()
	script := createScript(
		0x03, // PUSHBYTES3
		Instruction(byte('a')),
		Instruction(byte('b')),
		Instruction(byte('c')),
		0x02, // PUSHBYTES2
		Instruction(byte('d')),
		Instruction(byte('e')),
		APPEND,
	)
	vm.Run(script)
	assert.Equal(t, NewStackItem([]byte("abcde")), vm.estack.Pop())
}

func TestPack(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH1, PUSH2, PUSH3, PUSH3, PACK)
	vm.Run(script)
	assert.Equal(t, 1, vm.estack.Len())
	assert.Equal(t, 3, len(vm.estack.Pop().Array()))
	assert.Equal(t, 0, vm.estack.Len())
}

func TestPickItem(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH1, PUSH2, PUSH3, PUSH3, PACK, PUSH1, PICKITEM)
	vm.Run(script)
	assert.Equal(t, NewStackItem(2), vm.estack.Pop())
}

func TestSetItem(t *testing.T) {
	vm := NewVM()
	script := createScript(PUSH1, PUSH2, PUSH3, PUSH3, PACK, PUSH1, PUSH7, SETITEM, PUSH1, PICKITEM) //, PUSH1, PICKITEM)
	vm.Run(script)
	assert.Equal(t, NewStackItem(7), vm.estack.Pop())

}

func createScript(instructions ...Instruction) []byte {
	script := make([]byte, len(instructions))
	for i, instr := range instructions {
		script[i] = byte(instr)
	}
	return script
}
