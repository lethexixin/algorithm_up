package o45

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(minNumber([]int{3, 30, 34, 1, 2}))
}
