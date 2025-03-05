/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	minV := nums[0]
	maxV := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		tmpMax, tmpMin := maxV, minV
		minV = min(nums[i], min(tmpMin*nums[i], tmpMax*nums[i]))
		maxV = max(nums[i], max(tmpMax*nums[i], tmpMin*nums[i]))
		ans = max(maxV, ans)
	}
	return ans
}

// @lc code=end

