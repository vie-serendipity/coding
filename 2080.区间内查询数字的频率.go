/*
 * @lc app=leetcode.cn id=2080 lang=golang
 *
 * [2080] 区间内查询数字的频率
 */

// @lc code=start
type RangeFreqQuery struct {
	table map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	table := make(map[int][]int)
	for k, v := range arr {
		if table[v] == nil {
			table[v] = make([]int, 0)
		}
		table[v] = append(table[v], k)
	}
	return RangeFreqQuery{
		table: table,
	}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	if _, ok := this.table[value]; !ok {
		return 0
	}
	indices := this.table[value]

	l := lowerBound(indices, left)
	r := upperBound(indices, right)

	return r - l
}

func lowerBound(pos []int, target int) int {
	low, high := 0, len(pos)-1
	for low <= high {
		mid := low + (high-low)/2
		if pos[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func upperBound(pos []int, target int) int {
	low, high := 0, len(pos)-1
	for low <= high {
		mid := low + (high-low)/2
		if pos[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
// @lc code=end

