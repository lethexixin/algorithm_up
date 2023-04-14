package o10

//剑指 Offer 10- II. 青蛙跳台阶问题
//一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n级的台阶总共有多少种跳法。
//
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//
//示例 1：
//
//输入：n = 2
//输出：2
//示例 2：
//
//输入：n = 7
//输出：21
//示例 3：
//
//输入：n = 0
//输出：1
//提示：
//
//0 <= n <= 100
//注意：本题与主站 70 题相同：https://leetcode-cn.com/problems/climbing-stairs/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/57hyl5/

// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 1.9 MB, 在所有 Go 提交中击败了 13.32% 的用户
func numWays(n int) int {
	a, b, sum := 1, 1, 0
	for i := 0; i < n; i++ {
		sum = (a + b) % 1000000007
		a = b
		b = sum
	}

	return a
}
