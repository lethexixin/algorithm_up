package o28

//剑指 Offer 28. 对称的二叉树
//请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
//
//例如，二叉树[1,2,2,3,4,4,3] 是对称的。
//
//			  1
//			/   \
//		   2     2
//		  / \   / \
// 		 3  4  4   3
//但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:
//
//			 1
//			/ \
//		   2   2
//			\   \
//			 3   3
//
//
//示例 1：
//
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//示例 2：
//
//输入：root = [1,2,2,null,3,null,3]
//输出：false

//
//限制：
//
//0 <= 节点个数 <= 1000
//
//注意：本题与主站 101 题相同：https://leetcode-cn.com/problems/symmetric-tree/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5d412v/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 执行用时： 4 ms , 在所有 Go 提交中击败了 50.12% 的用户
// 内存消耗： 2.7 MB , 在所有 Go 提交中击败了 79.18% 的用户
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)
}

func recur(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return recur(left.Right, right.Left) && recur(left.Left, right.Right)
}
