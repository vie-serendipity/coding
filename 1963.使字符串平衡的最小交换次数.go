/*
 * @lc app=leetcode.cn id=1963 lang=golang
 *
 * [1963] 使字符串平衡的最小交换次数
 */

// @lc code=start
func minSwaps(s string) int {
	pos := make([]int, 0)
	for i := 0; i < len(s); i++ {
		pos = append(pos, i)
	}

	closing := 0
	opening := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			opening++
		} else {
			closing++
		}
		if closing > opening {
			pos = pos[:len(pos)-1]
			opening++
			closing--
			ans++
		}
	}
	return ans
}

// @lc code=end

