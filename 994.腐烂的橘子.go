/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 */

// @lc code=start
func orangesRotting(grid [][]int) int {
	nodes := make([][]int, 0)
	fresh := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				nodes = append(nodes, []int{i, j})
			}
			if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	ans := bfs(grid, nodes, fresh)
	return ans
}

func bfs(grid [][]int, nodes [][]int, fresh int) int {
	cur := 0
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, -1, 0, 1}
	for ; len(nodes) != 0; cur++ {
		q := len(nodes)
		for i := 0; i < q; i++ {
			row := nodes[0][0]
			col := nodes[0][1]
			nodes = nodes[1:]
			for k := 0; k < 4; k++ {
				r := row + dr[k]
				c := col + dc[k]
				if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == 1 {
					fresh--
					grid[r][c] = 2
					nodes = append(nodes, []int{r, c})
				}
			}
			grid[row][col] = 0
		}
	}
	if fresh != 0 {
		return -1
	}
	return cur - 1
}

// @lc code=end

