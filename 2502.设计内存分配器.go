/*
 * @lc app=leetcode.cn id=2502 lang=golang
 *
 * [2502] 设计内存分配器
 */

// @lc code=start
type Allocator struct {
	memory   []int
	capacity int
}

func Constructor(n int) Allocator {
	return Allocator{
		memory:   make([]int, n),
		capacity: n,
	}
}

func (this *Allocator) Allocate(size int, mID int) int {
	cnt := 0
	var i int
	for i = 0; i < this.capacity; i++ {
		if this.memory[i] == 0 {
			cnt++
		} else {
			cnt = 0
		}
		if cnt == size {
			break
		}
	}

	if i == this.capacity {
		return -1
	}

	for j := i; j > i-size; j-- {
		this.memory[j] = mID
	}

	return i - size + 1
}

func (this *Allocator) FreeMemory(mID int) int {
	ans := 0
	for i := 0; i < this.capacity; i++ {
		if this.memory[i] == mID {
			this.memory[i] = 0
			ans++
		}
	}
	return ans
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.FreeMemory(mID);
 */
// @lc code=end

