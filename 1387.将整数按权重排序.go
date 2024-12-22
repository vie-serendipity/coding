/*
 * @lc app=leetcode.cn id=1387 lang=golang
 *
 * [1387] 将整数按权重排序
 */

// @lc code=start
var kv = make(map[int]int)

func getKth(lo int, hi int, k int) int {
	weights := make([]int, 0)
	indices := make([]int, 0)
	for i := lo; i <= hi; i++ {
		weights = append(weights, getWeight(i))
		indices = append(indices, i)
	}
	return sort(weights, indices, 0, len(weights)-1, k-1)
}

func getWeight(x int) int {
	if weight, ok := kv[x]; ok {
		return weight
	}
	if x == 1 {
		kv[0] = 0
		return 0
	}

	if x%2 == 0 {
		kv[x] = getWeight(x/2) + 1
	} else {
		kv[x] = getWeight(3*x+1) + 1
	}
	return kv[x]
}

func sort(nums, indices []int, lo, hi int, k int) int {
	if lo >= hi {
		return indices[lo]
	}
	pivot := partition(nums, indices, lo, hi)
	if pivot == k {
		return indices[pivot]
	} else if pivot > k {
		return sort(nums, indices, lo, pivot-1, k)
	} else {
		return sort(nums, indices, pivot+1, hi, k)
	}
}

func partition(nums, indices []int, lo, hi int) int {
	pivot := lo
	lo = lo + 1
	for lo <= hi {
		for lo <= hi && less(nums, indices, lo, pivot) {
			lo++
		}
		for lo <= hi && !less(nums, indices, hi, pivot) {
			hi--
		}
		if lo >= hi {
			break
		}
		nums[lo], nums[hi] = nums[hi], nums[lo]
		indices[lo], indices[hi] = indices[hi], indices[lo]
	}
	nums[pivot], nums[hi] = nums[hi], nums[pivot]
	indices[pivot], indices[hi] = indices[hi], indices[pivot]
	return hi
}

func less(nums, indices []int, i, j int) bool {
	if nums[i] < nums[j] {
		return true
	} else if nums[i] > nums[j] {
		return false
	} else {
		return indices[i] < indices[j]
	}
}

// @lc code=end

