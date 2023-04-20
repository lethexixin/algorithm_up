package o56

//剑指 Offer 56 - I. 数组中数字出现的次数
//一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。
//请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
//
//
//
//示例 1：
//
//输入：nums = [4,1,4,6]
//输出：[1,6] 或 [6,1]
//示例 2：
//
//输入：nums = [1,2,10,4,1,4,3,3]
//输出：[2,10] 或 [10,2]
//
//
//限制：
//
//2 <= nums.length <= 10000
//
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/eubbnm/

// 执行用时： 28 ms, 在所有 Go 提交中击败了 9.89% 的用户
// 内存消耗： 6.1 MB, 在所有 Go 提交中击败了 99.54% 的用户
func singleNumbers(nums []int) []int {
	x, y, n, m := 0, 0, 0, 1
	for _, num := range nums {
		n ^= num
	}

	for n&m == 0 {
		m <<= 1
	}

	for _, num := range nums {
		if num&m == 0 {
			x ^= num
		} else {
			y ^= num
		}
	}

	return []int{x, y}
}
