package o21

//剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
//输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
//使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。
//
//
//示例：
//
//输入：nums =[1,2,3,4]
//输出：[1,3,2,4]
//注：[3,1,2,4] 也是正确的答案之一。
//
//提示：
//
//0 <= nums.length <= 50000
//0 <= nums[i] <= 10000
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5v8a6t/

// 执行用时： 16 ms, 在所有 Go 提交中击败了 89.94% 的用户
// 内存消耗： 6.3 MB, 在所有 Go 提交中击败了 66.63% 的用户
func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[i]%2 == 1 {
			i++
		}
		for i < j && nums[j]%2 == 0 {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}
