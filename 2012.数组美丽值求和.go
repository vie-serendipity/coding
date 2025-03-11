/*
 * @lc app=leetcode.cn id=2012 lang=golang
 *
 * [2012] 数组美丽值求和
 */

// @lc code=start
func sumOfBeauties(nums []int) int {
	n := len(nums)
	array := make([]int, n)
	maxV := nums[0]
	for i := 1; i < n-1; i++ {
		array[i] = 2
		if nums[i] <= maxV {
			array[i] = 1
		}
		maxV = max(nums[i], maxV)
	}

	minV := nums[n-1]
	for i := n - 2; i >= 1; i-- {
		if nums[i] >= minV {
			array[i] = 1
		}
		if nums[i] <= nums[i-1] || nums[i] >= nums[i+1] {
			array[i] = 0
		}
		minV = min(nums[i], minV)
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += array[i]
	}
	return sum
}


// @lc code=end

