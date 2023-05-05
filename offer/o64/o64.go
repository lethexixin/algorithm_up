package o64

//剑指 Offer 64. 求 1 + 2 + … + n
//求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
//
//
//
//示例 1：
//
//输入: n = 3
//输出: 6
//示例 2：
//
//输入: n = 9
//输出: 45
//
//
//限制：
//
//1 <= n <= 10000
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/9h44cj/

// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 2.6 MB, 在所有 Go 提交中击败了 6.25% 的用户
func sumNums(n int) int {
	res := 0
	recur(n, &res)
	return res
}

func recur(n int, res *int) int {
	_ = n > 1 && recur(n-1, res) > 0
	*res += n
	return *res
}
