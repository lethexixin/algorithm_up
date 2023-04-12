package o24

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	n := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	list := reverseList(n)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}

	//测试传入空链表
	list2 := reverseList(nil)
	for list2 != nil {
		fmt.Println(list2.Val)
		list2 = list2.Next
	}
}
