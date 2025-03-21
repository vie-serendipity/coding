/*
 * @lc app=leetcode.cn id=2680 lang=golang
 *
 * [2680] 最大或值
 */

// @lc code=start
func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	if n == 1 {
		var ans int64
		ans = int64(nums[0])
		for i := 0; i < k; i++ {
			ans = ans * 2
		}
		return ans
	}
	prefix := make([]int64, n)
	sufix := make([]int64, n)
	var v int64
	v = 0
	for i := 0; i < n; i++ {
		prefix[i] = v | int64(nums[i])
		v = prefix[i]
	}
	v = 0
	for i := n - 1; i >= 0; i-- {
		sufix[i] = v | int64(nums[i])
		v = sufix[i]
	}

	var ans int64
	ans = -1
	for i := 0; i < n; i++ {
		var tmp int64
		tmp = int64(nums[i])
		for j := 0; j < k; j++ {
			tmp *= 2
		}
		if i == 0 {
			ans = max(ans, tmp|int64(sufix[i+1]))
		} else if i == n-1 {
			ans = max(ans, int64(prefix[i-1])|tmp)
		} else {
			ans = max(ans, int64(prefix[i-1])|tmp|int64(sufix[i+1]))
		}
	}

	return ans
}

// @lc code=end

