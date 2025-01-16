/*
 * @lc app=leetcode.cn id=3095 lang=golang
 *
 * [3095] 或值至少 K 的最短子数组 I
 */

// @lc code=start
func minimumSubarrayLength(nums []int, k int) int {
	if k <= 0 {
		return 1
	}
	count := make([]int, 6)
	ans := 51
	left, right := 0, 0
	for left < len(nums) {
		for right < len(nums) && valueOf(count) < k {
			orOperation(nums[right], count)
			right++
		}
		if valueOf(count) >= k {
			ans = min(ans, right-left)
		}
		remove(nums[left], count)
		left++
	}
	if ans == 51 {
		return -1
	}
	return ans
}

func orOperation(num int, count []int) {
	pos := 0
	for num > 0 {
		if num&1 == 1 {
			count[pos]++
		}
		num >>= 1
		pos++
	}
}

func remove(num int, count []int) {
	pos := 0
	for num > 0 {
		if num&1 == 1 {
			count[pos]--
		}
		num >>= 1
		pos++
	}
}

func valueOf(count []int) int {
	ans := 0
	num := 1
	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			ans += num
		}
		num <<= 1
	}
	return ans
}

// @lc code=end

