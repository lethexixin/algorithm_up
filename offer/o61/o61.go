package o61

import "math"

//剑指 Offer 61. 扑克牌中的顺子
//从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。
//2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
//
//
//示例1:
//
//输入: [1,2,3,4,5]
//输出: True
//
//示例2:
//
//输入: [0,0,1,2,5]
//输出: True
//
//限制：
//
//数组长度为 5
//
//数组的数取值为 [0, 13] .
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/57mpoj/

// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 1.8 MB, 在所有 Go 提交中击败了 87.80% 的用户
func isStraight(nums []int) bool {
	m := make(map[int]int)
	min, max := 14, 0
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return false
		}
		if num != 0 {
			m[num]++
			min = int(math.Min(float64(min), float64(num)))
			max = int(math.Max(float64(max), float64(num)))
		}
	}
	if max-min >= 5 {
		return false
	}
	return true
}
