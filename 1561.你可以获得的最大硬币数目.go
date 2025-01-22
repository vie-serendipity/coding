/*
 * @lc app=leetcode.cn id=1561 lang=golang
 *
 * [1561] 你可以获得的最大硬币数目
 */

// @lc code=start
func maxCoins(piles []int) int {
	sort.Slice(piles, func(i, j int) bool {
		return piles[i] > piles[j]
	})
	ans := 0
	n := len(piles)
	group := n / 3
	for i := 0; i < group; i++ {
		ans += piles[i*2+1]
	}
	return ans
}
// @lc code=end

