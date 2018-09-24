package vm

import (
	"testing"
)

func TestVm(t *testing.T) {
	vm := NewVM()
	script := []byte{0x03, byte('a'), byte('n'), byte('t')}
	vm.Run(script)
	vm.estack.Inspect()
}
