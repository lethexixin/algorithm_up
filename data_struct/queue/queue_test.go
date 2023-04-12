package queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	// 初始化队列，使用队列的通用接口
	capacity := 10
	queue := NewListQueue(capacity)

	// 元素入队
	queue.push(1)
	queue.push(3)
	queue.push(2)
	queue.push(5)
	queue.push(4)
	fmt.Println("队列 queue = ", queue.toList())

	// 访问队首元素
	peekFront := queue.peekFront()
	fmt.Println("队首元素 peekFront =", peekFront)
	peekBack := queue.peekBack()
	fmt.Println("队尾元素 peekBack =", peekBack)

	// 元素出队
	pop := queue.pop()
	fmt.Print("出队元素 pop = ", pop, ", 出队后 queue = ", queue.toList())

	// 获取队的长度
	size := queue.size()
	fmt.Println("队的长度 size =", size)

	// 判断是否为空
	isEmpty := queue.empty()
	fmt.Println("队是否为空 =", isEmpty)

	/* 测试环形数组 */
	for i := 0; i < 10; i++ {
		queue.push(i)
		tmp := *queue
		queue.pop()
		fmt.Println("第", i, "轮入队, queue =", tmp.toList(), ", 出队后 queue =", queue.toList())
	}
}
