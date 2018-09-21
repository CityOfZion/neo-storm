package vm

type StackItem interface{}

type Stack struct {
	data []StackItem
	cap  uint
}

// NewStack creates a new stack with the given capacity.
func NewStack(size uint) *Stack {
	return &Stack{
		data: []StackItem{},
	}
}

// Data returns the underlying data that is currently on the stack.
func (s *Stack) Data() []StackItem {
	return s.data
}

// Len returns the number of items that are currently on the stack.
func (s *Stack) Len() int {
	return len(s.data)
}

// Push pushes an item on to the stack.
func (s *Stack) Push(item StackItem) {
	s.data = append(s.data, item)
}

// Pop pops an item of the stack.
func (s *Stack) Pop() StackItem {
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

// Peek peeks the first item that is on the stack without consuming it.
func (s *Stack) Peek() StackItem {
	return s.data[s.Len()-1]
}
