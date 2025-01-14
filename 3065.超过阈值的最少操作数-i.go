/*
 * @lc app=leetcode.cn id=3065 lang=golang
 *
 * [3065] 超过阈值的最少操作数 I
 */

// @lc code=start
func minOperations(nums []int, k int) int {
	ans := 0
	for i := range nums {
		if nums[i] < k {
			ans++
		}
	}
	return ans
}

// @lc code=end

