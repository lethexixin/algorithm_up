package linked

// LkNode 链表结点结构体
type LkNode struct {
	Val  int     //节点值
	Next *LkNode //指向下一个节点的指针
}

// NewLkNode 构造函数,创建一个新的链表
func NewLkNode(val int) *LkNode {
	return &LkNode{
		Val:  val,
		Next: nil,
	}
}

func initLN() {
	node1 := NewLkNode(1)
	node2 := NewLkNode(3)
	node3 := NewLkNode(2)
	node4 := NewLkNode(5)
	node5 := NewLkNode(4)

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
}

// insertLN 在链表的结点 n0 之后插入结点 p
func insertLN(n0 *LkNode, p *LkNode) {
	next := n0.Next
	p.Next = next
	n0.Next = p
}

// removeLN 删除链表的结点 n0 之后的首个结点
func removeLN(n0 *LkNode) {
	if n0.Next == nil {
		return
	}
	n0.Next = n0.Next.Next
}

// rangeLN 访问链表中索引为 index 的结点 */
func rangeLN(head *LkNode, idx int) *LkNode {
	for i := 0; i < idx; i++ {
		if head.Next == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

// findLN 在链表中查找值为 target 的首个结点
func findLN(head *LkNode, target int) *LkNode {
	if head.Val == target {
		return head
	}

	for head.Next != nil {
		if head.Val == target {
			return head
		}
		head = head.Next
	}
	return nil
}

// DoubleLkNode 双向链表
type DoubleLkNode struct {
	Val  int
	Prev *DoubleLkNode //指向前驱节点的指针
	Next *DoubleLkNode //指向后继节点的指针
}

// NewDoubleLkNode 初始化
func NewDoubleLkNode(val int) *DoubleLkNode {
	return &DoubleLkNode{
		Val:  val,
		Prev: nil,
		Next: nil,
	}
}
