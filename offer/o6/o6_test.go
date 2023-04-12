package o6

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	fmt.Println(reversePrint(h))
}
