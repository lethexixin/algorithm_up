package o47

import "math"

//剑指 Offer 47. 礼物的最大价值
//在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
//你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
//给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
//
//
//示例 1:
//
//输入:
//[
//[1,3,1],
//[1,5,1],
//[4,2,1]
//]
//输出: 12
//解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
//
//提示：
//
//0 < grid.length <= 200
//0 < grid[0].length <= 200
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5vokvr/

// 执行用时： 12 ms, 在所有 Go 提交中击败了 9.90% 的用户
// 内存消耗： 3.7 MB, 在所有 Go 提交中击败了 84.12% 的用户
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 && j != 0 {
				grid[i][j] += grid[i][j-1]
			}
			if i != 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			}
			if i != 0 && j != 0 {
				grid[i][j] += int(math.Max(float64(grid[i-1][j]), float64(grid[i][j-1])))
			}
		}
	}
	return grid[m-1][n-1]
}
