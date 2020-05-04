/*
 * @lc app=leetcode.cn id=1424 lang=golang
 *
 * [1424] 对角线遍历 II
 *
 * https://leetcode-cn.com/problems/diagonal-traverse-ii/description/
 *
 * algorithms
 * Medium (33.39%)
 * Likes:    8
 * Dislikes: 0
 * Total Accepted:    2.2K
 * Total Submissions: 6.6K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个列表 nums ，里面每一个元素都是一个整数列表。请你依照下面各图的规则，按顺序返回 nums 中对角线上的整数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：nums = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,4,2,7,5,3,8,6,9]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：nums = [[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]
 * 输出：[1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]
 * 
 * 
 * 示例 3：
 * 
 * 输入：nums = [[1,2,3],[4],[5,6,7],[8],[9,10,11]]
 * 输出：[1,4,2,5,3,8,6,9,7,10,11]
 * 
 * 
 * 示例 4：
 * 
 * 输入：nums = [[1,2,3,4,5,6]]
 * 输出：[1,2,3,4,5,6]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i].length <= 10^5
 * 1 <= nums[i][j] <= 10^9
 * nums 中最多有 10^5 个数字。
 * 
 * 
 */

// @lc code=start
func findDiagonalOrder(nums [][]int) []int {
	ansMap := make(map[int][]int)
	size := 0

	for i := range nums {
		for j := range nums[i] {
			size ++
			if _ ,ok := ansMap[i+j]; !ok {
				ansMap[i+j] = []int{nums[i][j]}
				continue
			}

			ansMap[i+j] = append(ansMap[i+j],nums[i][j])
		}
	}

	ansArr := make([][]int,0,len(ansMap))
	for i := 0; i < len(ansMap); i ++ {
		a := make([]int,0,len(ansMap[i]))
		ansArr = append(ansArr,a)
	}

	for k,v := range ansMap {
		for i := len(v)-1;i >= 0; i -- {
			ansArr[k] = append(ansArr[k],v[i])
		}
	}

	ans := make([]int,0,size)
	for i := range ansArr {
		ans = append(ans,ansArr[i]...)
	}
	
	return ans
}
// @lc code=end

