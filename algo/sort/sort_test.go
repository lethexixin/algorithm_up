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

func TestMergeSort(t *testing.T) {
	nums := []int{7, 3, 2, 6, 0, 1, 8, 5, 4}
	MergeSort(nums, 0, len(nums)-1, "TestMergeSort")
	fmt.Println(nums)
}

func TestBucketSort(t *testing.T) {
	nums := []float64{0.49, 0.96, 0.82, 0.09, 0.57, 0.43, 0.91, 0.75, 0.15, 0.37}
	BucketSort(nums)
	fmt.Println(nums)
}

func TestCountingSort1(t *testing.T) {
	nums := []int{1, 0, 1, 2, 0, 4, 0, 2, 2, 4}
	CountingSort1(nums)
	fmt.Println(nums)
}

func TestCountingSort2(t *testing.T) {
	nums := []int{1, 0, 1, 2, 0, 4, 0, 2, 2, 4}
	CountingSort2(nums)
	fmt.Println(nums)
}

func TestCountingSort3(t *testing.T) {
	num := []Number{
		{"name1", 1}, {"name2", 0}, {"name3", 1},
		{"name4", 2}, {"name5", 0}, {"name6", 4},
		{"name7", 0}, {"name8", 2}, {"name9", 2},
		{"name10", 4}}
	CountingSort3(num)
	fmt.Println(num)
	//[{name2 0} {name5 0} {name7 0} {name1 1} {name3 1} {name4 2} {name8 2} {name9 2} {name6 4} {name10 4}]
}

func TestRadixSort(t *testing.T) {
	nums := []int{10546151, 35663510, 42865989, 34862445, 81883077,
		88906420, 72429244, 30524779, 82060337, 63832996}
	RadixSort(nums)
	fmt.Println(nums)
}
