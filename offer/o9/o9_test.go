package o9

import (
	"algorithm_up/offer/o9/f1"
	"algorithm_up/offer/o9/f2"
	"fmt"
	"testing"
)

func TestF1(t *testing.T) {
	q := f1.Constructor()
	fmt.Println(q.DeleteHead())
	q.AppendTail(5)
	q.AppendTail(2)
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
}

func TestF2(t *testing.T) {
	q := f2.Constructor()
	fmt.Println(q.DeleteHead())
	q.AppendTail(5)
	q.AppendTail(2)
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
}
