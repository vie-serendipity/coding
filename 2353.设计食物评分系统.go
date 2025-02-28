/*
 * @lc app=leetcode.cn id=2353 lang=golang
 *
 * [2353] 设计食物评分系统
 */

// @lc code=start

type FoodRatings struct {
	foodMap map[string]struct {
		rating  int
		cuisine string
	}
	ratingMap map[string]*PriorityQueue
	n         int
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	n := len(foods)
	foodMap := make(map[string]struct {
		rating  int
		cuisine string
	})
	ratingMap := make(map[string]*PriorityQueue)

	for i := 0; i < n; i++ {
		food := foods[i]
		cuisine := cuisines[i]
		rating := ratings[i]
		foodMap[food] = struct {
			rating  int
			cuisine string
		}{rating, cuisine}
		if ratingMap[cuisine] == nil {
			ratingMap[cuisine] = &PriorityQueue{}
		}
		heap.Push(ratingMap[cuisine], Pair{n - rating, food})
	}

	return FoodRatings{
		foodMap:   foodMap,
		ratingMap: ratingMap,
		n:         n,
	}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	entry := this.foodMap[food]
	cuisine := entry.cuisine
	heap.Push(this.ratingMap[cuisine], Pair{this.n - newRating, food})
	entry.rating = newRating
	this.foodMap[food] = entry
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	pq := this.ratingMap[cuisine]
	for pq.Len() > 0 {
		top := (*pq)[0]
		if this.n-top.rating == this.foodMap[top.food].rating {
			return top.food
		}
		heap.Pop(pq)
	}
	return ""
}

type Pair struct {
	rating int
	food   string
}

type PriorityQueue []Pair

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].rating == pq[j].rating {
		return pq[i].food < pq[j].food
	}
	return pq[i].rating < pq[j].rating
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Pair))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
// @lc code=end

