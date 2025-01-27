/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	maxPosition := 0
	ans := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(i+nums[i], maxPosition)
		if i == end {
			end = maxPosition
			ans++
		}
	}
	return ans
}
// @lc code=end
