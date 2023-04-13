package o3

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))

	fmt.Println(findRepeatNumber2([]int{2, 3, 1, 0, 2, 5, 3}))
}
