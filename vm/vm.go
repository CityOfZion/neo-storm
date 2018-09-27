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
	vm.istack.PushVal(NewContext(script))
	vm.state = StateRunning
	for {
		switch vm.state {
		case StateRunning:
			vm.Step()
		case StateBreak, StateFault, StateHalt:
			fmt.Printf("VM is stopped in state %v\n", vm.state)
			return
		}
	}
}

func (vm *VM) context() *Context {
	if vm.istack.Len() == 0 {
		return nil
	}
	ctx, ok := vm.istack.Peek().value.(*Context)
	if !ok {
		panic("Expected to peek (*Context)")
	}
	return ctx
}

// Step advances the stack pointer by one at the time.
func (vm *VM) Step() {
	ctx := vm.context()
	instr := ctx.NextInstruction()
	vm.exec(ctx, instr)
}

func (vm *VM) exec(ctx *Context, instr Instruction) {
	// Catch all panics occured during VM execution.
	defer func() {
		if err := recover(); err != nil {
			log.Printf("error encountered at instruction %s at instruction pointer %d => %s", instr, ctx.ip, err)
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

	case PUSHDATA1:
		b := ctx.readVarBytes()
		vm.estack.PushVal(b)

	case PUSHDATA2:
		n := ctx.readUint16()
		b := ctx.readBytes(int(n))
		vm.estack.PushVal(b)

	case PUSHDATA4:
		n := ctx.readUint32()
		b := ctx.readBytes(int(n))
		vm.estack.PushVal(b)

	case TOALTSTACK:
		vm.astack.Push(vm.estack.Pop())

	case FROMALTSTACK:
		vm.estack.Push(vm.astack.Pop())

	case DUPFROMALTSTACK:
		vm.estack.Push(vm.astack.Dup())

	case DUP:
		vm.estack.Push(vm.estack.Dup())

	case SWAP:
		a := vm.estack.Pop()
		b := vm.estack.Pop()
		vm.estack.Push(a)
		vm.estack.Push(b)

	case XSWAP:
		n := int(vm.estack.Pop().BigInt().Int64())
		if n <= 0 {
			panic("XSWAP: invalid length")
		}

		a := vm.estack.PeekN(n)
		b := vm.estack.Peek()
		aval := a.value
		bval := b.value
		a.value = bval
		b.value = aval

	case TUCK:
		n := int(vm.estack.Pop().BigInt().Int64())
		if n <= 0 {
			panic("TUCK: invalid length")
		}
		vm.estack.InsertAt(vm.estack.Peek(), n)
		vm.estack.Inspect()

	case ROT:
		c := vm.estack.Pop()
		b := vm.estack.Pop()
		a := vm.estack.Pop()

		vm.estack.Push(b)
		vm.estack.Push(c)
		vm.estack.Push(a)

	case NIP:
		item := vm.estack.Pop()
		_ = vm.estack.Pop()
		vm.estack.Push(item)

	case OVER:
		b := vm.estack.Pop()
		a := vm.estack.Pop()

		vm.estack.Push(b)
		vm.estack.Push(a)

	case ROLL:
		n := int(vm.estack.Pop().BigInt().Int64())
		if n < 0 {
			panic("ROLL: popped negative value from the stack")
		}
		if n == 0 {
			panic("ROLL: cannot roll on index 0")
		}
		vm.estack.Push(vm.estack.RemoveAt(n))

	case ADD:
		a := vm.estack.Pop().BigInt()
		b := vm.estack.Pop().BigInt()

		a.Add(a, b)
		vm.estack.PushVal(a)

	case SUB:
		a := vm.estack.Pop().BigInt()
		b := vm.estack.Pop().BigInt()

		b.Sub(b, a)
		vm.estack.PushVal(b)

	case MUL:
		a := vm.estack.Pop().BigInt()
		b := vm.estack.Pop().BigInt()

		a.Mul(a, b)
		vm.estack.PushVal(a)

	case DIV:
		a := vm.estack.Pop().BigInt()
		b := vm.estack.Pop().BigInt()

		b.Div(b, a)
		vm.estack.PushVal(b)

	case MOD:
		a := vm.estack.Pop().BigInt()
		b := vm.estack.Pop().BigInt()

		b.Mod(b, a)
		vm.estack.PushVal(b)

	case SHL:
		a := vm.estack.Pop().BigInt()
		b := vm.estack.Pop().BigInt()

		b.Lsh(b, uint(a.Int64()))
		vm.estack.PushVal(b)

	case SHR:
		a := vm.estack.Pop().BigInt()
		b := vm.estack.Pop().BigInt()

		b.Rsh(b, uint(a.Int64()))
		vm.estack.PushVal(b)

	case RET:
		vm.state = StateHalt
	}
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("[STORM VM] ")
}
