package o30

//剑指 Offer 30. 包含 min 函数的栈
//定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
//调用 min、push 及 pop 的时间复杂度都是 O(1)。
//
//
//示例:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.min();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.min();   --> 返回 -2.
//
//提示：
//
//各函数的调用总次数不超过 20000 次
//
//注意：本题与主站 155 题相同：https://leetcode-cn.com/problems/min-stack/
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/50bp33/

type MinStack struct {
	data []int
	sup  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	s.data = append(s.data, x)
	if len(s.sup) == 0 {
		s.sup = append(s.sup, x)
	} else {
		if x <= s.sup[len(s.sup)-1] {
			s.sup = append(s.sup, x)
		}
	}
}

func (s *MinStack) Pop() {
	top := s.Top()
	s.data = s.data[:len(s.data)-1]
	if s.sup[len(s.sup)-1] == top {
		s.sup = s.sup[:len(s.sup)-1]
	}
}

func (s *MinStack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *MinStack) Min() int {
	return s.sup[len(s.sup)-1]
}

//执行用时： 16 ms , 在所有 Go 提交中击败了 45.90% 的用户
//内存消耗： 7.1 MB , 在所有 Go 提交中击败了 87.11% 的用户
