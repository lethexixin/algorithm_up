package o50

//剑指 Offer 50. 第一个只出现一次的字符
//在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
//
//示例 1:
//
//输入：s = "abaccdeff"
//输出：'b'
//示例 2:
//
//输入：s = ""
//输出：' '
//
//限制：
//
//0 <= s 的长度 <= 50000
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5viisg/

// 执行用时： 8 ms, 在所有 Go 提交中击败了 74.05% 的用户
// 内存消耗： 5.2 MB, 在所有 Go 提交中击败了 45.19% 的用户
func firstUniqChar(s string) byte {
	m := make([]int, 128)
	for _, v := range s {
		m[v]++
	}
	for _, v := range s {
		if m[v] == 1 {
			return byte(v)
		}
	}
	return byte(' ')
}
