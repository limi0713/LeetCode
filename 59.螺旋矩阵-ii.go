/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (76.92%)
 * Likes:    173
 * Dislikes: 0
 * Total Accepted:    30.9K
 * Total Submissions: 40.2K
 * Testcase Example:  '3'
 *
 * 给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
 * 
 * 示例:
 * 
 * 输入: 3
 * 输出:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 * 
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	ans := make([][]int,0)
	for i:=0;i<n;i++{
		a := make([]int,n)
		ans =append(ans,a)
	}
	
	ring := (n-1)/2 + 1

	num := 1

	for i :=0 ; i < ring; i ++ {
		
		for j := i; j <= n - i - 1; j ++ {
			ans[i][j] = num
			num ++
		}

		for j := i+1; j <= n - i - 1;j ++ {
			ans[j][n-1-i] = num
			num ++
		}

		if n -1 -i > i {
			for j := n - 2 - i ; j >= i; j -- {
				ans[n-1-i][j] = num
				num ++
			}
		}

		if i < n - 1 - i {
			for j := n - 2 - i ; j >= i+1; j -- {
				ans[j][i] = num
				num ++
			}
		}
	}

	return ans
}
// @lc code=end

