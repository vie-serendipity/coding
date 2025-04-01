/*
 * @lc app=leetcode.cn id=2140 lang=golang
 *
 * [2140] 解决智力问题
 */

// @lc code=start
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = make([]int, 2)
		if i+1 < n {
			dp[i][0] = max(dp[i+1][0], dp[i+1][1])
		}

		if i+questions[i][1]+1 < n && i+1 < n {
			dp[i][1] = max(questions[i][0]+dp[i+questions[i][1]+1][1], questions[i][0]+dp[i+1][0])
		} else {
			dp[i][1] = questions[i][0]
		}
	}
	return int64(max(dp[0][0], dp[0][1]))
}

// @lc code=end

