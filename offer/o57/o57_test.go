package o57

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	fmt.Println(twoSum([]int{10, 26, 30, 31, 47, 60}, 40))

	fmt.Println(findContinuousSequence(9))
	fmt.Println(findContinuousSequence(15))
}
