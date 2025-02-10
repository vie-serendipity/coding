/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	row, col := 0, 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir := 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		if row+dirs[dir][0] < n &&
			row+dirs[dir][0] >= 0 &&
			col+dirs[dir][1] < n &&
			col+dirs[dir][1] >= 0 &&
			matrix[row+dirs[dir][0]][col+dirs[dir][1]] == 0 {
			row += dirs[dir][0]
			col += dirs[dir][1]
		} else {
			dir = (dir + 1) % 4
			row += dirs[dir][0]
			col += dirs[dir][1]
		}
	}
	return matrix
}

// @lc code=end

