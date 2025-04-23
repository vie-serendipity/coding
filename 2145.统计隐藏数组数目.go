/*
 * @lc app=leetcode.cn id=2145 lang=golang
 *
 * [2145] 统计隐藏数组数目
 */

// @lc code=start
func numberOfArrays(differences []int, lower int, upper int) int {
	left, right := lower, upper
	ans := upper - lower + 1
	for i := 0; i < len(differences); i++ {
		left = max(left+differences[i], lower)
		right = min(differences[i]+right, upper)
		ans = min(ans, right-left+1)
	}
	if ans < 0 {
		return 0
	}
	return ans
}

// @lc code=end

