/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivot := partition(nums, lo, hi)
	sort(nums, lo, pivot-1)
	sort(nums, pivot+1, hi)
}

func partition(nums []int, lo, hi int) int {
	pivot := lo
	lo = lo + 1
	for lo <= hi {
		for lo <= hi && nums[lo] < nums[pivot] {
			lo++
		}
		for lo <= hi && nums[hi] >= nums[pivot] {
			hi--
		}
		if lo >= hi {
			break
		}
		nums[lo], nums[hi] = nums[hi], nums[lo]
	}
	nums[pivot], nums[hi] = nums[hi], nums[pivot]
	return hi
}

// 1 2 3 5

// @lc code=end
