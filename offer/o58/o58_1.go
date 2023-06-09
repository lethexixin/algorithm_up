package o58

import "strings"

//剑指 Offer 58 - I. 翻转单词顺序
//输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。
//例如输入字符串"I am a student. "，则输出"student. a am I"。
//
//
//
//示例 1：
//
//输入: "the sky is blue"
//输出: "blue is sky the"
//示例 2：
//
//输入: " hello world! "
//输出: "world! hello"
//解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//示例 3：
//
//输入: "a good example"
//输出: "example good a"
//解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
//
//
//说明：
//
//无空格字符构成一个单词。
//输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
//注意：本题与主站 151 题相同：https://leetcode-cn.com/problems/reverse-words-in-a-string/
//
//注意：此题对比原题有改动
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/586ecg/

// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 3.6 MB, 在所有 Go 提交中击败了 64.36% 的用户
func reverseWords(s string) string {
	s = strings.TrimSpace(s)

	i, j := len(s)-1, len(s)-1
	newS := strings.Builder{}
	for i >= 0 {
		for i >= 0 && s[i] != ' ' {
			i--
		}
		newS.WriteString(s[i+1:j+1] + " ")
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
	}
	return strings.TrimSpace(newS.String())
}

// reverseWords2: 同reverseWords
func reverseWords2(s string) string {
	s = strings.TrimSpace(s)
	newStr := strings.Builder{}

	i := len(s) - 1
	for i >= 0 {
		for i >= 0 && s[i] != ' ' {
			i--
		}
		newStr.WriteString(s[i+1:] + " ")
		for i >= 0 && s[i] == ' ' {
			i--
		}
		//reverseWords 中使用 j 记录最后的下标, 而 reverseWords2 这里选择直接截断已经翻转过的数据
		s = s[:i+1]
	}
	return strings.TrimSpace(newStr.String())
}
