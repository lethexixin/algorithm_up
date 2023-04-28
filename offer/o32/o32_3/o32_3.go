package o32_3

import "container/list"

//剑指 Offer 32 - III. 从上到下打印二叉树 III
//请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，
//第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
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
//  [20,9],
//  [15,7]
//]
//
//提示：
//
//节点总数 <= 1000
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5vnp91/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 执行用时： 4 ms , 在所有 Go 提交中击败了 38.11% 的用户
// 内存消耗： 2.8 MB , 在所有 Go 提交中击败了 20.05% 的用户

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	l := list.New()
	l.PushBack(root)

	v := make([][]int, 0)

	for l.Len() != 0 {
		t := make([]int, 0)
		lg := l.Len()
		for i := 0; i < lg; i++ {
			node := l.Remove(l.Front()).(*TreeNode)
			t = append(t, node.Val)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
		if len(v)%2 == 0 {
			v = append(v, t)
		} else {
			tm := make([]int, 0)
			for i := len(t) - 1; i >= 0; i-- {
				tm = append(tm, t[i])
			}
			v = append(v, tm)
		}
	}

	return v
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 2.6 MB , 在所有 Go 提交中击败了 65.74% 的用户

func LevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

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

		if len(res)&1 == 1 {
			for i := 0; i < len(tmp)/2; i++ {
				tmp[i], tmp[len(tmp)-i-1] = tmp[len(tmp)-i-1], tmp[i]
			}
		}
		res = append(res, tmp)

	}

	return res
}
