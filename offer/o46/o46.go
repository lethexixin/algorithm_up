package o46

//剑指 Offer 46. 把数字翻译成字符串
//给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，
//11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。
//请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
//
//
//示例 1:
//
//输入: 12258
//输出: 5
//解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
//
//提示：
//
//0 <= num < 231
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/99wd55/

// todo 有点难
// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 1.9 MB, 在所有 Go 提交中击败了 12.03% 的用户
func translateNum(num int) int {
	a, b, x, y := 1, 1, num%10, num%10
	for num > 9 {
		num /= 10
		x = num % 10
		tmp := 10*x + y
		c := 0
		if tmp >= 10 && tmp <= 25 {
			c = a + b
		} else {
			c = a
		}
		b = a
		a = c
		y = x
	}
	return a
}
