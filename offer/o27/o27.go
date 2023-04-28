package o27

import "container/list"

//剑指 Offer 27. 二叉树的镜像
//请完成一个函数，输入一个二叉树，该函数输出它的镜像。
//
//例如输入：
//
//				 4
//			  /    \
//			 2      7
//			/ \    / \
//		   1   3  6   9
//镜像输出：
//
// 				 4
//			   /   \
//			  7     2
//			 / \   / \
//			9   6 3   1
//
//
//示例 1：
//
//输入：root = [4,2,7,1,3,6,9]
//输出：[4,7,2,9,6,3,1]
//
//限制：
//
//0 <= 节点个数 <= 1000
//
//注意：本题与主站 226 题相同：https://leetcode-cn.com/problems/invert-binary-tree/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/59zt5i/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 1.9 MB , 在所有 Go 提交中击败了 46.57% 的用户
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l := list.New()
	l.PushBack(root)

	for l.Len() != 0 {
		node := l.Remove(l.Front()).(*TreeNode)
		tmp := node.Left
		if node.Left != nil {
			l.PushBack(node.Left)
		}
		if node.Right != nil {
			l.PushBack(node.Right)
		}
		node.Left = node.Right
		node.Right = tmp
	}
	return root
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 1.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
func mirrorTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := root.Left
	root.Left = mirrorTree2(root.Right)
	root.Right = mirrorTree2(left)
	return root
}
