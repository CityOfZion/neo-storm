package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	s := NewStack(1024)
	s.Push(NewItem(1))
	assert.Equal(t, s.Len(), 1)
	assert.Equal(t, s.Pop(), NewItem(1))
	assert.Equal(t, s.Len(), 0)
}

func TestStackPeek(t *testing.T) {
	s := NewStack(1024)
	s.Push(NewItem(1))
	assert.Equal(t, s.Len(), 1)
	assert.Equal(t, s.Peek(), NewItem(1))
	assert.Equal(t, s.Len(), 1)
}

func TestStackSwap(t *testing.T) {
	s := NewStack(1024)
	s.Push(NewItem(1))
	s.Push(NewItem(2))
	s.Push(NewItem(3))
	s.Push(NewItem(4))

	s.Swap(4)
	assert.Equal(t, s.Peek(), NewItem(1))
	s.Swap(3)
	assert.Equal(t, s.Peek(), NewItem(2))
}
