/*
 * @lc app=leetcode.cn id=5394 lang=golang
 *
 * [5394] 对角线遍历 II
 *
 * https://leetcode-cn.com/problems/diagonal-traverse-ii/description/
 *
 * algorithms
 * Medium (22.95%)
 * Likes:    1
 * Dislikes: 0
 * Total Accepted:    884
 * Total Submissions: 3.9K
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
	ans := make([]int,0)

	maxCol := 0
	for i := range nums {
		if len(nums[i]) > maxCol{
			maxCol = len(nums[i])
		}

		for j := i;j>=0;j--{
			if i-j < len(nums[j]) {
				ans = append(ans,nums[j][i-j])
			}
		}
	}

	for j := 1;j < maxCol;j++{
		for i := len(nums)-1;i>=0;i--{
			if len(nums)-1+j -i < len(nums[i]){
				ans = append(ans,nums[i][len(nums)-1+j -i])
			}
		}
	}

	return ans
}
// @lc code=end

