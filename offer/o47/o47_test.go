package o47

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	grid := make([][]int, 0)
	grid = append(grid, []int{1, 3, 1})
	grid = append(grid, []int{1, 5, 1})
	grid = append(grid, []int{4, 2, 1})
	fmt.Println(maxValue(grid))
}
