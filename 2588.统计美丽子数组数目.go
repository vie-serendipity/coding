/*
 * @lc app=leetcode.cn id=2588 lang=golang
 *
 * [2588] 统计美丽子数组数目
 */

// @lc code=start
func beautifulSubarrays(nums []int) int64 {
	cnt := make(map[int]int)
	v := 0
	var ans int64
	cnt[0] = 1
	for i := range nums {
		v ^= nums[i]
		ans += int64(cnt[v])
		cnt[v]++
	}
	return ans
}

// @lc code=end

