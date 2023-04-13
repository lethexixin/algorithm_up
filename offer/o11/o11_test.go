package o11

import (
	"fmt"
	"testing"
)

func TestF1(t *testing.T) {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray([]int{7, 0, 1, 1, 1, 1, 1, 2, 3, 4}))
}

func TestF2(t *testing.T) {
	fmt.Println(minArray2([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray2([]int{7, 0, 1, 1, 1, 1, 1, 2, 3, 4}))
}
