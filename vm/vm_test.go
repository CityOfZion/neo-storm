package vm

import "testing"

func TestVm(t *testing.T) {
	vm := NewVM()
	script := []byte{0x00, byte(RET)}
	vm.Run(script)
}
