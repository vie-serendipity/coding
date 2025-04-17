/*
 * @lc app=leetcode.cn id=2176 lang=golang
 *
 * [2176] 统计数组中相等且可以被整除的数对
 */

// @lc code=start
func countPairs(nums []int, k int) int {
	index := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := index[nums[i]]; !ok {
			index[nums[i]] = make([]int, 0)
		}
		index[nums[i]] = append(index[nums[i]], i)
	}
	ans := 0
	for _, v := range index {
		ans += findPairs(v, k)
	}
	return ans
}

func findPairs(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]*nums[j]%k == 0 {
				ans++
			}
		}
	}
	return ans
}

// @lc code=end

