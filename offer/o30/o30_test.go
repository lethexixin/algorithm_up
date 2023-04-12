package o30

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.Min())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.Min())
}
