package o18

import "testing"

func TestF(t *testing.T) {
	n := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  8,
							Next: nil,
						},
					},
				},
			},
		},
	}

	deleteNode(n, 5)
}
