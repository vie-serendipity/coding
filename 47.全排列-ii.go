/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	isvisited := make([]int, len(nums))
	var backtrack func()
	ans := make([][]int, 0)
	element := make([]int, 0)
	backtrack = func() {
		if len(element) == n {
			tmp := make([]int, n)
			copy(tmp, element)
			ans = append(ans, tmp)
		}
		prev := -11
		for i := 0; i < len(nums); i++ {
			if nums[i] == prev {
				continue
			}
			if isvisited[i] == 0 {
				prev = nums[i]
				isvisited[i] = 1
				element = append(element, nums[i])
				backtrack()
				element = element[:len(element)-1]
				isvisited[i] = 0
			}
		}
	}
	backtrack()
	return ans
}

// @lc code=end

