package queue

import (
	"container/list"
	"fmt"
)

// imitateQueue 模拟队列相关操作
func imitateQueue() {
	//初始化队列, 在 Go 中, 将 list 作为队列来使用
	queue := list.New()

	//入队
	queue.PushBack(1)
	queue.PushBack(3)
	queue.PushBack(2)
	queue.PushBack(5)
	queue.PushBack(4)

	//访问队首元素
	fmt.Println(queue.Front().Value)

	//出队
	queue.Remove(queue.Front())

	//获取队列长度
	fmt.Println(queue.Len())

	//判空
	fmt.Println(queue.Len() == 0)
}

// LinkedQueue 基于链表实现的队列
type LinkedQueue struct {
	// 使用内置包 list 来实现队列
	data *list.List
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{data: list.New()}
}

func (q *LinkedQueue) size() int {
	return q.data.Len()
}

func (q *LinkedQueue) empty() bool {
	return q.data.Len() == 0
}

// 获取队首元素
func (q *LinkedQueue) peekFront() any {
	if q.empty() {
		return nil
	}

	return q.data.Front().Value
}

// 获取队尾元素
func (q *LinkedQueue) peekBack() any {
	if q.empty() {
		return nil
	}

	return q.data.Back().Value
}

// push 入队
func (q *LinkedQueue) push(val any) {
	q.data.PushBack(val)
}

// pop 出队
func (q *LinkedQueue) pop() any {
	if q.empty() {
		return nil
	}

	val := q.data.Front()

	q.data.Remove(val)

	return val.Value
}

// toList 获取 List 用于打印
func (q *LinkedQueue) toList() *list.List {
	return q.data
}

// ListQueue todo 基于环形数组实现的队列
type ListQueue struct {
	data []any
}
