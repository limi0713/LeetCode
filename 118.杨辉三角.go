/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	if numRows <1{
		return nil
	}
	var res [][]int
	for i:=0;i<numRows;i++{
		r := make([]int,i+1)
		res = append(res,r)
	}
	res[0][0]=1
	for i:=0;i<numRows;i++{
		res[i][0] = 1
		for j:=1;j<i;j++{
			res[i][j]=res[i-1][j-1]+res[i-1][j]
		}
		res[i][i]=1
	}
	return res
}
// @lc code=end

