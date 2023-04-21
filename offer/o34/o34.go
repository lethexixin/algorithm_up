package o34

//剑指 Offer 34. 二叉树中和为某一值的路径
//给你二叉树的根节点 root 和一个整数目标和 targetSum ，
//找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
//叶子节点 是指没有子节点的节点。
//
//
//
//示例 1：
//					  5
//					/   \
//				  4		  8
//				/		/   \
//			   11     13     4
//			 /	  \        /   \
//			7	   2      5     1
//
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：[[5,4,11,2],[5,8,4,5]]
//
//示例 2：
//          1
//         / \
//        2   3
//输入：root = [1,2,3], targetSum = 5
//输出：[]
//
//示例 3：
//          1
//         /
//        2
//输入：root = [1,2], targetSum = 0
//输出：[]
//
//
//提示：
//
//树中节点总数在范围 [0, 5000] 内
//-1000 <= Node.val <= 1000
//-1000 <= targetSum <= 1000
//注意：本题与主站 113题相同：https://leetcode-cn.com/problems/path-sum-ii/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5dy6pt/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 4.4 MB , 在所有 Go 提交中击败了 51.63% 的用户
func pathSum(root *TreeNode, target int) [][]int {
	stack := make([]int, 0)
	res := make([][]int, 0)
	recur(root, target, stack, &res)
	return res
}

func recur(root *TreeNode, tar int, stack []int, res *[][]int) {
	if root == nil {
		return
	}
	stack = append(stack, root.Val)
	tar -= root.Val
	if tar == 0 && root.Left == nil && root.Right == nil {
		//slice是一个指向底层的数组的指针结构体
		//因为是先序遍历，如果 root.Right != nil ,arr 切片底层的数组会被修改
		//所以这里需要 copy stack 到 tmp，再添加进 res，防止 stack 底层数据修改带来的错误
		tmp := make([]int, len(stack))
		copy(tmp, stack)
		*res = append(*res, tmp)

		//或者使用下面的方式
		//*res = append(*res, append([]int{}, stack...))
	}
	recur(root.Left, tar, stack, res)
	recur(root.Right, tar, stack, res)
	stack = stack[:len(stack)-1]
}
