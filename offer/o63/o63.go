package o63

import "math"

//剑指 Offer 63. 股票的最大利润
//假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
//
//
//
//示例 1:
//
//输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
//示例 2:
//
//输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
//
//
//限制：
//
//0 <= 数组长度 <= 10^5
//
//
//
//注意：本题与主站 121 题相同：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/58nn7r/

// [7,1,5,3,6,4]
// 执行用时： 4 ms, 在所有 Go 提交中击败了 82.13% 的用户
// 内存消耗： 2.9 MB, 在所有 Go 提交中击败了 52.16% 的用户
func maxProfit(prices []int) int {
	cost, profit := math.MaxInt64, 0

	for _, price := range prices {
		cost = int(math.Min(float64(cost), float64(price)))
		profit = int(math.Max(float64(profit), float64(price-cost)))
	}
	return profit
}
