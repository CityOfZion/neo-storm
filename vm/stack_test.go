package vm

import (
	"fmt"
	"log"
	"testing"
)

func TestStackPush(t *testing.T) {
	s := NewStack(1024)
	s.Push(1)
	log.Println(s.data)
	item := s.Peek()
	fmt.Println(item)
}
