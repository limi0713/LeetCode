/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (65.21%)
 * Likes:    485
 * Dislikes: 0
 * Total Accepted:    39.8K
 * Total Submissions: 61.1K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
 * 
 * 示例:
 * 
 * 输入: 3
 * 输出: 5
 * 解释:
 * 给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
 * 
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 * 
 */

// @lc code=start
func numTrees(n int) int {
	if n <= 1{
		return 1
	}

	ans := make([]int,n+1)
	ans[0],ans[1],ans[2] =1,1,2

	for i:=3;i<=n;i++{
		for j:=0;j< i/2;j++{
			ans[i] += ans[j] * ans[i-1-j]
		}

		ans[i] *=2
		if i % 2 != 0 {
			ans[i] += ans[(i-1)/2] * ans[(i-1)/2]
		}
	}

	return ans[n]
}
// @lc code=end

