package o55

import "math"

//剑指 Offer 55 - I. 二叉树的深度
//输入一棵二叉树的根节点，求该树的深度。
//从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
//
//例如：
//
//给定二叉树 [3,9,20,null,null,15,7]，
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最大深度3 。
//
//
//
//提示：
//
//节点总数 <= 10000
//注意：本题与主站 104题相同：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/9hgr5i/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 执行用时： 4 ms , 在所有 Go 提交中击败了 81.10% 的用户
// 内存消耗： 4.1 MB , 在所有 Go 提交中击败了 11.76% 的用户
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 4.1 MB , 在所有 Go 提交中击败了 26.64% 的用户
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := 0

	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res++
		queue = queue[length:]
	}

	return res
}
