/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (48.38%)
 * Likes:    1433
 * Dislikes: 0
 * Total Accepted:    134.8K
 * Total Submissions: 278.4K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 * 
 * 示例:
 * 
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 * 
 * 
 * 进阶:
 * 
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 * 
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	ans := make([]int,len(nums))
	ans[0] =nums[0]
	maxIndex := 0
	for i := 1; i < len(nums); i ++ {
		if nums[i] > nums[i] + ans[i-1] {
			ans[i] = nums[i]
		} else {
			ans[i] = nums[i] + ans[i-1]
		}

		if ans[i] > ans[maxIndex] {
			maxIndex = i
		}
	}
	
	return ans[maxIndex]
}
// @lc code=end

