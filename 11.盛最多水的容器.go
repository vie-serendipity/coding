/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
    ans:=0
    left,right:=0, len(height)-1
    ans=max(ans, min(height[left], height[right])*(right-left))
    for left<right{
        if height[left]<height[right]{
            left++
        }else{
            right--
        }
        ans=max(ans, min(height[left], height[right])*(right-left))
    }
    return ans
}

func max(a, b int) int {
    if a>b{
        return a
    }
    return b
}
func min(a, b int) int {
    if a>b{
        return b
    }
    return a
}
// @lc code=end

