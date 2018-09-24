package vm

import (
	"fmt"
	"log"
)

// State represents the state the VM is currently in.
type State int

// String implements the fmt.Stringer interface.
func (s State) String() string {
	switch s {
	case StateRunning:
		return "running"
	case StateBreak:
		return "break"
	case StateHalt:
		return "halt"
	case StateFault:
		return "fault"
	default:
		return "unknown"
	}
}

// List of viable VM state constants.
const (
	StateRunning State = iota
	StateBreak
	StateFault
	StateHalt
)

// VM represents the NEO virtual machine that is compatible with .avm bytecode.
type VM struct {
	estack *Stack // evaluation stack
	astack *Stack // alt stack
	istack *Stack // invocation stack

	state State
}

// NewVM returns a pointer to a newly created VM.
func NewVM() *VM {
	return &VM{
		estack: NewStack(1024),
		astack: NewStack(1024),
		istack: NewStack(1024),
	}
}

// Run executes the given script.
func (vm *VM) Run(script []byte) {
	vm.istack.Push(NewContext(script))
	vm.state = StateRunning
	for {
		switch vm.state {
		case StateRunning:
			vm.Step()
		case StateBreak, StateFault, StateHalt:
			fmt.Printf("VM is stopped in state %v", vm.state)
			return
		}
	}
}

func (vm *VM) context() *Context {
	if vm.istack.Len() == 0 {
		return nil
	}
	ctx, ok := vm.istack.Peek().(*Context)
	if !ok {
		panic("Expected to peek (*Context)")
	}
	return ctx
}

// Step advances the stack pointer by one at the time.
func (vm *VM) Step() {
	ctx := vm.context()
	instr := ctx.NextInstruction()
	log.Println(instr)
	// be out of gas.
	vm.exec(ctx, instr)
}

func (vm *VM) exec(ctx *Context, instr Instruction) {
	// Catch all panics occured during VM execution.
	defer func() {
		if err := recover(); err != nil {
			log.Printf("error encountered at instruction %d", ctx.ip)
		}
	}()

	if instr >= PUSHBYTES1 && instr <= PUSHBYTES75 {
		b := ctx.readBytes(int(instr))
		vm.estack.PushVal(b)
		return
	}

	switch instr {
	case PUSHM1, PUSH1, PUSH2, PUSH3, PUSH4,
		PUSH5, PUSH6, PUSH7, PUSH8, PUSH9, PUSH10,
		PUSH11, PUSH12, PUSH13, PUSH14, PUSH15, PUSH16:
		val := int(instr) - int(PUSH1) + 1
		vm.estack.PushVal(val)

	case PUSH0:
		vm.estack.PushVal(0)

	case RET:
		vm.state = StateHalt
	}
}
