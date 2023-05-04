package o55

import "math"

//剑指 Offer 55 - II. 平衡二叉树
//输入一棵二叉树的根节点，判断该树是不是平衡二叉树。
//如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
//
//
//示例 1:
//
//给定二叉树 [3,9,20,null,null,15,7]
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回 true 。
//
//示例 2:
//
//给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//       1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
//返回false 。
//
//
//限制：
//
//0 <= 树的结点个数 <= 10000
//注意：本题与主站 110题相同：https://leetcode-cn.com/problems/balanced-binary-tree/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/9hzffg/

// 执行用时： 4 ms, 在所有 Go 提交中击败了 93.52% 的用户
// 内存消耗： 5.7 MB, 在所有 Go 提交中击败了 14.10% 的用户
func isBalanced(root *TreeNode) bool {
	return recur(root) != -1
}

func recur(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := recur(node.Left)
	if left == -1 {
		return -1
	}
	right := recur(node.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(float64(left-right)) < 2 {
		return int(math.Max(float64(left), float64(right))) + 1
	} else {
		return -1
	}
}
