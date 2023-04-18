package o39

//剑指 Offer 39. 数组中出现次数超过一半的数字
//数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
//
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//示例1:
//
//输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
//输出: 2
//
//限制：
//
//1 <= 数组长度 <= 50000
//
//
//注意：本题与主站 169 题相同：https://leetcode-cn.com/problems/majority-element/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/99iy4g/

// 执行用时： 16 ms, 在所有 Go 提交中击败了 67.45% 的用户
// 内存消耗： 5.9 MB, 在所有 Go 提交中击败了 10.22% 的用户
func majorityElement(nums []int) int {
	m := make(map[int]int)

	for _, num := range nums {
		if m[num] >= len(nums)/2 {
			return num
		}
		m[num]++
	}
	return nums[0]
}

//todo 众数法
