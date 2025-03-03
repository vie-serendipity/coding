/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
func minCut(s string) int {
	n := len(s)

	table := make([][]bool, n)
	for i := 0; i < n; i++ {
		table[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			table[i][j] = true
		}
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] != s[j] || !table[i+1][j-1] {
				table[i][j] = false
			}
		}
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if table[0][i] {
			dp[i] = 0
		} else {
			tmp := math.MaxInt
			for j := 0; j < i; j++ {
				if table[j+1][i] {
					tmp = min(dp[j]+1, tmp)
				}
			}
			dp[i] = tmp
		}
	}

	return dp[n-1]
}

// @lc code=end

