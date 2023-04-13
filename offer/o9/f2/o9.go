package f2

//剑指 Offer 09. 用两个栈实现一个队列。
//队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
//分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//
//
//示例 1：
//
//输入：
//["CQueue","appendTail","deleteHead","deleteHead","deleteHead"]
//[[],[3],[],[],[]]
//输出：[null,null,3,-1,-1]
//示例 2：
//
//输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
//提示：
//
//1 <= values <= 10000
//最多会对 appendTail、deleteHead 进行 10000 次调用
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/5d3i87/

// 方法2： 通过 []int 数组 实现
//执行用时： 168 ms , 在所有 Go 提交中击败了 63.83% 的用户
//内存消耗： 7.9 MB , 在所有 Go 提交中击败了 57.71% 的用户

type CQueue struct {
	inStack, outStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (cq *CQueue) AppendTail(value int) {
	cq.inStack = append(cq.inStack, value)
}

func (cq *CQueue) In2out() {
	for len(cq.inStack) != 0 {
		cq.outStack = append(cq.outStack, cq.inStack[len(cq.inStack)-1])
		cq.inStack = cq.inStack[:len(cq.inStack)-1]
	}
}

func (cq *CQueue) DeleteHead() int {
	if len(cq.outStack) == 0 {
		if len(cq.inStack) == 0 {
			return -1
		}
		cq.In2out()
	}
	x := cq.outStack[len(cq.outStack)-1]
	cq.outStack = cq.outStack[:len(cq.outStack)-1]
	return x
}
