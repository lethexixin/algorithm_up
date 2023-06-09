package o52

//剑指 Offer 52. 两个链表的第一个公共节点
//输入两个链表，找出它们的第一个公共节点。
//
//如下面的两个链表：
//         a1 - a2
//                 \
//                  c1 - c2 - c3
//                 /
//    b1 - b2 - b3
//在节点 c1 开始相交。
//
//示例 1：
// A:         4 - 1
//                  \
//                   8 - 4 - 5
//                  /
// B:     5 - 0 - 1
//输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
//输出：Reference of the node with value = 8
//输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，
//链表 A 为 [4,1,8,4,5]， 链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；
//在 B 中，相交节点前有 3 个节点。
//
//
//示例2：
// A:     0 - 9 - 1
//                  \
//                   2 - 4
//                  /
// B:             3
//输入：intersectVal= 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
//输出：Reference of the node with value = 2
//输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，
//链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；
//在 B 中，相交节点前有 1 个节点。
//
//
//示例3：
// A:      2 - 6 - 4
//
// B:         1 - 5
//输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
//输出：null
//输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
//由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
//解释：这两个链表不相交，因此返回 null。
//
//
//注意：
//
//如果两个链表没有交点，返回 null.
//在返回结果后，两个链表仍须保持原有的结构。
//可假定整个链表结构中没有循环。
//程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
//本题与主站 160 题相同：https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/oe5os3/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 执行用时： 20 ms , 在所有 Go 提交中击败了 99.88% 的用户
// 内存消耗： 6.6 MB , 在所有 Go 提交中击败了 95.86% 的用户
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}
