/*
 * @lc app=leetcode.cn id=2269 lang=golang
 *
 * [2269] 找到一个数字的 K 美丽值
 */

// @lc code=start
func divisorSubstrings(num int, k int) int {
	tmp := num
	number := make([]int, 11)
	cnt := 0
	for i := 0; i < 11; i++ {
		number[i] = -1
		if num <= 0 {
			continue
		}
		number[i] = num % 10
		num /= 10
		cnt++
	}
	ans := 0
	for i := k - 1; number[i] != -1; i++ {
		value := 0
		for j := i; j >= i-(k-1); j-- {
			value = value*10 + number[j]
		}
		if value == 0 {
			continue
		}
		if tmp%value == 0 {
			ans++
		}
	}
	return ans
}

// @lc code=end

