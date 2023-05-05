package o7

import (
	"testing"
)

func TestF(t *testing.T) {
	//buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	buildTree([]int{3, 9, 2, 1, 7}, []int{9, 3, 1, 2, 7})
	/*buildTree([]int{3, 9, 1, 7, 2}, []int{1, 9, 7, 3, 2})
	buildTree([]int{3, 9, 1, 2}, []int{1, 9, 3, 2})*/
}
