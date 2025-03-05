/*
 * @lc app=leetcode.cn id=1328 lang=golang
 *
 * [1328] 破坏回文串
 */

// @lc code=start
func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}
	var ans string
	for i := 0; i < len(palindrome); i++ {
		if len(palindrome)%2 == 1 && i == (len(palindrome)+1)/2 {
			continue
		}
		if palindrome[i] != 'a' {
			tmp := []byte(palindrome)
			tmp[i] = 'a'
			ans = string(tmp)
			return ans
		}
	}
	tmp := []byte(palindrome)
	tmp[len(palindrome)-1] = 'b'
	ans = string(tmp)
	return ans
}

// @lc code=end

