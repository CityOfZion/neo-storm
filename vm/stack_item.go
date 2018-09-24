package vm

// StackItem is an abstraction for data that lives on the stack.
type StackItem interface {
	// Value returns the underlying value that is carried by the stack item.
	Value() interface{}
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

// Item represents an item on the stack.
type Item struct {
	value interface{}
	kind  StackItemType
}

// NewItem creates a new Item from the given value. It will automatically convert
// the given value to the correct stack item type. Will panic if the given value
// is not suitable as a stack item.
func NewItem(value interface{}) *Item {
	var kind StackItemType

	switch value.(type) {
	case int64, int:
		kind = BigIntType
	case []byte:
		kind = ByteArrayType
	case *Context:
		kind = ContextType
	case []StackItem:
		kind = ArrayType
	default:
		// TODO: be more specific for this error to the end user.
		panic("Invalid value to construct a stack item")
	}

	return &Item{value, kind}
}

// Value implements the StackItem interface.
func (i *Item) Value() interface{} {
	return i.value
}
