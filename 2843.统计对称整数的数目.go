/*
 * @lc app=leetcode.cn id=2843 lang=golang
 *
 * [2843] 统计对称整数的数目
 */

// @lc code=start
func countSymmetricIntegers(low int, high int) int {
	ans := 0
	for i := low; i <= high; i++ {
		if valid(strconv.Itoa(i)) {
			ans++
		}
	}
	return ans
}

func valid(num string) bool {
	if len(num)%2 == 1 {
		return false
	}
	left, right := 0, 0
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		left += int(num[i] - '0')
		right += int(num[j] - '0')
	}
	return left == right
}

// @lc code=end

