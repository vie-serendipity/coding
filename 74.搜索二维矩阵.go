/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}

	low, high := 0, m-1
	for low <= high {
		mid := low + (high-low)/2
		if target < matrix[mid][0] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	// high
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[high][mid] == target {
			return true
		}
		if target < matrix[high][mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// @lc code=end

