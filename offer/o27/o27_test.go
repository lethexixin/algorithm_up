package o27

import (
	"testing"
)

func TestF(t *testing.T) {
	n1 := &TreeNode{Val: 4}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 7}
	n4 := &TreeNode{Val: 1}
	n5 := &TreeNode{Val: 3}
	n6 := &TreeNode{Val: 6}
	n7 := &TreeNode{Val: 8}

	n1.Left = n2
	n1.Right = n3

	n2.Left = n4
	n2.Right = n5

	n3.Left = n6
	n3.Right = n7

	mirrorTree(n1)
}
