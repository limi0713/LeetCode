/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 *
 * https://leetcode-cn.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (59.18%)
 * Likes:    214
 * Dislikes: 0
 * Total Accepted:    44.9K
 * Total Submissions: 75.8K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 给定如下二叉树，以及目标和 sum = 22，
 * 
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \    / \
 * ⁠       7    2  5   1
 * 
 * 
 * 返回:
 * 
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
 * ]
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
func pathSum(root *TreeNode, sum int) [][]int {
	ans := make([][]int,0)
	DFS(root,0,sum,[]int{},&ans)
	return ans
}

func DFS(node *TreeNode,curSum,sum int, cur []int,ans *[][]int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		if curSum + node.Val == sum {
			temp := make([]int,len(cur))
			copy(temp,cur)
			temp = append(temp,node.Val)
			*ans = append(*ans,temp)
		}
		return
	}

	// temp := make([]int,len(cur)+1)
	// copy(temp,cur)

	DFS(node.Left,curSum+node.Val,sum,append(cur,node.Val),ans)

	DFS(node.Right,curSum+node.Val,sum,append(cur,node.Val),ans)
}
// @lc code=end

