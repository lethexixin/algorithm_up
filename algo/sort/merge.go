package sort

import "fmt"

func MergeSort(nums []int, left, right int, s string) {
	if left >= right {
		return
	}
	//划分阶段
	mid := (left + right) / 2
	MergeSort(nums, left, mid, "1")
	MergeSort(nums, mid+1, right, "2")
	//合并阶段
	merge(nums, left, mid, right, s)
}

/* 合并左子数组和右子数组 */
// 左子数组区间 [left, mid]
// 右子数组区间 [mid + 1, right]
func merge(nums []int, left, mid, right int, s string) {
	fmt.Println("type:", s, ", nums:", nums)
	fmt.Println("left:", left, ", right:", right, ", mid:", mid)

	// 初始化辅助数组 借助 copy 模块
	tmp := make([]int, right-left+1)
	for i := left; i <= right; i++ {
		tmp[i-left] = nums[i]
	}
	fmt.Println("tmp:", tmp)
	// 左子数组的起始索引和结束索引
	leftStart, leftEnd := left-left, mid-left
	// 右子数组的起始索引和结束索引
	rightStart, rightEnd := mid+1-left, right-left
	fmt.Println("leftStart:", leftStart, "leftEnd:", leftEnd, "rightStart:", rightStart, "rightEnd:", rightEnd)
	fmt.Println()
	// i, j 分别指向左子数组、右子数组的首元素
	i, j := leftStart, rightStart
	// 通过覆盖原数组 nums 来合并左子数组和右子数组
	for k := left; k <= right; k++ {
		if i > leftEnd {
			// 若 '左子数组已全部合并完' , 则选取右子数组元素, 并且 j++
			nums[k] = tmp[j]
			j++
		} else if j > rightEnd || tmp[i] <= tmp[j] {
			// 否则, 若 '右子数组已全部合并完' 或 '左子数组元素 <= 右子数组元素' , 则选取左子数组元素, 并且 i++
			nums[k] = tmp[i]
			i++
		} else {
			// 否则, 若 '左右子数组都未全部合并完' 且 '左子数组元素 > 右子数组元素' , 则选取右子数组元素, 并且 j++
			nums[k] = tmp[j]
			j++
		}
	}
}
