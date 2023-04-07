package stack

import (
	"container/list"
	"fmt"
)

// imitateStack 模拟栈相关操作
func imitateStack() {
	var stack []int

	//入栈
	stack = append(stack, 1)
	stack = append(stack, 3)
	stack = append(stack, 2)
	stack = append(stack, 5)
	stack = append(stack, 4)

	//访问栈顶元素
	peek := stack[len(stack)-1]
	fmt.Println(peek)

	//出栈
	pop := stack[len(stack)-1]
	fmt.Println(pop)
	stack = stack[:len(stack)-1]

	//获取栈的长度
	size := len(stack)
	fmt.Println(size)
}

// LinkedStack 基于链表实现的栈
type LinkedStack struct {
	//使用内置包 list 来实现栈
	data *list.List
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{data: list.New()}
}

// toList 获取 List 用于打印
func (s *LinkedStack) toList() *list.List {
	return s.data
}

// 获取栈的长度
func (s *LinkedStack) size() int {
	return s.data.Len()
}

func (s *LinkedStack) empty() bool {
	return s.data.Len() == 0
}

// push 入栈
func (s *LinkedStack) push(val int) {
	s.data.PushBack(val)
}

// pop 出栈
func (s *LinkedStack) pop() any {
	if s.empty() {
		return nil
	}

	v := s.data.Back()
	s.data.Remove(v)
	return v.Value
}

// peek 访问栈顶元素
func (s *LinkedStack) peek() any {
	if s.empty() {
		return nil
	}

	return s.data.Back().Value
}

// ListStack 基于数组实现的栈
type ListStack struct {
	data []any
}

func NewListStack() *ListStack {
	return &ListStack{data: make([]any, 0, 16)}
}

// toList 获取 List 用于打印
func (s *ListStack) toList() []any {
	return s.data
}

// 获取栈的长度
func (s *ListStack) size() int {
	return len(s.data)
}

func (s *ListStack) empty() bool {
	return len(s.data) == 0
}

// push 入栈
func (s *ListStack) push(val int) {
	s.data = append(s.data, val)
}

// pop 出栈
func (s *ListStack) pop() any {
	if s.empty() {
		return nil
	}

	v := s.peek()
	s.data = s.data[:s.size()-1]
	return v
}

// peek 访问栈顶元素
func (s *ListStack) peek() any {
	if s.empty() {
		return nil
	}

	return s.data[s.size()-1]
}
