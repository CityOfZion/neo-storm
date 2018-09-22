package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	s := NewStack(1024)
	s.Push(1)
	assert.Equal(t, s.Len(), 1)
	assert.Equal(t, s.Pop(), 1)
	assert.Equal(t, s.Len(), 0)
}

func TestStackPeek(t *testing.T) {
	s := NewStack(1024)
	s.Push(1)
	assert.Equal(t, s.Len(), 1)
	assert.Equal(t, s.Peek(), 1)
	assert.Equal(t, s.Len(), 1)
}

func TestStackSwap(t *testing.T) {
	s := NewStack(1024)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	s.Swap(4)
	assert.Equal(t, s.Peek(), 1)
	s.Swap(3)
	assert.Equal(t, s.Peek(), 2)
}
