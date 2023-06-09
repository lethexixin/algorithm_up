package o66

//剑指 Offer 66. 构建乘积数组
//给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，
//其中B[i] 的值是数组 A 中除了下标 i 以外的元素的积,
//即B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
//
//
//示例:
//
//输入: [1,2,3,4,5]
//输出: [120,60,40,30,24]
//
//提示：
//
//所有元素乘积之和不会溢出 32 位整数
//a.length <= 100000
//相关标签
//
//Go
//
//
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/57d8cm/

func constructArr(a []int) []int {
	if len(a) == 0 {
		return []int{0}
	}

	b := make([]int, len(a))
	b[0] = 1
	tmp := 1

	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}

	for i := len(a) - 2; i >= 0; i-- {
		tmp *= a[i+1]
		b[i] *= tmp
	}

	return b
}
