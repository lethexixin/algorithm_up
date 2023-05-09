package o31

//剑指 Offer 31. 栈的压入、弹出序列
//输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
//假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，
//序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
//
//
//
//示例 1：
//
//输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
//输出：true
//解释：我们可以按以下顺序执行：
//push(1), push(2), push(3), push(4), pop() -> 4,
//push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//示例 2：
//
//输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
//输出：false
//解释：1 不能在 2 之前弹出。
//
//
//提示：
//
//0 <= pushed.length == popped.length <= 1000
//0 <= pushed[i], popped[i] < 1000
//pushed 是 popped 的排列。
//注意：本题与主站 946 题相同：https://leetcode-cn.com/problems/validate-stack-sequences/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5wh1hj/

// 执行用时：4 ms, 在所有 Go 提交中击败了 93.21% 的用户
// 内存消耗：3.7 MB, 在所有 Go 提交中击败了 40.92% 的用户
func validateStackSequences(pushed []int, popped []int) bool {
	tmp := make([]int, 0)
	i := 0

	for _, v := range pushed {
		tmp = append(tmp, v)
		for len(tmp) > 0 && tmp[len(tmp)-1] == popped[i] {
			tmp = tmp[:len(tmp)-1]
			i++
		}
	}
	return len(tmp) == 0
}
