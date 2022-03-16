/*
 * @lc app=leetcode.cn id=34 lang=java
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
class Solution {
    public int[] searchRange(int[] nums, int target) {
        int[] res=new int[2];
        int pos=binarySearch(nums,target,0,nums.length-1);
        res[0]=pos;
        res[1]=pos;
        if(pos==-1){
            return res;
        }
        for(int i=pos-1;i>=0;--i){
            if(nums[i]!=target){
                res[0]=i+1;
                break;
            }
            res[0]=i;
        }
        for(int i=pos+1;i<nums.length;++i){
            if(nums[i]!=target){
                res[1]=i-1;
                break;
            }
            res[1]=i;
        }
        return res;
    }
    private int binarySearch(int[] nums,int target,int left,int right){
        if(nums.length==0){
            return -1;
        }
        if(left==right){
            return nums[left]==target?left:-1;
        }
        int mid=(left+right)/2;
        if(target>nums[mid]){
            return binarySearch(nums,target,mid+1,right);
        }
        else if(target<nums[mid]){
            return binarySearch(nums,target,left,mid);
        }
        else{
            return mid;
        }
    }
}
// @lc code=end

