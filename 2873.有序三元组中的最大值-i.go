/*
 * @lc app=leetcode.cn id=2873 lang=golang
 *
 * [2873] 有序三元组中的最大值 I
 */

// @lc code=start
func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	var ans int64
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				ans = max(int64((nums[i]-nums[j])*nums[k]), ans)
			}
		}
	}
	return ans
}

// @lc code=end

