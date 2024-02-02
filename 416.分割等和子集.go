/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum := 0
	maxVal := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if maxVal > target {
		return false
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}

// @lc code=end

