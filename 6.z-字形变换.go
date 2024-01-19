/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	ans := ""
	for i := 0; i < numRows; i++ {
		t := 2*numRows - 2
		for j := 0; j+i < len(s); j += t {
			ans += string(s[j+i])
			if i > 0 && i < numRows-1 && j+t-i < len(s) {
				ans += string(s[j+t-i])
			}
		}
	}
	return ans
}

// @lc code=end

