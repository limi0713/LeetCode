/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	var res [][]int
	for i:=0;i<2;i++{
		r := make([]int,rowIndex+1)
		res=append(res,r)
	}
	res[0][0] = 1
	for i:=0;i<=rowIndex;i++{
		res[i%2][0]=1
		for j:=1;j<i;j++{
			res[i%2][j]=res[(i-1)%2][j-1]+res[(i-1)%2][j]
		}
		res[i%2][i]=1
	}
	return res[rowIndex%2]
}
// @lc code=end

