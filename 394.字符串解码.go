/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	str := ""
	num := 0
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			numStack = append(numStack, num)
			strStack = append(strStack, str)
			num = 0
			str = ""
		} else if 'a' <= s[i] && s[i] <= 'z' {
			str += string(s[i])
		} else {
			num := numStack[len(numStack)-1]
			strTmp := strStack[len(strStack)-1]
			numStack = numStack[:len(numStack)-1]
			strStack = strStack[:len(strStack)-1]
			for j := 0; j < num; j++ {
				strTmp += str
			}
			str = strTmp
		}
	}
	return str
}

// @lc code=end

