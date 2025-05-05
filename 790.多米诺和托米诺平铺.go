/*
 * @lc app=leetcode.cn id=790 lang=golang
 *
 * [790] 多米诺和托米诺平铺
 */

// @lc code=start
func numTilings(n int) int {
	const mod int = 1e9 + 7
	dp := make([][4]int, n+1)
	dp[0][3] = 1
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][3]
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % mod
		dp[i][2] = (dp[i-1][0] + dp[i-1][1]) % mod
		dp[i][3] = (((dp[i-1][0]+dp[i-1][1])%mod+dp[i-1][2])%mod + dp[i-1][3]) % mod
	}
	return dp[n][3]
}

// @lc code=end

