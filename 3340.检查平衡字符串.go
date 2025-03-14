/*
 * @lc app=leetcode.cn id=3340 lang=golang
 *
 * [3340] 检查平衡字符串
 */

// @lc code=start
func isBalanced(num string) bool {
	odd, even := 0, 0
	for k, v := range num {
		if k%2 == 0 {
			even += int(v - '0')
		} else {
			odd += int(v - '0')
		}
	}
	return odd == even
}

// @lc code=end

