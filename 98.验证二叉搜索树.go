/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (29.97%)
 * Likes:    537
 * Dislikes: 0
 * Total Accepted:    101.6K
 * Total Submissions: 335.5K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 * 
 * 假设一个二叉搜索树具有如下特征：
 * 
 * 
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
 * 
 * 
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
 func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    vals := make([]int,0)
    
	return MFS(root,&vals)
}

// 中序遍历
func MFS(root *TreeNode, vals *[]int) bool {
    if root == nil {
        return true
    }
	
	res := MFS(root.Left,vals)
	if !res {
		return false
	}

	if len(*vals)-1 >=0 && root.Val <= (*vals)[len(*vals)-1] {
		return false
	}
    *vals = append(*vals,root.Val)
    
	return MFS(root.Right,vals)
}
// @lc code=end

