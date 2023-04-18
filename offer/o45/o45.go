package o45

import (
	"sort"
	"strconv"
	"strings"
)

//剑指 Offer 45. 把数组排成最小的数
//输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
//
//
//示例 1:
//
//输入: [10,2]
//输出: "102"
//示例2:
//
//输入: [3,30,34,5,9]
//输出: "3033459"
//
//提示:
//
//0 < nums.length <= 100
//说明:
//
//输出结果可能非常大，所以你需要返回一个字符串而不是整数
//拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/59ypcj/

func minNumber(nums []int) string {
	a := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		a = append(a, strconv.Itoa(nums[i]))
	}
	sort.Slice(a, func(i, j int) bool {
		a1 := a[i] + a[j]
		b1 := a[j] + a[i]
		//fmt.Println(i, j, a1, b1)
		a2, _ := strconv.Atoi(a1)
		b2, _ := strconv.Atoi(b1)
		return a2 < b2
	})
	return strings.Join(a, "")
}
