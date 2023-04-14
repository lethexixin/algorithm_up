package o32

import (
	"algorithm_up/offer/o32/o32_1"
	"algorithm_up/offer/o32/o32_2"
	"algorithm_up/offer/o32/o32_3"
	"fmt"
	"testing"
)

func TestF1(t *testing.T) {
	n1 := &o32_1.TreeNode{Val: 3}
	n2 := &o32_1.TreeNode{Val: 9}
	n3 := &o32_1.TreeNode{Val: 20}
	n4 := &o32_1.TreeNode{Val: 15}
	n5 := &o32_1.TreeNode{Val: 7}

	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5

	fmt.Println(o32_1.LevelOrder(n1))

	n6 := &o32_1.TreeNode{Val: 5}
	n7 := &o32_1.TreeNode{Val: 8}
	n2.Left = n6
	n2.Right = n7
	fmt.Println(o32_1.LevelOrder(n1))
}

func TestF2(t *testing.T) {
	n1 := &o32_2.TreeNode{Val: 3}
	n2 := &o32_2.TreeNode{Val: 9}
	n3 := &o32_2.TreeNode{Val: 20}
	n4 := &o32_2.TreeNode{Val: 15}
	n5 := &o32_2.TreeNode{Val: 7}

	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5

	fmt.Println(o32_2.LevelOrder(n1))

	n6 := &o32_2.TreeNode{Val: 5}
	n7 := &o32_2.TreeNode{Val: 8}
	n2.Left = n6
	n2.Right = n7
	fmt.Println(o32_2.LevelOrder(n1))
}

func TestF3(t *testing.T) {
	n1 := &o32_3.TreeNode{Val: 3}
	n2 := &o32_3.TreeNode{Val: 9}
	n3 := &o32_3.TreeNode{Val: 20}
	n4 := &o32_3.TreeNode{Val: 15}
	n5 := &o32_3.TreeNode{Val: 7}

	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5

	fmt.Println(o32_3.LevelOrder(n1))

	n6 := &o32_3.TreeNode{Val: 5}
	n7 := &o32_3.TreeNode{Val: 8}
	n2.Left = n6
	n2.Right = n7
	fmt.Println(o32_3.LevelOrder(n1))
}
