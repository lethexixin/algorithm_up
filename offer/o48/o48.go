package o48

import "math"

//剑指 Offer 48. 最长不含重复字符的子字符串
//请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
//
//
//示例1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
//    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
//
//
//提示：
//
//s.length <= 40000
//注意：本题与主站 3 题相同：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5dgr0c/

// 执行用时： 8 ms, 在所有 Go 提交中击败了 65.68% 的用户
// 内存消耗： 3.9 MB, 在所有 Go 提交中击败了 8.67% 的用户
func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	current, max := 0, 0

	for j, v := range s {
		i, ok := m[string(v)]
		if !ok {
			i = -1
		}
		m[string(v)] = j
		if current >= j-i {
			current = j - i
		} else {
			current++
		}
		max = int(math.Max(float64(current), float64(max)))
	}

	return max
}
