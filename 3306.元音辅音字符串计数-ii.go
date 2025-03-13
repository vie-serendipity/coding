/*
 * @lc app=leetcode.cn id=3306 lang=golang
 *
 * [3306] 元音辅音字符串计数 II
 */

// @lc code=start
func countOfSubstrings(word string, k int) int64 {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	count := func(m int) int {
		n := len(word)
		var res int = 0
		consonants := 0
		occur := make(map[byte]int)
		for i, j := 0, 0; i < n; i++ {
			for j < n && (consonants < m || len(occur) < 5) {
				if vowels[word[j]] {
					occur[word[j]]++
				} else {
					consonants++
				}
				j++
			}
			if consonants >= m && len(occur) == 5 {
				res += n - j + 1
			}
			if vowels[word[i]] {
				occur[word[i]]--
				if occur[word[i]] == 0 {
					delete(occur, word[i])
				}
			} else {
				consonants--
			}
		}
		return res
	}
	return int64(count(k) - count(k+1))
}
// @lc code=end

