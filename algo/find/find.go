package find

import "algorithm_up/data_struct/linked"

// linearSearchArray 数组线性查找
func linearSearchArray(nums []int, target int) (idx int) {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

// linearSearchLinked 链表线性查找
func linearSearchLinked(node *linked.LkNode, target int) *linked.LkNode {
	if node.Val == target {
		return node
	}

	for node.Next != nil {
		if node.Val == target {
			return node
		}
		node = node.Next
	}
	return nil
}

// binarySearch 双闭区间二分查找
func binarySearch(nums []int, target int) (idx int) {
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		if nums[m] > target {
			j = m - 1
		} else if nums[m] < target {
			i = m + 1
		} else {
			return m
		}
	}
	return -1
}

// hashSearchArray HASH查找, m: key:值, value:索引
func hashSearchArray(m map[int]int, target int) int {
	if idx, ok := m[target]; ok {
		return idx
	}
	return -1
}

// hashSearchLinked HASH查找, m: key:节点值, value:节点
func hashSearchLinked(m map[int]*linked.LkNode, target int) *linked.LkNode {
	if lk, ok := m[target]; ok {
		return lk
	}
	return nil
}
