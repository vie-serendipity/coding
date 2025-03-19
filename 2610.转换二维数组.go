/*
 * @lc app=leetcode.cn id=2610 lang=golang
 *
 * [2610] 转换二维数组
 */

// @lc code=start
func findMatrix(nums []int) [][]int {
	n := len(nums)
	cnt := make(map[int]int)
	maxCnt := 0
	for i := 0; i < n; i++ {
		cnt[nums[i]]++
		maxCnt = max(maxCnt, cnt[nums[i]])
	}
	ans := make([][]int, maxCnt)
	for k, v := range cnt {
		for i := 0; i < v; i++ {
			if ans[i] == nil {
				ans[i] = make([]int, 0)
			}
			ans[i] = append(ans[i], k)
		}
	}
	return ans
}

// @lc code=end

