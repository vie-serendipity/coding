/*
 * @lc app=leetcode.cn id=2829 lang=golang
 *
 * [2829] k-avoiding 数组的最小总和
 */

// @lc code=start
func minimumSum(n int, k int) int {
	hash := make(map[int]bool)
	ans := 0
	for i := 1; len(hash) < n; i++ {
		if _, ok := hash[k-i]; ok {
			continue
		}
		ans += i
		hash[i] = true
	}
	return ans
}

// @lc code=end

