package sort

// RadixSort 基数排序
func RadixSort(nums []int) {
	max := 0
	for _, num := range nums {
		if num >= max {
			max = num
		}
	}

	for exp := 1; exp <= max; exp = exp * 10 {
		countingSortDigit(nums, exp)
	}
}

func digit(num int, exp int) int {
	return (num / exp) % 10
}

func countingSortDigit(nums []int, exp int) {
	tmp := make([]int, 10)
	n := len(nums)
	for _, num := range nums {
		idx := digit(num, exp)
		tmp[idx]++
	}
	for i := 1; i < 10; i++ {
		tmp[i] += tmp[i-1]
	}
	// 倒序遍历，根据桶内统计结果, 将各元素填入 res
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		d := digit(nums[i], exp)
		j := tmp[d] - 1  // 获取 d 在数组中的索引 j
		res[j] = nums[i] // 将当前元素填入索引 j
		tmp[d]--         // 将 d 的数量减 1
	}
	// 使用结果覆盖原数组 nums
	copy(nums, res)
}
