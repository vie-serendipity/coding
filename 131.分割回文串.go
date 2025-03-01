/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start

func partition(s string) [][]string {
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

	ans := [][]string{}

	var backtrack func(string, int, []string)

	backtrack = func(s string, start int, item []string) {
		if start >= len(s) {
			tmp := make([]string, 0)
			tmp = append(tmp, item...)
			ans = append(ans, tmp)
		}
		for i := start; i < len(s); i++ {
			if table[start][i] {
				item = append(item, s[start:i+1])
				backtrack(s, i+1, item)
				item = item[:len(item)-1]
			}
		}
	}
	item := []string{}
	backtrack(s, 0, item)
	return ans
}

// @lc code=end

