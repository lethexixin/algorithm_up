package o32_1

import "container/list"

//剑指 Offer 32 - I. 从上到下打印二叉树
//从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
//
//
//例如:
//给定二叉树:[3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回：
//
//[3,9,20,15,7]
//
//提示：
//
//节点总数 <= 1000
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/9ackoe/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 执行用时： 4 ms , 在所有 Go 提交中击败了 30.89% 的用户
// 内存消耗： 2.5 MB , 在所有 Go 提交中击败了 37.86% 的用户

func LevelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	v := make([]int, 0)

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		v = append(v, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return v
}

// 执行用时： 4 ms , 在所有 Go 提交中击败了 30.89% 的用户
// 内存消耗： 2.6 MB , 在所有 Go 提交中击败了 11.61% 的用户

func LevelOrder2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	l := list.New()
	l.PushBack(root)
	v := make([]int, 0)

	for l.Len() != 0 {
		node := l.Remove(l.Front())
		v = append(v, node.(*TreeNode).Val)
		if node.(*TreeNode).Left != nil {
			l.PushBack(node.(*TreeNode).Left)
		}
		if node.(*TreeNode).Right != nil {
			l.PushBack(node.(*TreeNode).Right)
		}
	}
	return v
}
