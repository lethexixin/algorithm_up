package o4

//剑指 Offer 04. 二维数组中的查找
//在一个 n * m 的二维数组中，每一行都按照从左到右非递减的顺序排序，每一列都按照从上到下非递减的顺序排序。
//请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
//
//
//示例:
//
//现有矩阵 matrix 如下：
//
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
//给定 target=5，返回true。
//
//给定target=20，返回false。
//
//
//限制：
//
//0 <= n <= 1000
//
//0 <= m <= 1000
//
//
//注意：本题与主站 240 题相同：https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5v76yi/

// 执行用时： 16 ms, 在所有 Go 提交中击败了 95.94% 的用户
// 内存消耗： 6.6 MB, 在所有 Go 提交中击败了 35.55% 的用户
func findNumberIn2DArray(matrix [][]int, target int) bool {
	i, j := len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}
