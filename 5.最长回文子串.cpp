/*
 * @lc app=leetcode.cn id=5 lang=cpp
 *
 * [5] 最长回文子串
 */

// @lc code=start
#include<vector>
#include<iostream>
#include<string>
class Solution {
public:
    string longestPalindrome(string s) {
        int len=s.size();
        if(len<2){
            return s;
        }
        int maxlen=0;
        string res;
        int dp[len][len];
        for (int i = 0; i < len;++i){
            memset(dp[i],1,len);
        }
        for(int left=len-1;left>=0;--left){
            for(int right=left;right<len;++right){
                if(s[left]==s[right]){
                    if(right-left>1){
                        dp[left][right]=dp[left+1][right-1];
                    }
                    else{
                        dp[left][right]=true;
                    }
                }
                else{
                    dp[left][right]=0;
                }
                if((right-left+1>maxlen)&&dp[left][right]){
                    maxlen=right-left+1;
                    res=s.substr(left,maxlen);
                }
            }
        }
        return res;
    }
};
// @lc code=end

