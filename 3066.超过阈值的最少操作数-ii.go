/*
 * @lc app=leetcode.cn id=3066 lang=golang
 *
 * [3066] 超过阈值的最少操作数 II
 */

// @lc code=start
func minOperations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := 0
	for {
		if len(nums) < 2 {
			return ans
		}
		newNum := min(nums[0], nums[1])*2 + max(nums[0], nums[1])
		nums = append([]int{}, nums[2:]...)
		for i, num := range nums {
			if newNum < num {
				nums = append(nums[:i], append([]int{newNum}, nums[i:]...)...)
				break
			}
		}
		ans++
		if nums[0] >= k {
			break
		}
	}
	return ans
}

// @lc code=end

