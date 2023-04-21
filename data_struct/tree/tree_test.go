package tree

import (
	"fmt"
	"testing"
)

func TestBinaryInsert(t *testing.T) {
	generateBinary().binaryInsert(16)
}

func TestInOrder(t *testing.T) {
	n := make([]int, 0)
	generateBinary().inOrder(&n)
	fmt.Println(n)
}
