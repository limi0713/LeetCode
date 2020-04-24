/*
 * @lc app=leetcode.cn id=993 lang=golang
 *
 * [993] 二叉树的堂兄弟节点
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

type Node struct{
	Deep int
	Par int
}

func DF(root *TreeNode,record map[int]Node,deep int,par int){
	if root ==nil{
		return
	}
	DF(root.Left,record,deep+1,root.Val)
	record[root.Val]=Node{Deep:deep,Par:par}
	DF(root.Right,record,deep+1,root.Val)
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	record := make(map[int]Node)
	DF(root,record,0,0)
	nodeX := record[x]
	nodeY := record[y]
	if nodeX.Deep == nodeY.Deep && nodeX.Par != nodeY.Par {
		return true
	}
	return false
}
// @lc code=end

