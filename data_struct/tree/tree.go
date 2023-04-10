package tree

import "container/list"

// TrNode 二叉树结点结构体
type TrNode struct {
	Val    int
	Height int //节点高度
	Left   *TrNode
	Right  *TrNode
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

// preOrder 前序遍历
func (t *TrNode) preOrder(nums []int) {
	if t == nil {
		return
	}

	// 根节点, 左节点, 右节点
	nums = append(nums, t.Val)
	t.Left.preOrder(nums)
	t.Right.preOrder(nums)
}

// inOrder 中序遍历
func (t *TrNode) inOrder(nums []int) {
	if t == nil {
		return
	}

	// 左子树, 根子树, 右子树
	t.Left.inOrder(nums)
	nums = append(nums, t.Val)
	t.Right.inOrder(nums)
}

// postOrder 后序遍历
func (t *TrNode) postOrder(nums []int) {
	if t == nil {
		return
	}

	// 左子树, 右子树, 根子树
	t.Left.postOrder(nums)
	t.Right.postOrder(nums)
	nums = append(nums, t.Val)
}

func generateBinary() (t *TrNode) {
	root := NewTrNode(8)
	n2 := NewTrNode(4)
	n3 := NewTrNode(12)
	n4 := NewTrNode(2)
	n5 := NewTrNode(6)
	n6 := NewTrNode(10)
	n7 := NewTrNode(14)
	n8 := NewTrNode(1)
	n9 := NewTrNode(3)
	n10 := NewTrNode(5)
	n11 := NewTrNode(7)
	n12 := NewTrNode(9)
	n13 := NewTrNode(11)
	n14 := NewTrNode(13)
	n15 := NewTrNode(15)

	root.Left = n2
	root.Right = n3

	n2.Left = n4
	n2.Right = n5

	n3.Left = n6
	n3.Right = n7

	n4.Left = n8
	n4.Right = n9

	n5.Left = n10
	n5.Right = n11

	n6.Left = n12
	n6.Right = n13

	n7.Left = n14
	n7.Right = n15

	return root
}

func (t *TrNode) binarySearch(target int) *TrNode {
	r := t
	for r != nil {
		if r.Val < target {
			r = r.Right
		} else if r.Val > target {
			r = r.Left
		} else {
			return r
		}
	}
	return nil
}

func (t *TrNode) binaryInsert(target int) {
	r := t
	var pre *TrNode
	for r != nil {
		if r.Val == target {
			return
		}
		pre = r
		if r.Val > target {
			r = r.Left
		} else {
			r = r.Right
		}
	}

	newTr := NewTrNode(target)
	if pre.Val > target {
		pre.Left = newTr
	} else {
		pre.Right = newTr
	}
}

// so hard, todo
func (t *TrNode) binaryRemove(target int) {

}

// so hard, avl tree, todo
