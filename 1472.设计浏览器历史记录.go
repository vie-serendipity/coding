/*
 * @lc app=leetcode.cn id=1472 lang=golang
 *
 * [1472] 设计浏览器历史记录
 */

// @lc code=start
type BrowserHistory struct {
	history []string
	ptr     int
}

func Constructor(homepage string) BrowserHistory {
	history := []string{homepage}
	return BrowserHistory{
		history: history,
		ptr:     0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.history = this.history[:this.ptr+1]
	this.history = append(this.history, url)
	this.ptr++
}

func (this *BrowserHistory) Back(steps int) string {
	if this.ptr-steps < 0 {
		this.ptr = 0
		return this.history[0]
	}
	this.ptr -= steps
	return this.history[this.ptr]
}

func (this *BrowserHistory) Forward(steps int) string {
	cap := len(this.history)
	if this.ptr+steps >= cap {
		this.ptr = cap - 1
		return this.history[this.ptr]
	}
	this.ptr += steps
	return this.history[this.ptr]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
// @lc code=end

