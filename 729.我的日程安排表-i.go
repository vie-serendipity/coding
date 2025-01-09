/*
 * @lc app=leetcode.cn id=729 lang=golang
 *
 * [729] 我的日程安排表 I
 */

// @lc code=start
type MyCalendar struct {
	tasks [][]int
}

func Constructor() MyCalendar {
	tasks := make([][]int, 0)
	return MyCalendar{
		tasks: tasks,
	}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	for i := range this.tasks {
		if startTime >= this.tasks[i][1] || endTime <= this.tasks[i][0] {
			continue
		} else {
			return false
		}
	}

	this.tasks = append(this.tasks, []int{startTime, endTime})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
// @lc code=end

