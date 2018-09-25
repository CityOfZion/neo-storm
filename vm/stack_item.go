package vm

import (
	"fmt"
	"math/big"
	"reflect"
)

// StackItem represents an item on the stack.
type StackItem struct {
	value interface{}
	kind  StackItemType
}

// BigInt attempts to return the stack item as a big integer type.
func (s *StackItem) BigInt() *big.Int {
	val, ok := s.value.(*big.Int)
	if !ok {
		panic("this stack item is not of type big integer")
	}
	return val
}

// Inspect returns underlying information about an item on the stack.
func (s *StackItem) Inspect() {
	switch s.kind {
	case BigIntType:
		fmt.Printf("<type: %s, value: %d>\n", s.kind, s.value)
	case ByteArrayType:
		fmt.Printf("<type: %s, value: %v string: %s>\n", s.kind, s.value, s.value)
	}
}

// StackITemType represents the underlying type of an item on the stack.
type StackItemType int

// Viable list of stack item types.
const (
	BigIntType StackItemType = iota
	ByteArrayType
	ArrayType
	ContextType
)

// String implements the fmt.Stringer interface.
func (s StackItemType) String() string {
	switch s {
	case BigIntType:
		return "BigInteger"
	case ByteArrayType:
		return "ByteArray"
	case ArrayType:
		return "Array"
	case ContextType:
		return "Context"
	default:
		return "Unknown"
	}
}

// NewStackItem creates a new StackItem from the given value. It will automatically convert
// the given value to the correct stack item type. Will panic if the given value
// is not suitable as a stack item.
func NewStackItem(value interface{}) *StackItem {
	var kind StackItemType

	switch t := value.(type) {
	case *big.Int:
		kind = BigIntType
	case int:
		return &StackItem{big.NewInt(int64(t)), BigIntType}
	case int64:
		return &StackItem{big.NewInt(t), BigIntType}
	case []byte:
		kind = ByteArrayType
	case *Context:
		kind = ContextType
	case []StackItem:
		kind = ArrayType
	default:
		// TODO: be more specific for this error to the end user.
		fmt.Println(reflect.TypeOf(t))
		panic("Invalid value to construct a stack item")
	}

	return &StackItem{value, kind}
}
