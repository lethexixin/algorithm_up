package o56

//剑指 Offer 56 - II. 数组中数字出现的次数 II
//在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
//
//
//示例 1：
//
//输入：nums = [3,4,3,3]
//输出：4
//示例 2：
//
//输入：nums = [9,1,7,9,7,9,7]
//输出：1
//
//限制：
//
//1 <= nums.length <= 10000
//1 <= nums[i] < 2^31
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/9hyq1r/

// 执行用时： 32 ms, 在所有 Go 提交中击败了 23.76% 的用户
// 内存消耗： 6.3 MB, 在所有 Go 提交中击败了 91.71% 的用户
func singleNumber(nums []int) int {
	o, t := 0, 0
	for _, num := range nums {
		o = o ^ num & ^t
		t = t ^ num & ^o
	}
	return o
}
