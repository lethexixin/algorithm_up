package o54

//剑指 Offer 54. 二叉搜索树的第 k 大节点
//给定一棵二叉搜索树，请找出其中第 k 大的节点的值。
//
//示例 1:
//
//输入: root = [3,1,4,null,2], k = 1
//     3
//    / \
//   1   4
//    \
//    2
//输出: 4
//示例 2:
//
//输入: root = [5,3,6,2,4,null,null,1], k = 3
//         5
//        / \
//       3   6
//      / \
//     2   4
//    /
//   1
//输出: 4
//
//
//限制：
//
//1 ≤ k ≤ 二叉搜索树元素个数
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/58df23/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 执行用时： 8 ms , 在所有 Go 提交中击败了 84.44% 的用户
// 内存消耗： 6.1 MB , 在所有 Go 提交中击败了 61.84% 的用户
func kthLargest(root *TreeNode, k int) int {
	res := 0
	dfs(root, &k, &res)
	return res
}

func dfs(root *TreeNode, k *int, res *int) {
	if root == nil {
		return
	}

	dfs(root.Right, k, res)
	*k--
	//fmt.Println(root.Val, *k)
	if *k == 0 {
		*res = root.Val
		return
	}

	dfs(root.Left, k, res)
}
