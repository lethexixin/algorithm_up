package sort

import "sort"

// QuickSort 快速排序
func QuickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}

	base := partition(nums, left, right)
	QuickSort(nums, left, base-1)
	QuickSort(nums, base+1, right)
}

func partition(nums []int, left int, right int) (idx int) {
	l, r := left, right

	for l < r {
		for l < r && nums[r] >= nums[left] {
			r--
		}
		for l < r && nums[l] <= nums[left] {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[l], nums[left] = nums[left], nums[l]
	return l
}

// quickSortGetMidIdx 获取数组left,mid,right索引处的值的中位数对应的索引
func quickSortGetMidIdx(nums []int, left, mid, right int) (idx int) {
	t := []int{nums[left], nums[mid], nums[right]}
	sort.Ints(t)
	if nums[left] == t[1] {
		return left
	}
	if nums[right] == t[1] {
		return right
	}
	return mid
}

// MidQuickSort 中位数快排
func MidQuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	base := midPartition(nums, left, right)
	MidQuickSort(nums, left, base-1)
	MidQuickSort(nums, base+1, right)
}

// midPartition 中位数快排
func midPartition(nums []int, left, right int) int {
	mid := quickSortGetMidIdx(nums, left, (left+right)/2, right)
	nums[mid], nums[left] = nums[left], nums[mid]

	l, r := left, right
	for l < r {
		for l < r && nums[r] >= nums[left] {
			r--
		}
		for l < r && nums[l] <= nums[left] {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}

	nums[l], nums[left] = nums[left], nums[l]
	return l
}

// TailCallQuickSort 尾递归快排
func TailCallQuickSort(nums []int, left, right int) {
	for left < right {
		base := partition(nums, left, right)
		if base-left < right-base {
			TailCallQuickSort(nums, left, base-1)
			left = base + 1
		} else {
			TailCallQuickSort(nums, base+1, right)
			right = base - 1
		}
	}
}
