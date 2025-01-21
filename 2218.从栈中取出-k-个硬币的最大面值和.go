/*
 * @lc app=leetcode.cn id=2218 lang=golang
 *
 * [2218] 从栈中取出 K 个硬币的最大面值和
 */

// @lc code=start
func maxValueOfCoins(piles [][]int, k int) int {
	prefixSum := make([][]int, len(piles)+1)
	prefixSum[0] = make([]int, k+1)
	for i := 1; i <= len(piles); i++ {
		prefixSum[i] = make([]int, k+1)
		sum := 0
		j := 0
		for j = 1; j <= len(piles[i-1]) && j <= k; j++ {
			sum += piles[i-1][j-1]
			prefixSum[i][j] = sum
		}
		for j <= k {
			prefixSum[i][j] = sum
			j++
		}
	}

	dp := make([][]int, len(piles)+1)
	dp[0] = make([]int, k+1)
	for i := 1; i <= len(piles); i++ {
		dp[i] = make([]int, k+1)
		for j := 1; j <= k; j++ {
			value := 0
			for m := 0; m <= len(piles[i-1]) && m <= j; m++ {
				value = max(dp[i-1][j-m]+prefixSum[i][m], value)
			}
			dp[i][j] = max(value, prefixSum[i][j])
		}
	}

	return dp[len(piles)][k]
}

// @lc code=end
