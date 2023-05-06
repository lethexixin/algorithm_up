package o29

// 剑指 Offer 29. 顺时针打印矩阵
// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
//
// 示例 1：
//
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
// 示例 2：
//
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
// 限制：
//
// 0 <= matrix.length <= 100
// 0 <= matrix[i].length <= 100
// 注意：本题与主站 54 题相同：https://leetcode-cn.com/problems/spiral-matrix/
// 作者：Krahets
// 链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5vfh9g/

// 执行用时： 12 ms, 在所有 Go 提交中击败了 52.60% 的用户
// 内存消耗： 6.3 MB, 在所有 Go 提交中击败了 25.43% 的用户
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	// 1  2  3
	// 4  5  6
	// 7  8  9
	res := make([]int, 0)
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
			if i == right {
				top++
			}
		}
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
			if i == bottom {
				right--
			}
		}
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
			if i == left {
				bottom--
			}
		}
		if bottom < top {
			break
		}
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
			if i == top {
				left++
			}
		}
		if left > right {
			break
		}

	}
	return res
}
