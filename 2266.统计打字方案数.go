/*
 * @lc app=leetcode.cn id=2266 lang=golang
 *
 * [2266] 统计打字方案数
 */

// @lc code=start
var solutionThree = []int{1, 1, 2, 4}
var solutionFour = []int{1, 1, 2, 4, 8}

func countTexts(pressedKeys string) int {
	array := make([]int, 0)
	cnt := 1
	for i := 0; i < len(pressedKeys); i++ {
		if i < len(pressedKeys)-1 && pressedKeys[i] == pressedKeys[i+1] {
			cnt++
		} else {
			array = append(array, consecutive(cnt, pressedKeys[i])%(1000000000+7))
			cnt = 1
		}
	}
	ans := 1

	for i := 0; i < len(array); i++ {
		ans = (ans * array[i]) % (1000000000 + 7)
	}
	return ans
}

func consecutive(num int, alpha byte) int {
	if alpha == '7' || alpha == '9' {
		if len(solutionFour) > num {
			return solutionFour[num]
		} else {
			for i := len(solutionFour); i <= num; i++ {
				solutionFour = append(solutionFour, solutionFour[i-1]%(1000000000+7)+solutionFour[i-2]%(1000000000+7)+solutionFour[i-3]%(1000000000+7)+solutionFour[i-4]%(1000000000+7))
			}
			return solutionFour[num]
		}
	}
	if len(solutionThree) > num {
		return solutionThree[num]
	} else {
		for i := len(solutionThree); i <= num; i++ {
			solutionThree = append(solutionThree, solutionThree[i-1]%(1000000000+7)+solutionThree[i-2]%(1000000000+7)+solutionThree[i-3]%(1000000000+7))
		}
	}
	return solutionThree[num]
}

// @lc code=end

