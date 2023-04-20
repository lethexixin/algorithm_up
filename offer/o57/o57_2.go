package o57

//剑指 Offer 57 - II. 和为 s 的连续正数序列
//输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//
//序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
//
//
//示例 1：
//
//输入：target = 9
//输出：[[2,3,4],[4,5]]
//示例 2：
//
//输入：target = 15
//输出：[[1,2,3,4,5],[4,5,6],[7,8]]
//
//限制：
//
//1 <= target <= 10^5
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/eufzm7/

// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 2.1 MB, 在所有 Go 提交中击败了 46.23% 的用户
func findContinuousSequence(target int) [][]int {
	i, j, sum := 1, 2, 3

	//滑动窗口
	res := make([][]int, 0)
	for i < j {
		if sum == target {
			tmp := make([]int, 0)
			for x := i; x <= j; x++ {
				tmp = append(tmp, x)
			}
			res = append(res, tmp)
		}
		if sum >= target {
			sum -= i
			i++
		} else {
			j++
			sum += j
		}
	}
	return res
}
