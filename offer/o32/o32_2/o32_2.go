package o32_2

import "container/list"

//剑指 Offer 32 - II. 从上到下打印二叉树 II
//从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
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
//返回其层次遍历结果：
//
//[
//  [3],
//  [9,20],
//  [15,7]
//]
//
//提示：
//
//节点总数 <= 1000
//注意：本题与主站 102 题相同：https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5vawr3/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 执行用时： 4 ms , 在所有 Go 提交中击败了 38.55% 的用户
// 内存消耗： 2.8 MB , 在所有 Go 提交中击败了 10.30% 的用户

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	l := list.New()
	l.PushBack(root)
	v := make([][]int, 0)

	for l.Len() != 0 {
		t := make([]int, 0)
		length := l.Len()
		for i := 0; i < length; i++ {
			node := l.Remove(l.Front())
			t = append(t, node.(*TreeNode).Val)
			if node.(*TreeNode).Left != nil {
				l.PushBack(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				l.PushBack(node.(*TreeNode).Right)
			}
		}
		v = append(v, t)
	}
	return v
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 2.6 MB , 在所有 Go 提交中击败了 58.88% 的用户

func LevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	res := make([][]int, 0)

	//建议用 LevelOrder3 更好理解
	for len(queue) != 0 {
		tmp := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			tmp = append(tmp, cur.Val)
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, tmp)
	}

	return res
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 2.6 MB , 在所有 Go 提交中击败了 35.41% 的用户

func LevelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	res := make([][]int, 0)

	for len(queue) != 0 {
		tmp := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			tmp = append(tmp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
		res = append(res, tmp)
	}

	return res
}
