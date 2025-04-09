/*
 * @lc app=leetcode.cn id=3375 lang=golang
 *
 * [3375] 使数组的值全部为 K 的最少操作次数
 */

// @lc code=start
func minOperations(nums []int, k int) int {
	n := len(nums)
	table := make(map[int]bool)
	for i := 0; i < n; i++ {
		if nums[i] > k {
			table[nums[i]] = true
		}
		if nums[i] < k {
			return -1
		}
	}
	return len(table)
}

// @lc code=end

