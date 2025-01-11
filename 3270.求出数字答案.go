/*
 * @lc app=leetcode.cn id=3270 lang=golang
 *
 * [3270] 求出数字答案
 */

// @lc code=start
func generateKey(num1 int, num2 int, num3 int) int {
	ans := 0
	for p := 1; num1 > 0 && num2 > 0 && num3 > 0; p *= 10 {
		ans += min(num1%10, min(num2%10, num3%10)) * p
		num1, num2, num3 = num1/10, num2/10, num3/10
	}
	return ans
}

// @lc code=end

