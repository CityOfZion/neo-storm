package compiler

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/CityOfZion/neo-storm/vm"
)

func emit(w *bytes.Buffer, instr vm.Instruction, b []byte) error {
	if err := w.WriteByte(byte(instr)); err != nil {
		return err
	}
	_, err := w.Write(b)
	return err
}

func emitOpcode(w *bytes.Buffer, instr vm.Instruction) error {
	return w.WriteByte(byte(instr))
}

func emitBool(w *bytes.Buffer, ok bool) error {
	if ok {
		return emitOpcode(w, vm.PUSHT)
	}
	return emitOpcode(w, vm.PUSHF)
}

func emitInt(w *bytes.Buffer, i int64) error {
	if i == -1 {
		return emitOpcode(w, vm.PUSHM1)
	}
	if i == 0 {
		return emitOpcode(w, vm.PUSHF)
	}
	if i > 0 && i < 16 {
		val := vm.Instruction((int(vm.PUSH1) - 1 + int(i)))
		return emitOpcode(w, val)
	}

	bInt := big.NewInt(i)
	val := arrayReverse(bInt.Bytes())
	return emitBytes(w, val)
}

func emitString(w *bytes.Buffer, s string) error {
	return emitBytes(w, []byte(s))
}

func emitBytes(w *bytes.Buffer, b []byte) error {
	var (
		err error
		n   = len(b)
	)

	if n <= int(vm.PUSHBYTES75) {
		return emit(w, vm.Instruction(n), b)
	} else if n < 0x100 {
		err = emit(w, vm.PUSHDATA1, []byte{byte(n)})
	} else if n < 0x10000 {
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, uint16(n))
		err = emit(w, vm.PUSHDATA2, buf)
	} else {
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, uint32(n))
		err = emit(w, vm.PUSHDATA4, buf)
	}
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func emitSyscall(w *bytes.Buffer, api string) error {
	if len(api) == 0 {
		return errors.New("syscall api cannot be of length 0")
	}
	buf := make([]byte, len(api)+1)
	buf[0] = byte(len(api))
	copy(buf[1:len(buf)], []byte(api))
	return emit(w, vm.SYSCALL, buf)
}

func emitCall(w *bytes.Buffer, instr vm.Instruction, label int16) error {
	return emitJmp(w, instr, label)
}

func emitJmp(w *bytes.Buffer, instr vm.Instruction, label int16) error {
	if !isInstrJmp(instr) {
		// TODO: Generate stringer for the instructions so we can use %s in formats.
		return fmt.Errorf("instruction %v is not a jump or call type", instr)
	}
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(label))
	return emit(w, instr, buf)
}

func isInstrJmp(instr vm.Instruction) bool {
	if instr == vm.JMP || instr == vm.JMPIFNOT || instr == vm.JMPIF || instr == vm.CALL {
		return true
	}
	return false
}

func arrayReverse(b []byte) []byte {
	// Protect from big.Ints that have 1 len bytes.
	if len(b) < 2 {
		return b
	}
	dest := make([]byte, len(b))
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		dest[i], dest[j] = b[j], b[i]
	}
	return dest
}
