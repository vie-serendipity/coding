/*
 * @lc app=leetcode.cn id=213 lang=java
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
class Solution {
    public int rob(int[] nums) {
        if(nums.length<2){
            return nums[0];
        }
        int n=nums.length;
        int[][] dp=new int[2][n];
        dp[0][0]=nums[0];//偷了0号房子
        dp[1][0]=0;
        dp[0][1]=nums[0];
        dp[1][1]=nums[1];
        for(int i=2;i<n;++i){
            dp[0][i]=Math.max(dp[0][i-2]+nums[i],dp[0][i-1]);
            dp[1][i]=Math.max(dp[1][i-2]+nums[i],dp[1][i-1]);
        }
        return Math.max(dp[1][n-1],dp[0][n-2]);
    }
}
// @lc code=end

