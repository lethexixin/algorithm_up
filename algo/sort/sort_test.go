package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{2, 4, 1, 0, 3, 5, 7, 6}

	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func TestMidQuickSort(t *testing.T) {
	nums := []int{2, 4, 1, 0, 3, 5, 7, 6}

	MidQuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func TestTailCallQuickSort(t *testing.T) {
	nums := []int{2, 4, 1, 0, 3, 5, 7, 6}

	TailCallQuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
