/*
 * @lc app=leetcode.cn id=2944 lang=golang
 *
 * [2944] 购买水果需要的最少金币数
 */

// @lc code=start
func minimumCoins(prices []int) int {
	state := make(map[int]int)
	var dfs func(int) int
	dfs = func(i int) int {
		if 2*i+2 >= len(prices) {
			return prices[i]
		}
		if v, ok := state[i]; ok {
			return v
		}

		minValue := math.MaxInt

		for j := i + 1; j <= 2*i+2; j++ {
			minValue = min(minValue, dfs(j))
		}
		state[i] = prices[i] + minValue
		return state[i]
	}
	return dfs(0)
}

// @lc code=end

