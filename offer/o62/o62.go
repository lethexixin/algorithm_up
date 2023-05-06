package o62

import "fmt"

//剑指 Offer 62. 圆圈中最后剩下的数字
//0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，
//每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。
//
//例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，
//则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
//
//
//示例 1：
//
//输入: n = 5, m = 3
//输出:3
//示例 2：
//
//输入: n = 10, m = 17
//输出:2
//
//限制：
//
//1 <= n <= 10^5
//1 <= m <= 10^6
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/oxrkot/

// 执行用时： 4 ms, 在所有 Go 提交中击败了 99.75% 的用户
// 内存消耗： 1.9 MB, 在所有 Go 提交中击败了 51.73% 的用户
func lastRemaining(n int, m int) int {
	lastFn := 0 //代表 f(n=1) 问题的解, 恒为 0
	for i := 2; i <= n; i++ {
		lastFn = (lastFn + m) % i
		fmt.Println(lastFn)
	}
	return lastFn
}
