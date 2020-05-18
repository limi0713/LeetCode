/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (38.03%)
 * Likes:    566
 * Dislikes: 0
 * Total Accepted:    65.6K
 * Total Submissions: 167.2K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 * 
 * 
 * 示例 2:
 * 
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 * 
 */

// @lc code=start
func maxProduct(nums []int) int {

	if len(nums) <= 0 {
		return 0
	}

	ans := nums[0]
	// max := make([]int,len(nums))
	// min := make([]int,len(nums))

	max,min := nums[0],nums[0]
	for i:=1;i<len(nums);i++{
		max,min = maxF(max * nums[i],maxF(min * nums[i],nums[i])),minF(max * nums[i],minF(min * nums[i],nums[i]))

		if max > ans {
			ans = max
		}
	}

	return ans
}

func maxF(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func minF(a,b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

