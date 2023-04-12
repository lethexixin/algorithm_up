package o35

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	n0 := &Node{Val: 7}
	n1 := &Node{Val: 13}
	n2 := &Node{Val: 11}
	n3 := &Node{Val: 10}
	n4 := &Node{Val: 1}

	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	n0.Random = nil
	n1.Random = n0
	n2.Random = n4
	n3.Random = n2
	n4.Random = n0

	list := copyRandomList(n0)
	for list != nil {
		if list.Random != nil {
			fmt.Println(list.Val, list.Random.Val)
		} else {
			fmt.Println(list.Val, "nil")
		}

		list = list.Next
	}
}
