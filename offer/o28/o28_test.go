package o28

import (
	"fmt"
	"testing"
)

func TestF1(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 2}
	n4 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 4}
	n6 := &TreeNode{Val: 4}
	n7 := &TreeNode{Val: 3}

	n1.Left = n2
	n1.Right = n3

	n2.Left = n4
	n2.Right = n5

	n3.Left = n6
	n3.Right = n7

	fmt.Println(isSymmetric(n1))
}

func TestF2(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 2}
	n4 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 3}

	n1.Left = n2
	n1.Right = n3

	n2.Right = n4

	n3.Right = n5

	fmt.Println(isSymmetric(n1))
}
