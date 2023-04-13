package o53

//剑指 Offer 53 - I. 在排序数组中查找数字 I
//统计一个数字在排序数组中出现的次数。
//
//
//示例 1:
//
//输入: nums = [5,7,7,8,8,10], target = 8
//输出: 2
//示例2:
//
//输入: nums = [5,7,7,8,8,10], target = 6
//输出: 0
//
//提示：
//
//0 <= nums.length <= 105
//-109<= nums[i]<= 109
//nums是一个非递减数组
//-109<= target<= 109
//
//注意：本题与主站 34 题相同（仅返回值不同）：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5874p1/

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 3.8 MB , 在所有 Go 提交中击败了 35.12% 的用户
func search(nums []int, target int) int {
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return 0
	}

	return helper(nums, target) - helper(nums, target-1)
}

func helper(nums []int, tar int) int {
	i, j := 0, len(nums)-1

	for i <= j {
		m := (i + j) / 2
		if nums[m] > tar {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return i
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 3.8 MB , 在所有 Go 提交中击败了 56.83% 的用户
func search2(nums []int, target int) int {
	ans := 0
	for _, v := range nums {
		if v == target {
			ans++
			continue
		}
		if ans > 0 {
			return ans
		}
	}
	return ans
}
