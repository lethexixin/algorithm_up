package o65

import "fmt"

//剑指 Offer 65. 不用加减乘除做加法
//写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
//
//
//
//示例:
//
//输入: a = 1, b = 1
//输出: 2
//
//
//提示：
//
//a,b均可能是负数或 0
//结果不会溢出 32 位整数
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5vz6d1/

// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 1.9 MB, 在所有 Go 提交中击败了 8.31% 的用户
func add(a int, b int) int {
	for b != 0 {
		c := a & b << 1
		a = a ^ b
		fmt.Println(fmt.Sprintf("a & b << 1 : %b\n", c), fmt.Sprintf("a^b : %b\n", a))
		b = c
	}
	return a
}
