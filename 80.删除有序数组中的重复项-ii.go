/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	cnt := 1
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			cnt++
			if cnt > 2 {
				continue
			}
		} else {
			cnt = 1
		}
		nums[index] = nums[i]
		index++
	}
	nums = nums[:index]
	return len(nums)
}
// @lc code=end

