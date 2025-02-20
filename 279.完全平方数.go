/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		minn := math.MaxInt
		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	}
	return dp[n]
}

// @lc code=end

