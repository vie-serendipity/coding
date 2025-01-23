/*
 * @lc app=leetcode.cn id=2920 lang=golang
 *
 * [2920] 收集所有金币可获得的最大积分
 */

// @lc code=start
func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for i := range edges {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	isVisited := make([]int, n)
	var dfs func(int)
	dfs = func(u int) {
		if isVisited[u] != 0 {
			return
		}
		isVisited[u] = 1
		for _, v := range graph[u] {
			if isVisited[v] == 0 {
				dfs(v)
				dp[u][0] += max(dp[v][0], dp[v][1])
				dp[u][1] += dp[v][1] / 2
			}
		}

		dp[u][0] += coins[u] - k
		dp[u][1] += coins[u] / 2
	}

	dfs(0)
	return max(dp[0][0], dp[0][1])
}


// @lc code=end

