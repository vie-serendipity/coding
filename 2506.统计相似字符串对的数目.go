/*
 * @lc app=leetcode.cn id=2506 lang=golang
 *
 * [2506] 统计相似字符串对的数目
 */

// @lc code=start
func similarPairs(words []string) int {
	ans := 0
	cnt := make(map[int]int)

	for _, word := range words {
		state := 0
		for _, c := range word {
			state |= 1 << (c - 'a')
		}
		ans += cnt[state]
		cnt[state]++
	}
	return ans
}

// @lc code=end

