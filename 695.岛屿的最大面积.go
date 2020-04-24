/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 *
 * https://leetcode-cn.com/problems/max-area-of-island/description/
 *
 * algorithms
 * Medium (58.73%)
 * Likes:    229
 * Dislikes: 0
 * Total Accepted:    33.5K
 * Total Submissions: 53.2K
 * Testcase Example:  '[[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]'
 *
 * 给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地)
 * 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
 * 
 * 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)
 * 
 * 示例 1:
 * 
 * 
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,1,1,0,1,0,0,0,0,0,0,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,0,1,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,1,1,0,0],
 * ⁠[0,0,0,0,0,0,0,0,0,0,1,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,0,0,0,0]]
 * 
 * 
 * 对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。
 * 
 * 示例 2:
 * 
 * 
 * [[0,0,0,0,0,0,0,0]]
 * 
 * 对于上面这个给定的矩阵, 返回 0。
 * 
 * 注意: 给定的矩阵grid 的长度和宽度都不超过 50。
 * 
 */

// @lc code=start

func dfs(grid [][]int,i,j int) int {
	grid[i][j] = 0

	res := 1
	if i-1>=0 && grid[i-1][j] == 1{
		res += dfs(grid,i-1,j)
	}

	if i+1 <len(grid) && grid[i+1][j] == 1{
		res += dfs(grid,i+1,j)
	}

	if j-1 >=0 && grid[i][j-1]==1{
		res+= dfs(grid,i,j-1)
	}

	if j+1 < len(grid[i]) && grid[i][j+1]==1{
		res += dfs(grid,i,j+1)
	}

	return res
}

func maxAreaOfIsland(grid [][]int) int {

	res := 0
	for i :=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			temp := 0
			if grid[i][j]==1{
				temp = dfs(grid,i,j)
			}

			if temp > res{
				res = temp
			}
		}
	}

	return res
}
// @lc code=end

