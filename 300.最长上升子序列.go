/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 *
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (44.08%)
 * Likes:    601
 * Dislikes: 0
 * Total Accepted:    82.2K
 * Total Submissions: 186.2K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 * 
 * 示例:
 * 
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4 
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 * 
 * 说明:
 * 
 * 
 * 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
 * 你算法的时间复杂度应该为 O(n^2) 。
 * 
 * 
 * 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
 * 
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	ans := make([]int,len(nums))
	ans[0] = 1

	for i := 1; i < len(nums); i ++ {
		leng := 0
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] && ans[j] > leng {
				leng = ans[j]
			}
		}

		ans[i] = leng + 1
	}

	return findMax(ans)
}

func findMax(nums []int) int {
	max := 0

	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
// @lc code=end

