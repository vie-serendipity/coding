/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	vote := 1
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == ans {
			vote++
		} else {
			vote--
		}
		if vote < 0 {
			ans = nums[i]
			vote = 1
		}
	}
	return ans
}

// @lc code=end

