/*
 * @lc app=leetcode.cn id=2264 lang=golang
 *
 * [2264] 字符串中最大的 3 位相同数字
 */

// @lc code=start
func largestGoodInteger(num string) string {
	ans := ""
	if len(num) < 3 {
		return ""
	}
	for i := 0; i+2 < len(num); i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			ans = max(ans, num[i:i+3])
		}
	}
	return ans
}


// @lc code=end

