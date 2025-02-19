/*
 * @lc app=leetcode.cn id=624 lang=golang
 *
 * [624] 数组列表中的最大距离
 */

// @lc code=start
func maxDistance(arrays [][]int) int {
	maxV, subMaxV := math.MinInt, math.MinInt
	minV, subMinV := math.MaxInt, math.MaxInt
	maxIndex := 0
	minIndex := 0
	for i := 0; i < len(arrays); i++ {
		if arrays[i][len(arrays[i])-1] > maxV {
			subMaxV = max(maxV, subMaxV)
			maxV = arrays[i][len(arrays[i])-1]
			maxIndex = i
		} else {
			subMaxV = max(arrays[i][len(arrays[i])-1], subMaxV)
		}
		if arrays[i][0] < minV {
			subMinV = minV
			minV = arrays[i][0]
			minIndex = i
		} else {
			subMinV = min(arrays[i][0], subMinV)
		}
	}
	if maxIndex != minIndex {
		return maxV - minV
	}
	return max(maxV-subMinV, subMaxV-minV)
}

// @lc code=end

