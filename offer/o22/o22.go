package o22

//剑指 Offer 22. 链表中倒数第 k 个节点
//输入一个链表，输出该链表中倒数第k个节点。
//为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//
//例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。
//这个链表的倒数第 3 个节点是值为 4 的节点。
//
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 k = 2.
//
//返回链表 4->5.
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/58tl52/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 执行用时： 4 ms , 在所有 Go 提交中击败了 26.29% 的用户
// 内存消耗： 2.1 MB , 在所有 Go 提交中击败了 100.00% 的用户
func getKthFromEnd(head *ListNode, k int) *ListNode {
	f, l := head, head
	for i := 0; i < k; i++ {
		f = f.Next
	}
	for f != nil {
		f = f.Next
		l = l.Next
	}
	return l
}
