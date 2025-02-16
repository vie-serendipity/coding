/*
 * @lc app=leetcode.cn id=1299 lang=golang
 *
 * [1299] 将每个元素替换为右侧最大元素
 */

// @lc code=start
func replaceElements(arr []int) []int {
	n := len(arr)
	maxV := arr[n-1]
	arr[n-1] = -1

	for i := n - 2; i >= 0; i-- {
		tmp := maxV
		maxV = max(maxV, arr[i])
		arr[i] = tmp
	}
	return arr
}

// @lc code=end

