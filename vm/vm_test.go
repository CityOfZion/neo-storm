package vm

import "testing"

func TestVm(t *testing.T) {
	vm := NewVM()
	script := []byte{0x00, byte(RET)}
	// FIXME: continuous loop in this for the momement
	// vm.Run(script)
}
