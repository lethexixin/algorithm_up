package sort

import (
	"sort"
)

// BucketSort 桶排序
func BucketSort(nums []float64) {
	k := len(nums) / 2

	buckets := make([][]float64, k)
	for i := 0; i < len(buckets); i++ {
		buckets[i] = make([]float64, 0)
	}

	for _, num := range nums {
		i := int(num) * k
		buckets[i] = append(buckets[i], num)
	}

	for i := 0; i < len(buckets); i++ {
		sort.Float64s(buckets[i])
	}

	i := 0
	for _, bucket := range buckets {
		for _, n := range bucket {
			nums[i] = n
			i++
		}
	}
}
