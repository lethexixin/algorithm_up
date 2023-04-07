package deque

import (
	"container/list"
	"fmt"
)

// imitateDeque 模拟双向队列操作
func imitateDeque() {
	/* 初始化双向队列 */
	// 在 Go 中，将 list 作为双向队列使用
	deque := list.New()

	//入队
	deque.PushBack(2)
	deque.PushBack(5)
	deque.PushBack(4)
	deque.PushFront(3)
	deque.PushFront(1)

	peekFront := deque.Front()
	fmt.Println(peekFront.Value)
	peekBack := deque.Back()
	fmt.Println(peekBack.Value)

	deque.Remove(peekFront)
	deque.Remove(peekBack)

	fmt.Println(deque.Len())
	fmt.Println(deque.Len() == 0)
}

// LinkedDeque 基于双向链表实现的双向队列
type LinkedDeque struct {
	data *list.List
}

func NewLinkedDeque() *LinkedDeque {
	return &LinkedDeque{data: list.New()}
}

func (s *LinkedDeque) size() int {
	return s.data.Len()
}

func (s *LinkedDeque) empty() bool {
	return s.data.Len() == 0
}

// toList 获取 List 用于打印
func (s *LinkedDeque) toList() *list.List {
	return s.data
}

// pushFirst 队首元素入队
func (s *LinkedDeque) pushFirst(value any) {
	s.data.PushFront(value)
}

// pushLast 队尾元素入队
func (s *LinkedDeque) pushLast(value any) {
	s.data.PushBack(value)
}

// popFirst 队首元素出队
func (s *LinkedDeque) popFirst() any {
	if s.empty() {
		return nil
	}

	e := s.data.Front()
	s.data.Remove(e)
	return e.Value
}

// popLast 队尾元素出队
func (s *LinkedDeque) popLast() any {
	if s.empty() {
		return nil
	}

	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}

// peekFirst 访问队首元素
func (s *LinkedDeque) peekFirst() any {
	if s.empty() {
		return nil
	}

	return s.data.Front().Value
}

// peekLast 访问队尾元素
func (s *LinkedDeque) peekLast() any {
	if s.empty() {
		return nil
	}

	return s.data.Back().Value
}

// ListDeque todo 基于环形数组实现的双向队列
type ListDeque struct {
	nums []any
}
