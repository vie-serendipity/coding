/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		cur := key(strs[i])
		if len(hash[cur]) == 0 {
			hash[cur] = make([]string, 0)
		}
		hash[cur] = append(hash[cur], strs[i])
	}
	ans := make([][]string, len(hash))
	for i := 0; i < len(hash); i++ {
		ans[i] = make([]string, 0)
	}
	count := 0
	for k := range hash {
		ans[count] = append(ans[count], hash[k]...)
		count++
	}
	return ans
}

func key(a string) string {
	splitA := strings.Split(a, "")
	sort.Strings(splitA)
	return strings.Join(splitA, "")
}
// @lc code=end

