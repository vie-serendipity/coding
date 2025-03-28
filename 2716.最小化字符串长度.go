/*
 * @lc app=leetcode.cn id=2716 lang=golang
 *
 * [2716] 最小化字符串长度
 */

// @lc code=start
func minimizedStringLength(s string) int {
	hash := make(map[string]bool)
	for _, v := range s {
		hash[string(v)] = true
	}
	return len(hash)
}

// @lc code=end

