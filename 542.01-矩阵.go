/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return [][]int{}
	}

	ans := make([][]int, 0)
	for i := 0; i < len(matrix); i ++ {
		a := make([]int,len(matrix[i]))
		ans = append(ans, a)
	}

	if matrix[0][0] != 0 {
		ans[0][0] = 100000
	}

	dp(&ans, matrix)

	return ans
}

func dp(ans *[][]int,matrix [][]int){

	for i := 0; i < len(matrix); i ++{
		for j := 0; j < len(matrix[i]); j ++ {
			if matrix[i][j] == 0 {
				(*ans)[i][j] = 0
			} else {
				a, b := 100000, 100000
				if i - 1 >= 0 {
					a = (*ans)[i-1][j] + 1
				}

				if j - 1 >= 0 {
					b = (*ans)[i][j-1] + 1
				}
				(*ans)[i][j] = min(a,b)

				// if  j - 1 < 0 && i - 1 >= 0 {
				// 	(*ans)[i][j] = (*ans)[i-1][j] + 1
				// }else if j - 1 >= 0 && i - 1 < 0{
				// 	(*ans)[i][j] = (*ans)[i][j-1] + 1
				// }else if i - 1 >= 0 && j - 1 >= 0 {
				// 	(*ans)[i][j] = min((*ans)[i][j-1],(*ans)[i-1][j]) + 1
				// }
			}
		}
	}

	for i := len(matrix) - 1; i >= 0; i -- {
		for j := len(matrix[i]) - 1; j >= 0; j -- {
			if matrix[i][j] == 0 {
				(*ans)[i][j] = 0
			} else {
				a, b := 100000,100000
				if i + 1 < len(matrix) {
					a = (*ans)[i+1][j]+1
				}

				if j + 1 < len(matrix[i]) {
					b = (*ans)[i][j+1]+1
				}

				(*ans)[i][j] = min((*ans)[i][j],min(a, b))


				// if  j + 1 >= len(matrix[i]) && i + 1 < len(matrix) {
				// 	(*ans)[i][j] = min((*ans)[i][j],(*ans)[i+1][j]+1)
				// }else if j + 1 < len(matrix[i]) && i + 1 >= len(matrix){
				// 	(*ans)[i][j] = min((*ans)[i][j],(*ans)[i][j+1]+1)
				// }else if j + 1 < len(matrix[i]) && i + 1 < len(matrix) {
				// 	(*ans)[i][j] = min((*ans)[i][j],min((*ans)[i][j+1],(*ans)[i+1][j]) + 1)
				// }
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
// @lc code=end

