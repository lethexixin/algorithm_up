package o53

import (
	"fmt"
	"testing"
)

func TestF1(t *testing.T) {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}

func TestF2(t *testing.T) {
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
}
