package sort

// BubbleSort1 冒泡排序普通版
func BubbleSort1(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// BubbleSort2 冒泡排序flag优化版
func BubbleSort2(nums []int) {
	for i := len(nums) - 1; i > 0; i++ {
		flag := false
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			// 此轮冒泡未交换任何元素, 直接跳出
			break
		}
	}
}
