/*
 * @lc app=leetcode.cn id=1656 lang=golang
 *
 * [1656] 设计有序流
 */

// @lc code=start
type OrderedStream struct {
	ptr   int
	pairs map[int]string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		ptr:   1,
		pairs: make(map[int]string),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.pairs[idKey] = value
	ans := make([]string, 0)
	index := this.ptr
	for {
		if _, ok := this.pairs[index]; !ok {
			break
		}
		ans = append(ans, this.pairs[index])
		index++
	}
	this.ptr = index
	return ans
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
// @lc code=end

