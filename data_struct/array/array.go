package array

import (
	"fmt"
	"math/rand"
)

/* 随机返回数组中的一个值 */
func randomAccess(nums []int) (randomNum int) {
	idx := rand.Intn(len(nums))
	return nums[idx]
}

/* 在数组的索引 index 处插入元素 num */
func insert(nums []int, val int, idx int) {
	for i := len(nums) - 1; i > idx; i-- {
		nums[i] = nums[i-1]
	}
	nums[idx] = val
}

/* 删除索引 index 处元素 */
func remove(nums []int, idx int) (val int) {
	val = nums[idx]
	for i := idx; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
	return val
}

/* 遍历数组 */
func rangeArray(nums []int) {
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for i := range nums {
		fmt.Println(i)
	}
}

/* 在数组中查找指定元素 */
func find(nums []int, target int) (idx int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
