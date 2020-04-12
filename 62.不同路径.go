/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 *
 * https://leetcode-cn.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (59.74%)
 * Likes:    491
 * Dislikes: 0
 * Total Accepted:    91.8K
 * Total Submissions: 153.3K
 * Testcase Example:  '3\n2'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 * 
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 * 
 * 问总共有多少条不同的路径？
 * 
 * 
 * 
 * 例如，上图是一个7 x 3 的网格。有多少可能的路径？
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: m = 3, n = 2
 * 输出: 3
 * 解释:
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向右 -> 向下
 * 2. 向右 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向右
 * 
 * 
 * 示例 2:
 * 
 * 输入: m = 7, n = 3
 * 输出: 28
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= m, n <= 100
 * 题目数据保证答案小于等于 2 * 10 ^ 9
 * 
 * 
 */

// @lc code=start
func uniquePaths(m int, n int) int {

	ans := make([][]int,0)
	for i := 0;i<2;i++{
		a := make([]int,n)
		ans = append(ans, a)
	}

	for i := 0;i<m;i++{
		for j:=0;j<n;j++{
			if i == 0 || j == 0{
				ans[i%2][j] = 1
				continue
			}

			ans[i%2][j] = ans[i%2][j-1] + ans[(i-1)%2][j]
		}
	}

	return ans[(m-1)%2][n-1]
}
// @lc code=end

