/*
 * @lc app=leetcode.cn id=1399 lang=golang
 *
 * [1399] 统计最大组的数目
 */

// @lc code=start
func countLargestGroup(n int) int {
	sum := make(map[int]int)
	for i := 1; i <= n; i++ {
		sum[sumOf(i)]++
	}
	ans := 0
	cnt := make(map[int]int)
	for _, v := range sum {
		cnt[v]++
	}
	value := 0
	for k, v := range cnt {
		if k > value {
			value = k
			ans = v
		}
	}
	return ans
}

func sumOf(a int) int {
	sum := 0
	for a != 0 {
		sum += (a % 10)
		a = a / 10
	}
	return sum
}

// @lc code=end

