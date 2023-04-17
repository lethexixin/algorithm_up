package o25

//剑指 Offer 25. 合并两个排序的链表
//输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//
//示例1：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//限制：
//
//0 <= 链表长度 <= 1000
//
//注意：本题与主站 21 题相同：https://leetcode-cn.com/problems/merge-two-sorted-lists/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5vq98s/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 执行用时： 8 ms , 在所有 Go 提交中击败了 49.14% 的用户
// 内存消耗： 3.9 MB , 在所有 Go 提交中击败了 82.87% 的用户
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := &ListNode{}
	cur := tmp

	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return tmp.Next
}
