package o4

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	x := make([][]int, 0)
	x = append(x, []int{1, 4, 7, 11, 15})
	x = append(x, []int{2, 5, 8, 12, 19})
	x = append(x, []int{3, 6, 9, 16, 22})
	x = append(x, []int{10, 13, 14, 17, 24})
	x = append(x, []int{18, 21, 23, 26, 30})

	fmt.Println(findNumberIn2DArray(x, 8))
}
