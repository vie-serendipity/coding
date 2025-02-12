/*
 * @lc app=leetcode.cn id=1760 lang=golang
 *
 * [1760] 袋子里最少数目的球
 */

// @lc code=start
func minimumSize(nums []int, maxOperations int) int {
	sort.Ints(nums)
	left, right := 1, nums[len(nums)-1]
	ans := nums[0]
	for left <= right {
		mid := (left + right) / 2
		if valid(nums, mid, maxOperations) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func valid(nums []int, v int, maxOperations int) bool {
	for i := len(nums) - 1; i >= 0 && nums[i] > v; i-- {
		ops := (nums[i] - 1) / v
		maxOperations -= ops
		if maxOperations < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

