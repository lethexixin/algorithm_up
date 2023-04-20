package o20

//剑指 Offer 20. 表示数值的字符串
//请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
//
//数值（按顺序）可以分成以下几个部分：
//
//若干空格
//一个小数或者整数
//（可选）一个'e'或'E'，后面跟着一个整数
//若干空格
//小数（按顺序）可以分成以下几个部分：
//
//（可选）一个符号字符（'+' 或 '-'）
//下述格式之一：
//至少一位数字，后面跟着一个点 '.'
//至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
//一个点 '.' ，后面跟着至少一位数字
//整数（按顺序）可以分成以下几个部分：
//
//（可选）一个符号字符（'+' 或 '-'）
//至少一位数字
//部分数值列举如下：
//
//["+100", "5e2", "-123", "3.1416", "-1E-16", "0123"]
//部分非数值列举如下：
//
//["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]
//
//
//示例 1：
//
//输入：s = "0"
//输出：true
//示例 2：
//
//输入：s = "e"
//输出：false
//示例 3：
//
//输入：s = "."
//输出：false
//示例 4：
//
//输入：s = ".1"
//输出：true
//
//
//提示：
//
//1 <= s.length <= 20
//s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，空格 ' ' 或者点 '.' 。
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5d6vi6/

// 执行用时： 4 ms, 在所有 Go 提交中击败了 53.36% 的用户
// 内存消耗： 4.4 MB, 在所有 Go 提交中击败了 30.25% 的用户
func isNumber(s string) bool {
	var states []map[byte]int
	states = append(states, map[byte]int{' ': 0, 's': 1, 'd': 2, '.': 4})
	states = append(states, map[byte]int{'d': 2, '.': 4})
	states = append(states, map[byte]int{'d': 2, '.': 3, 'e': 5, ' ': 8})
	states = append(states, map[byte]int{'d': 3, 'e': 5, ' ': 8})
	states = append(states, map[byte]int{'d': 3})
	states = append(states, map[byte]int{'s': 6, 'd': 7})
	states = append(states, map[byte]int{'d': 7})
	states = append(states, map[byte]int{'d': 7, ' ': 8})
	states = append(states, map[byte]int{' ': 8})

	p := 0
	var t byte
	for _, c := range s {
		if c >= '0' && c <= '9' {
			t = 'd'
		} else if c == '+' || c == '-' {
			t = 's'
		} else if c == 'e' || c == 'E' {
			t = 'e'
		} else if c == '.' || c == ' ' {
			t = byte(c)
		} else {
			t = '?'
		}
		if _, ok := states[p][t]; !ok {
			return false
		}
		p = states[p][t]
	}
	return p == 2 || p == 3 || p == 7 || p == 8
}
