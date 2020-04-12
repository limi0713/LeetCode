/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 *
 * https://leetcode-cn.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (32.63%)
 * Likes:    265
 * Dislikes: 0
 * Total Accepted:    52.8K
 * Total Submissions: 161.4K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 * 
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 * 
 * 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
 * 
 * 
 * 
 * 网格中的障碍物和空位置分别用 1 和 0 来表示。
 * 
 * 说明：m 和 n 的值均不超过 100。
 * 
 * 示例 1:
 * 
 * 输入:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * 输出: 2
 * 解释:
 * 3x3 网格的正中间有一个障碍物。
 * 从左上角到右下角一共有 2 条不同的路径：
 * 1. 向右 -> 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右 -> 向右
 * 
 * 
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m <= 0 {
		return 0
	}

	n := len(obstacleGrid[0])

	ans := make([][]int,0)
	for i := 0;i<2;i++{
		a := make([]int,n)
		ans = append(ans, a)
	}

	for i := 0;i<m;i++{
		for j:=0;j<n;j++{
			if obstacleGrid[i][j] == 1 {
				ans[i%2][j] = 0
				continue
			}

			if i == 0 {
				if j - 1 >= 0 {
					ans[i%2][j] = ans[i%2][j-1]
				}else{
					ans[i%2][j] = 1
				}
				continue
			}

			if j == 0 {
				if i-1 >=0{
					ans[i%2][j] = ans[(i-1)%2][j]
				}else{
					ans[i%2][j] = 1
				}
				continue
			}

			ans[i%2][j] = ans[i%2][j-1] + ans[(i-1)%2][j]
		}
	}

	return ans[(m-1)%2][n-1]
}
// @lc code=end

