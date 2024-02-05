/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode{
    if left>right{
        return nil
    }
    mid:=(left+right)/2
    root:=&TreeNode{
        Val: nums[mid],
    }
    lt:=helper(nums, left, mid-1)
    rt:=helper(nums, mid+1, right)
    root.Left=lt
    root.Right=rt
    return root
}
// @lc code=end

