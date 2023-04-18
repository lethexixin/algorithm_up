package o40

//剑指 Offer 40. 最小的 k 个数
//输入整数数组 arr ，找出其中最小的 k 个数。例如，
//输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//
//
//示例 1：
//
//输入：arr = [3,2,1], k = 2
//输出：[1,2] 或者 [2,1]
//示例 2：
//
//输入：arr = [0,1,2,1], k = 1
//输出：[0]
//
//限制：
//
//0 <= k <= arr.length <= 10000
//0 <= arr[i]<= 10000
//
//作者：Krahets
//链接：https://leetcode.cn/leetbook/read/illustration-of-algorithm/ohvl0d/

// 执行用时： 28 ms, 在所有 Go 提交中击败了 54.12% 的用户
// 内存消耗： 6.4 MB, 在所有 Go 提交中击败了 37.23% 的用户
func getLeastNumbers(arr []int, k int) []int {
	if k >= len(arr) {
		return arr
	}

	quickSort(arr, 0, len(arr)-1)
	return arr[:k]
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	l, r := left, right
	for l < r {
		//两个循环顺序不能反, 先 r-- 后 l++
		for l < r && arr[r] >= arr[left] {
			r--
		}
		for l < r && arr[l] <= arr[left] {
			l++
		}

		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[l], arr[left] = arr[left], arr[l]
	quickSort(arr, left, l-1)
	quickSort(arr, l+1, right)
}
