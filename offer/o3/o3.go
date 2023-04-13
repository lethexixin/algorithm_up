package o3

//剑指 Offer 03. 数组中重复的数字
//找出数组中重复的数字。
//
//
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，
//但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
//
//示例 1：
//
//输入：
//[2, 3, 1, 0, 2, 5, 3]
//输出：2 或 3
//
//限制：
//
//2 <= n <= 100000
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/59bjss/

// 执行用时： 28 ms , 在所有 Go 提交中击败了 89.67% 的用户
// 内存消耗： 7.7 MB , 在所有 Go 提交中击败了 75.70% 的用户
func findRepeatNumber(nums []int) int {
	m := make([]int, len(nums))

	for _, num := range nums {
		if m[num] == 1 {
			return num
		}
		m[num]++
	}

	return nums[0]
}

// 执行用时： 36 ms , 在所有 Go 提交中击败了 37.88% 的用户
// 内存消耗： 10.9 MB , 在所有 Go 提交中击败了 5.38% 的用户
func findRepeatNumber2(nums []int) int {
	m := make(map[int]any)

	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		}
		m[num] = nil
	}

	return nums[0]
}
