/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	m := map[int]int{}
	m[0] = 1
	ans, pre := 0, 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			ans += m[pre-k]
		}
		m[pre]++
	}
	return ans
}

// @lc code=end

