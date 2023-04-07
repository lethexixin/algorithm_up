package tree

import "container/list"

// TrNode 二叉树结点结构体
type TrNode struct {
	Val   int
	Left  *TrNode
	Right *TrNode
}

func NewTrNode(val int) *TrNode {
	return &TrNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// imitateTrNode 模拟树操作
func imitateTrNode() {
	/*
		             n1
				  /      \
				n2        n3
			  /	   \
			n4	   n5
	*/

	n1 := NewTrNode(1)
	n2 := NewTrNode(2)
	n3 := NewTrNode(3)
	n4 := NewTrNode(4)
	n5 := NewTrNode(5)

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5

	/* 插入与删除结点 */
	// 在 n1 -> n2 中间插入结点 P
	p := NewTrNode(0)
	n1.Left = p
	p.Left = n2

	// 删除结点 P
	n1.Left = n2
}

//todo 平衡二叉树的理解

// levelOrder 层序遍历
func (t *TrNode) levelOrder() []int {
	queue := list.New()
	queue.PushBack(t)

	vs := make([]int, 0)

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TrNode)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
		vs = append(vs, node.Val)
	}

	return vs
}
