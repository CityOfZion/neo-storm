package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	s := NewStack(1024)
	s.Push(NewStackItem(1))
	assert.Equal(t, s.Len(), 1)
	assert.Equal(t, s.Pop(), NewStackItem(1))
	assert.Equal(t, s.Len(), 0)
}

func TestStackPeek(t *testing.T) {
	s := NewStack(1024)
	s.Push(NewStackItem(1))
	assert.Equal(t, s.Len(), 1)
	assert.Equal(t, s.Peek(), NewStackItem(1))
	assert.Equal(t, s.Len(), 1)
}

func TestStackSwap(t *testing.T) {
	s := NewStack(1024)
	s.Push(NewStackItem(1))
	s.Push(NewStackItem(2))
	s.Push(NewStackItem(3))
	s.Push(NewStackItem(4))

	s.Swap(4)
	assert.Equal(t, s.Peek(), NewStackItem(1))
	s.Swap(3)
	assert.Equal(t, s.Peek(), NewStackItem(2))
}

func TestPeekN(t *testing.T) {
	s := NewStack(1024)
	s.Push(NewStackItem(1))
	s.Push(NewStackItem(2))
	s.Push(NewStackItem(3))

	assert.Equal(t, s.PeekN(1), NewStackItem(2))
	assert.Equal(t, s.PeekN(2), NewStackItem(1))
}

func TestInsertAt(t *testing.T) {
	s := NewStack(1024)
	s.Push(NewStackItem(1))
	s.Push(NewStackItem(3))
	s.Push(NewStackItem(4))

	s.InsertAt(NewStackItem(2), 2)
	assert.Equal(t, s.Pop(), NewStackItem(4))
	assert.Equal(t, s.Pop(), NewStackItem(3))
	assert.Equal(t, s.Pop(), NewStackItem(2))
	assert.Equal(t, s.Pop(), NewStackItem(1))
}

func TestRemoveAt(t *testing.T) {
	s := NewStack(1024)
	s.Push(NewStackItem(1))
	s.Push(NewStackItem(2))
	s.Push(NewStackItem(3))
	s.Push(NewStackItem(4))
	s.Push(NewStackItem(5))
	s.Push(NewStackItem(6))

	item := s.RemoveAt(1)
	assert.Equal(t, item, NewStackItem(5))
}
