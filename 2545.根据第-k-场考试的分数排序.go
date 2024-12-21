/*
 * @lc app=leetcode.cn id=2545 lang=golang
 *
 * [2545] 根据第 K 场考试的分数排序
 */

// @lc code=start
func sortTheStudents(score [][]int, k int) [][]int {
	numOfStudent := len(score)
	if numOfStudent == 0 {
		return [][]int{}
	}
	if len(score[0]) == 0 {
		return [][]int{}
	}
	scoreCopy := make([][]int, numOfStudent)
	order := make([]int, len(score))

	for i := range scoreCopy {
		scoreCopy[i] = make([]int, len(score[i]))
		order[i] = i
		for j := range score[i] {
			scoreCopy[i][j] = score[i][j]
		}
	}

	// high to low
	for i := 0; i < numOfStudent-1; i++ {
		for j := 0; j < numOfStudent-1-i; j++ {
			if score[j][k] < score[j+1][k] {
				score[j][k], score[j+1][k] = score[j+1][k], score[j][k]
				order[j], order[j+1] = order[j+1], order[j]
			}
		}
	}
	for i := range score {
		score[i] = scoreCopy[order[i]]
	}
	return score
}
// @lc code=end

