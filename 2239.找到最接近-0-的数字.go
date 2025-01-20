/*
 * @lc app=leetcode.cn id=2239 lang=golang
 *
 * [2239] 找到最接近 0 的数字
 */

// @lc code=start
func findClosestNumber(nums []int) int {
	ans := -100001
	closest := 100001
	for i := range nums {
		if nums[i] >= 0 && (nums[i]-0) < closest {
			closest = nums[i] - 0
			ans = nums[i]
		} else if nums[i] < 0 && (0-nums[i]) < closest {
			closest = 0 - nums[i]
			ans = nums[i]
		} else if (nums[i] - 0) == closest {
			ans = max(ans, nums[i])
		}
	}
	return ans
}

// @lc code=end

