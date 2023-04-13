package o5

import "strings"

//剑指 Offer 05. 替换空格
//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
//
//示例 1：
//
//输入：s = "We are happy."
//输出："We%20are%20happy."
//
//限制：
//
//0 <= s 的长度 <= 10000
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/50ywkd/

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 1.8 MB , 在所有 Go 提交中击败了 72.48% 的用户
func replaceSpace(s string) string {
	ns := strings.Builder{}

	for _, v := range s {
		if v == ' ' {
			ns.WriteString("%20")
		} else {
			ns.WriteString(string(v))
		}
	}
	return ns.String()
}
