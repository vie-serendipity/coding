/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
    if len(s)<len(p){
        return []int{}
    }
    ans:=[]int{}
    count:=[26]int{}
    for i, ch:=range p {
        count[ch-'a']--
        count[s[i]-'a']++
    }
    differ:=0
    for _, c:=range count{
        if c!=0{
            differ++
        }
    }
    if differ==0{
        ans=append(ans, 0)
    }
    for i, ch := range s[:len(s)-len(p)]{
        if count[ch-'a']==0{
            differ++
        }else if count[ch-'a']==1{
            differ--
        }
        count[ch-'a']--

        if count[s[i+len(p)]-'a']==0{
            differ++
        }else if count[s[i+len(p)]-'a']==-1{
            differ--
        }
        count[s[i+len(p)]-'a']++
        if differ==0{
            ans=append(ans, i+1)
        }
    }
    return ans
}
// @lc code=end

