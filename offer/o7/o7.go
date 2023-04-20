package o7

import "fmt"

//剑指 Offer 07. 重建二叉树
//输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
//
//假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
//
//示例 1:
//       3
//     /   \
//    9    20
//        /  \
//       15   7
//
//Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//Output: [3,9,20,null,null,15,7]
//示例 2:
//
//Input: preorder = [-1], inorder = [-1]
//Output: [-1]
//
//限制：
//
//0 <= 节点个数 <= 5000
//
//
//注意：本题与主站 105 题重复：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/99lxci/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 执行用时： 4 ms , 在所有 Go 提交中击败了 92.15% 的用户
// 内存消耗： 4.3 MB , 在所有 Go 提交中击败了 7.98% 的用户
func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}
	return recur(0, 0, len(inorder)-1, preorder, m)
}

func recur(root, left, right int, preorder []int, m map[int]int) *TreeNode {
	if left > right {
		return nil
	}

	node := &TreeNode{Val: preorder[root]}
	i := m[preorder[root]]
	fmt.Println(root, left, right, i, node.Val)
	node.Left = recur(root+1, left, i-1, preorder, m)
	node.Right = recur(root+i-left+1, i+1, right, preorder, m)
	return node
}
