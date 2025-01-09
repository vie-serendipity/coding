/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [3297] 统计重新排列后包含另一个字符串的子字符串数目 I
 */

// @lc code=start
func validSubstringCount(word1 string, word2 string) int64 {
	var ans int64
	count := make(map[byte]int)
	for k := range word2 {
		count[word2[k]] = count[word2[k]] - 1
	}

	left := 0
	for i := range word1 {
		count[word1[i]] += 1
		for valid(count) {
			count[word1[left]] -= 1
			left += 1
		}
		ans += int64(left)
	}
	return ans
}

func valid(count map[byte]int) bool {
	for _, v := range count {
		if v < 0 {
			return false
		}
	}
	return true
}