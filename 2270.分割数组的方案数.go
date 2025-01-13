/*
 * @lc app=leetcode.cn id=2270 lang=golang
 *
 * [2270] 分割数组的方案数
 */

// @lc code=start
func waysToSplitArray(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	ans := 0
	left, right := nums[0], sum-nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if left >= right {
			ans++
		}
		left += nums[i+1]
		right -= nums[i+1]
	}
	return ans
}

// @lc code=end

