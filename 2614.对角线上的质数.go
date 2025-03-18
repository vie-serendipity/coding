/*
 * @lc app=leetcode.cn id=2614 lang=golang
 *
 * [2614] 对角线上的质数
 */

// @lc code=start
func diagonalPrime(nums [][]int) int {
	ans := 0
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		if prime(nums[i][i]) {
			ans = max(ans, nums[i][i])
		}
		if prime(nums[n-1-i][i]) {
			ans = max(ans, nums[n-i-1][i])
		}
	}
	return ans
}

func prime(num int) bool {
	if num == 1 {
		return false
	}
	sqrt := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// @lc code=end

