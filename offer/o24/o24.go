package o24

//剑指 Offer 24. 反转链表
//定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
//限制：
//
//0 <= 节点个数 <= 5000
//
//
//注意：本题与主站 206 题相同：https://leetcode-cn.com/problems/reverse-linked-list/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/9pdjbm/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 2.4 MB , 在所有 Go 提交中击败了 85.27% 的用户
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre any

	for cur != nil {
		tmp := cur.Next
		if pre == nil {
			cur.Next = nil
		} else {
			cur.Next = pre.(*ListNode)
		}

		pre = cur
		cur = tmp
	}

	//当 head==nil 时, pre = interface{}
	if v, ok := pre.(*ListNode); ok {
		return v
	}
	return nil
}
