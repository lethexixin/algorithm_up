package o16

import "math"

//剑指 Offer 16. 数值的整数次方
//实现pow(x,n)，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。
//
//
//示例 1：
//
//输入：x = 2.00000, n = 10
//输出：1024.00000
//示例 2：
//
//输入：x = 2.10000, n = 3
//输出：9.26100
//示例 3：
//
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2^-2 = 1/(2^2) = 1/4 = 0.25
//
//提示：
//
//-100.0 <x< 100.0
//-2^31<= n <=2^31-1
//-10^4<= x^n<= 10^4
//
//注意：本题与主站 50 题相同：https://leetcode-cn.com/problems/powx-n/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/57rwmg/

// 普通乘法,n比较大时,容易超时
func myPow1(x float64, n int) float64 {
	mul := 1.0
	if n > 0 {
		for i := 1; i <= n; i++ {
			mul *= x
		}
		return mul
	} else if n < 0 {
		for i := 1; i <= -n; i++ {
			mul *= 1 / x
		}
		return mul
	}

	return mul
}

// 执行用时： 0 ms , 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗： 1.9 MB , 在所有 Go 提交中击败了 42.72% 的用户
func myPow(x float64, n int) float64 {
	res := 1.0
	if n < 0 {
		x = 1 / x
		n = -n
	}

	for n > 0 {
		if n%2 == 1 {
			res *= x
		}
		x *= x
		n = int(math.Floor(float64(n / 2)))
	}
	return res
}
