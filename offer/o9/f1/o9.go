package f1

import "container/list"

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

// 方法1： 通过 list.List 实现
//执行用时： 164 ms , 在所有 Go 提交中击败了 80.99% 的用户
//内存消耗： 7.9 MB , 在所有 Go 提交中击败了 66.57% 的用户

type CQueue struct {
	a *list.List
	b *list.List
}

func Constructor() CQueue {
	return CQueue{a: list.New(), b: list.New()}
}

func (q *CQueue) AppendTail(value int) {
	q.a.PushBack(value)
}

func (q *CQueue) DeleteHead() int {
	if q.b.Len() != 0 {
		v := q.b.Back()
		q.b.Remove(v)
		return v.Value.(int)
	}

	if q.a.Len() == 0 {
		return -1
	}

	for q.a.Len() > 0 {
		q.b.PushBack(q.a.Remove(q.a.Back()))
	}
	return q.b.Remove(q.b.Back()).(int)
}
