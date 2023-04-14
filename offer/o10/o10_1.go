package o10

//剑指 Offer 10- I. 斐波那契数列
//写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
//
//F(0) = 0, F(1)= 1
//F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
//斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
//
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//
//
//示例 1：
//
//输入：n = 2
//输出：1
//示例 2：
//
//输入：n = 5
//输出：5
//
//提示：
//
//0 <= n <= 100
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/50fxu1/

// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 1.8 MB, 在所有 Go 提交中击败了 58.46% 的用户
func fib(n int) int {
	a, b, sum := 0, 1, 0
	for i := 0; i < n; i++ {
		//sum = a + b
		//考虑到 n 很大时, 可能会超过 int/int64 的最大值, 所以 sum = (a + b) % 1000000007
		sum = (a + b) % 1000000007
		a = b
		b = sum
	}

	return a
}

// 0, 1, 1, 2, 3, 5, 8, 13
// 当 n 比较小时, 可以用这个, 但是递归很多重复计算的逻辑, 使用动态规划会更好
func fib2(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
