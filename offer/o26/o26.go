package o26

//剑指 Offer 26. 树的子结构
//输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
//
//B是A的子结构， 即 A中有出现和B相同的结构和节点值。
//
//例如:
//给定的树 A:
//
// 			 3
//          / \
//		   4   5
//        / \
//       1   2
//给定的树 B：
//
//		4
//	   /
//    1
//返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
//
//示例 1：
//
//输入：A = [1,2,3], B = [3,1]
//输出：false
//示例 2：
//
//输入：A = [3,4,5,1,2], B = [4,1]
//输出：true
//限制：
//
//0 <= 节点个数 <= 10000
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5dshwe/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// todo 有点难,需要重新理解
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return isSub(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSub(A, B *TreeNode) bool {
	if A == nil || A.Val != B.Val {
		return false
	}
	// if B.Left == nil && B.Right == nil {
	//     return true
	// }
	ans := true
	if B.Left != nil {
		ans = ans && isSub(A.Left, B.Left)
	}
	if B.Right != nil {
		ans = ans && isSub(A.Right, B.Right)
	}
	//fmt.Println(ans)
	return ans
}
