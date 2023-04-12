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
	data    []any //用于存储队列元素的数组
	front   int   //队首指针, 指向队首元素
	queSize int   //队列长度
	queCap  int   //队列容量(最大容纳元素数量)
}

func NewListQueue(queCap int) *ListQueue {
	return &ListQueue{
		data:    make([]any, queCap),
		front:   0,
		queSize: 0,
		queCap:  queCap,
	}
}

func (q *ListQueue) size() int {
	return q.queSize
}

func (q *ListQueue) empty() bool {
	return q.queSize == 0
}

// push 入队
func (q *ListQueue) push(val any) {
	if q.queSize == q.queCap {
		return
	}
	// 计算尾指针，指向队尾索引 + 1
	// 通过取余操作，实现 rear 越过数组尾部后回到头部
	rear := (q.front + q.queSize) % q.queCap
	// 将 num 添加至队尾
	q.data[rear] = val
	q.queSize++
}

// pop 出队
func (q *ListQueue) pop() any {
	val := q.peekFront()
	// 队首指针向后移动一位，若越过尾部则返回到数组头部
	q.front = (q.front + 1) % q.queCap
	q.queSize--
	return val
}

// peekFront 访问队首元素
func (q *ListQueue) peekFront() any {
	if q.empty() {
		return nil
	}

	return q.data[q.front]
}

// peekBack 访问队尾元素
func (q *ListQueue) peekBack() any {
	if q.empty() {
		return nil
	}

	if q.front-1 >= 0 && q.front-1 < q.queSize {
		return q.data[q.front-1]
	}
	return nil
}

func (q *ListQueue) toList() []any {
	rear := q.front + q.queSize
	if rear >= q.queCap {
		rear %= q.queCap
		return append(q.data[q.front:], q.data[:rear]...)
	}
	return q.data[q.front:rear]
}
