package o6

//剑指 Offer 06. 从尾到头打印链表
//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//
//示例 1：
//
//输入：head = [1,3,2]
//输出：[2,3,1]
//
//限制：
//
//0 <= 链表长度 <= 10000
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5dt66m/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 3 MB , 在所有 Go 提交中击败了 56.78% 的用户
func reversePrint(head *ListNode) []int {
	stack := make([]int, 0)

	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	res := make([]int, len(stack))
	for i := 0; i < len(stack); i++ {
		res[len(stack)-i-1] = stack[i]
	}

	return res
}
