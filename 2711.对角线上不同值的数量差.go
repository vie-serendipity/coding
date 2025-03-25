/*
 * @lc app=leetcode.cn id=2711 lang=golang
 *
 * [2711] 对角线上不同值的数量差
 */

// @lc code=start
func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			s1 := make(map[int]bool)
			x, y := i+1, j+1
			for x < m && y < n {
				s1[grid[x][y]] = true
				x++
				y++
			}
			s2 := make(map[int]bool)
			x, y = i-1, j-1
			for x >= 0 && y >= 0 {
				s2[grid[x][y]] = true
				x--
				y--
			}
			res[i][j] = abs(len(s1) - len(s2))
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

