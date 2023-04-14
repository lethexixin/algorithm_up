package o63

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))

	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
