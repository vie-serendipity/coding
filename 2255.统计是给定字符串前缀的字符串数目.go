/*
 * @lc app=leetcode.cn id=2255 lang=golang
 *
 * [2255] 统计是给定字符串前缀的字符串数目
 */

// @lc code=start
func countPrefixes(words []string, s string) int {
	n := len(words)
	ans := 0
	for i := 0; i < n; i++ {
		if len(words[i]) > len(s) {
			continue
		}
		if words[i] == s[:len(words[i])] {
			ans++
		}
	}
	return ans
}

// @lc code=end

