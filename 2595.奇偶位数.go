/*
 * @lc app=leetcode.cn id=2595 lang=golang
 *
 * [2595] 奇偶位数
 */

// @lc code=start
func evenOddBit(n int) []int {
    ans:=[]int{0,0}
    even:=true
    for n>0{
        if n&1==1{
            if even{
                ans[0]+=1
            }else{
                ans[1]+=1
            }
            
        }
        n>>=1
        even=!even
    }
    return ans
}
// @lc code=end

