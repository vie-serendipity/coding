/*
 * @lc app=leetcode.cn id=198 lang=java
 *
 * [198] 打家劫舍
 */

// @lc code=start
class Solution {
    public int rob(int[] nums) {
        int n=nums.length;
        int[][] dp=new int[2][n];
        dp[0][0]=0;
        dp[1][0]=nums[0];
        int res=nums[0];
        for(int i=1;i<n;++i){
            dp[0][i]=Math.max(dp[0][i-1],dp[1][i-1]);
            dp[1][i]=dp[0][i-1]+nums[i];
            res=Math.max(res,dp[0][i]);
            res=Math.max(res,dp[1][i]);
        }
        return res;
    }
}
// @lc code=end

