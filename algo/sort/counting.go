package sort

// CountingSort1 计数排序
func CountingSort1(nums []int) {
	max := 0
	for _, num := range nums {
		if num >= max {
			max = num
		}
	}

	tmp := make([]int, max+1)
	for _, num := range nums {
		tmp[num]++
	}

	i := 0
	for j, v := range tmp {
		for k := 0; k < v; k++ {
			nums[i] = j
			i++
		}
	}
}

// CountingSort2 计数排序
func CountingSort2(nums []int) {
	m := 0
	for _, num := range nums {
		if num >= m {
			m = num
		}
	}

	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}

	for i := 0; i < len(counter)-1; i++ {
		counter[i+1] += counter[i]
	}

	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		res[counter[num]-1] = num
		counter[num]--
	}

	copy(nums, res)
}

type Number struct {
	Name string
	Num  int
}

// CountingSort3 计数排序
func CountingSort3(nums []Number) {
	// 1. 统计数组 nums 对象中的最大Num元素 m
	m := 0
	for _, num := range nums {
		if num.Num >= m {
			m = num.Num
		}
	}

	// 2. 统计各Num的出现次数
	// counter[num.Num] 代表 Num 的出现次数
	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num.Num]++
	}

	// 3. 求 counter 的前缀和，将“出现次数”转换为“尾索引”
	// 即 counter[num]-1 是 num 在 res 中最后一次出现的索引
	for i := 0; i < m; i++ {
		counter[i+1] += counter[i]
	}

	// 4. 倒序遍历 nums，将各元素填入结果数组 res
	// 初始化数组 res 用于记录结果
	n := len(nums)
	res := make([]Number, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		// 将 num 放置到对应索引处
		res[counter[num.Num]-1] = num
		// 令前缀和自减 1 ，得到下次放置 num 的索引
		counter[num.Num]--
	}

	// 使用结果数组 res 覆盖原数组 nums
	copy(nums, res)
}
