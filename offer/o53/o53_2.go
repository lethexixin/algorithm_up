package o53

//剑指 Offer 53 - II. 0～n-1 中缺失的数字
//一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
//
//
//示例 1:
//
//输入: [0,1,3]
//输出: 2
//示例2:
//
//输入: [0,1,2,3,4,5,6,7,9]
//输出: 8
//
//限制：
//
//1 <= 数组长度 <= 10000
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/58iqo5/

//执行用时： 12 ms, 在所有 Go 提交中击败了 89.45% 的用户
//内存消耗： 6 MB, 在所有 Go 提交中击败了 99.37% 的用户

func missingNumber(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] != m {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return i
}
