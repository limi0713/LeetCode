/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 *
 * https://leetcode-cn.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (64.15%)
 * Likes:    195
 * Dislikes: 0
 * Total Accepted:    33.2K
 * Total Submissions: 51.8K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 * 
 * 示例:
 * 
 * 输入: [1,2,3,null,5,null,4]
 * 输出: [1, 3, 4]
 * 解释:
 * 
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)
	nodeLevMap := make(map[int]([]*TreeNode))

	nodeLevMap[1] = []*TreeNode{root}

	bfs(&ans,&nodeLevMap,1)
	return ans
}

func bfs(ans *[]int, nodeLevMap *map[int]([]*TreeNode), curLev int){

	*ans = append(*ans, (*nodeLevMap)[curLev][len((*nodeLevMap)[curLev])-1].Val)
	
	(*nodeLevMap)[curLev + 1] = make([]*TreeNode,0)
	
	for _, v := range (*nodeLevMap)[curLev] {
		if v.Left != nil {
			(*nodeLevMap)[curLev+1] = append((*nodeLevMap)[curLev+1], v.Left)
		}

		if v.Right != nil {
			(*nodeLevMap)[curLev+1] = append((*nodeLevMap)[curLev+1], v.Right)
		}
	}

	delete(*nodeLevMap, curLev)

	if len((*nodeLevMap)[curLev+1]) > 0 {
		bfs(ans, nodeLevMap, curLev+1)
	}
}
// @lc code=end

