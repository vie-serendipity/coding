/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] <= intervals[j][1]
		} else {
			return intervals[i][0] <= intervals[j][0]
		}
	})
	var ans [][]int
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i+1][0] = intervals[i][0]
			if intervals[i][1] > intervals[i+1][1] {
				intervals[i+1][1] = intervals[i][1]
			}
		} else {
			ans = append(ans, intervals[i])
		}
	}
	ans = append(ans, intervals[len(intervals)-1])
	return ans
}

// @lc code=end

