/*
 * @lc app=leetcode.cn id=5 lang=java
 *
 * [5] 最长回文子串
 */

// @lc code=start
class Solution {
    public String longestPalindrome(String s) {
        int len;
        String res=String.valueOf(s.charAt(0));
        int max_len=1;
        for(int i=0;i<s.length();++i){
            len=1;
            for(int j=1;j<Math.min(i,s.length()-i-1)+1;++j){
                if(s.charAt(i+j)==s.charAt(i-j)){
                    len+=2;
                    if(len>max_len){
                        max_len=len;
                        res=s.substring(i-j,i+j+1);
                    }
                }
                else{
                    break;
                }
            }       
            len=0;
            for(int j=0;j<Math.min(i,s.length()-i-2)+1;++j){
                if(s.charAt(i+j+1)==s.charAt(i-j)){
                    len+=2;
                    if(len>max_len){
                        max_len=len;
                        res=s.substring(i-j,i+j+2);
                    }
                }
                else{
                    break;
                }
            }
        }
        return res;
    }
}
// @lc code=end

