/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
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

func maxInt(a,b int)int{
	if a>b {
		return a
	}
	return b
}

var res = 0

func caculate(root *TreeNode)int{
	if root == nil {
		return 0
	}
	left := caculate(root.Left)
	right := caculate(root.Right)
	l := 0
	r := 0
	if root.Left != nil && root.Left.Val == root.Val {
		l = left + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		r = right + 1
	}
	res = maxInt(res,l+r)
	return maxInt(l,r)
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res = 0
	caculate(root)
	return res
}

// @lc code=end

