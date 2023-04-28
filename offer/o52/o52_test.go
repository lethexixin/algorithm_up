package o52

import "testing"

func TestF(t *testing.T) {
	headA := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 8,
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

	headB := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}

	getIntersectionNode(headA, headB)

	headA = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}
	headB = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	}

	getIntersectionNode(headA, headB)
}
