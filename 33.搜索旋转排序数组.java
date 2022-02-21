/*
 * @lc app=leetcode.cn id=33 lang=java
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
class Solution {
    public int search(int[] nums, int target) {
        return binarySearch(nums,0,nums.length-1,target);
    }
    private int binarySearch(int[] nums,int left,int right,int target){
        if(left>right){
            return -1;
        }
        int pos=(left+right)/2;
        if(nums[pos]==target){
            return pos;
        }
        if(nums[pos]<nums[right]){
            if(target<=nums[right]&&target>nums[pos]){
                return binarySearch(nums,pos+1,right,target);
            }
            else{
                return binarySearch(nums,left,pos-1,target);
            }
        }
        else{
            if(target<nums[pos]&&target>=nums[left]){
                return binarySearch(nums,left,pos-1,target);
            }
            else{
                return binarySearch(nums,pos+1,right,target);
            }
        }
    }
}
// @lc code=end

