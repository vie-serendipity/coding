/*
 * @lc app=leetcode.cn id=1920 lang=golang
 *
 * [1920] 基于排列构建数组
 */

// @lc code=start
func buildArray(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// use initial value
		nums[i] = nums[nums[i]%1000]%1000*1000 + nums[i]%1000
	}
	for i := 0; i < n; i++ {
		// final value is only needed here
		nums[i] = nums[i] / 1000
	}
	return nums
}

// @lc code=end

