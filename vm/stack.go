package vm

// Stack represents a general purpose stack with FIFO as its semantics.
type Stack struct {
	data []*StackItem
	cap  uint
}

// NewStack creates a new stack with the given capacity.
func NewStack(size uint) *Stack {
	return &Stack{
		data: []*StackItem{},
	}
}

// Data returns the underlying data that is currently on the stack.
func (s *Stack) Data() []*StackItem {
	return s.data
}

// Len returns the number of items that are currently on the stack.
func (s *Stack) Len() int {
	return len(s.data)
}

// Push pushes an item on to the stack.
func (s *Stack) Push(item *StackItem) {
	s.data = append(s.data, item)
}

// PushVal pushes the given value and will automaticaly convert it to a stack
// item.
func (s *Stack) PushVal(val interface{}) {
	s.data = append(s.data, NewStackItem(val))
}

// Pop pops an item of the stack.
func (s *Stack) Pop() *StackItem {
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

// Peek peeks the first item that is on the stack without consuming it.
func (s *Stack) Peek() *StackItem {
	return s.data[s.Len()-1]
}

// PeekN will peek the nth item from the top of the stack.
func (s *Stack) PeekN(n int) *StackItem {
	return s.data[s.Len()-(1+n)]
}

// Dup returns a copy of the top stack item.
func (s *Stack) Dup() *StackItem {
	item := s.Peek()
	if item == nil {
		return nil
	}
	return &StackItem{item.value, item.kind}
}

// Swap swaps the n(th) item with the top of the stack.
func (s *Stack) Swap(n int) {
	if n == 0 {
		panic("cannot swap with index 0")
	}
	s.data[s.Len()-n], s.data[s.Len()-1] = s.data[s.Len()-1], s.data[s.Len()-n]
}

// InsertAt inserts the given stack item at the given offset.
func (s *Stack) InsertAt(item *StackItem, n int) {
	if n == 0 {
		s.Push(item)
		return
	}
	if n > s.Len() {
		panic("cannot insert item on index that is longer than the stack itself")
	}

	s.data = append(s.data, &StackItem{})
	n = (len(s.data) - 1) - n
	copy(s.data[n+1:], s.data[n:])
	s.data[n] = item
}

// RemoveAt removes and returns a stack item on the given offset.
func (s *Stack) RemoveAt(n int) *StackItem {
	n = (len(s.data) - 1) - n
	item := s.data[n]
	s.data = append(s.data[:n], s.data[n+1:]...)
	return item
}

// Inspect prints out a human readable representation of all the items that are
// on the stack.
func (s *Stack) Inspect() {
	for _, item := range s.data {
		item.Inspect()
	}
}
