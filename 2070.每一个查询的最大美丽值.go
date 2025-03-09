/*
 * @lc app=leetcode.cn id=2070 lang=golang
 *
 * [2070] 每一个查询的最大美丽值
 */

// @lc code=start
func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] < items[j][0] {
			return true
		} else if items[i][0] > items[j][0] {
			return false
		} else {
			return items[i][1] > items[j][1]
		}
	})
	for i := 1; i < len(items); i++ {
		if items[i][1] < items[i-1][1] {
			items[i][1] = items[i-1][1]
		}
	}

	ans := make([]int, len(queries))
	for i, key := range queries {
		ans[i] = query(items, key)
	}
	return ans
}

func query(items [][]int, key int) int {
	if key < items[0][0] {
		return 0
	}
	left, right := 0, len(items)-1

	for left <= right {
		mid := left + (right-left)/2
		if items[mid][0] <= key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return items[right][1]
}

// @lc code=end

