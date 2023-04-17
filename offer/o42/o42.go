package o42

import (
	"math"
)

//剑指 Offer 42. 连续子数组的最大和
//输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
//
//要求时间复杂度为O(n)。
//
//
//示例1:
//
//输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释:连续子数组[4,-1,2,1] 的和最大，为6。
//
//提示：
//
//1 <=arr.length <= 10^5
//-100 <= arr[i] <= 100
//注意：本题与主站 53 题相同：https://leetcode-cn.com/problems/maximum-subarray/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/59gq9c/

// 执行用时： 20 ms, 在所有 Go 提交中击败了 61.73% 的用户
// 内存消耗： 7 MB, 在所有 Go 提交中击败了 31.13% 的用户
func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += int(math.Max(float64(nums[i-1]), 0))
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}
