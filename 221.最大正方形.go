/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 *
 * https://leetcode-cn.com/problems/maximal-square/description/
 *
 * algorithms
 * Medium (39.68%)
 * Likes:    310
 * Dislikes: 0
 * Total Accepted:    33.6K
 * Total Submissions: 84.7K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
 * 
 * 示例:
 * 
 * 输入: 
 * 
 * 1 0 1 0 0
 * 1 0 1 1 1
 * 1 1 1 1 1
 * 1 0 0 1 0
 * 
 * 输出: 4
 * 

[
["1","1","1","1","1","1","1","1"],
["1","1","1","1","1","1","1","0"],
["1","1","1","1","1","1","1","0"],
["1","1","1","1","1","0","0","0"],
["0","1","1","1","1","0","0","0"]]

 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	if len(matrix) <= 0 {
		return 0
	}

	maxAns := 0
	ans := make([][]int,0,2)
	for i := 0;i < 2;i++{
		a := make([]int,len(matrix[0]))
		ans = append(ans,a)
	}

	for i:= range matrix[0] {
		ans[0][i] = int(matrix[0][i]-'0')
		if ans[0][i] > maxAns {
			maxAns = ans[0][i]
		}
	}

	for i:=1; i < len(matrix); i ++ {
		for j := range matrix[i] {
			if matrix[i][j] == '0' {
				ans[i%2][j] = 0
				continue
			}

			if j == 0 {
				ans[i%2][j] = int(matrix[i][j]-'0')
			}else{
				ans[i%2][j] = min(ans[i%2][j-1],min(ans[(i-1)%2][j-1],ans[(i-1)%2][j]))+1
			}

			if ans[i%2][j] > maxAns {
				maxAns = ans[i%2][j]
			}
		}
	}

	return maxAns*maxAns
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

