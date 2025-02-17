/*
 * @lc app=leetcode.cn id=1287 lang=golang
 *
 * [1287] 有序数组中出现次数超过25%的元素
 */

// @lc code=start
func findSpecialInteger(arr []int) int {
	threshold := len(arr) / 4
	cnt := 1
	ans := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			cnt++
			if cnt > threshold {
				ans = arr[i]
				return ans
			}
		} else {
			cnt = 1
		}
	}
	return ans
}

// @lc code=end

