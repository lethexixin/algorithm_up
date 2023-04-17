package o57

//剑指 Offer 57. 和为 s 的两个数字
//输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。
//如果有多对数字的和等于s，则输出任意一对即可。
//
//
//示例 1：
//
//输入：nums = [2,7,11,15], target = 9
//输出：[2,7] 或者 [7,2]
//示例 2：
//
//输入：nums = [10,26,30,31,47,60], target = 40
//输出：[10,30] 或者 [30,10]
//
//限制：
//
//1 <= nums.length <= 10^5
//1 <= nums[i]<= 10^6
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5832fi/

// 执行用时： 140 ms, 在所有 Go 提交中击败了 75.40% 的用户
// 内存消耗： 8.5 MB, 在所有 Go 提交中击败了 61.93% 的用户
func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l]+nums[r] < target {
			l++
		}
		for l < r && nums[l]+nums[r] > target {
			r--
		}
		if nums[l]+nums[r] == target {
			return []int{nums[l], nums[r]}
		}
	}
	return nil
}
