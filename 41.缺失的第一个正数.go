/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 *
 * https://leetcode-cn.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (37.86%)
 * Likes:    456
 * Dislikes: 0
 * Total Accepted:    44.5K
 * Total Submissions: 117.4K
 * Testcase Example:  '[1,2,0]'
 *
 * 给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: [1,2,0]
 * 输出: 3
 * 
 * 
 * 示例 2:
 * 
 * 输入: [3,4,-1,1]
 * 输出: 2
 * 
 * 
 * 示例 3:
 * 
 * 输入: [7,8,9,11,12]
 * 输出: 1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。
 * 
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	leng := len(nums)
	if leng <= 0 {
		return 1
	}

	one := false 
	for i := 0; i < leng; i ++ {
		if nums[i] == 1 {
			one = true
		} else if nums[i] <= 0 {
			nums[i] = 1	
		}
	}

	if !one {
		return 1
	}

	for i := 0; i < leng; i ++ {
		index := nums[i]
		if index < 0 {
			index = - index
		}

		if index > 0 && index <= leng {
			if nums[index - 1] > 0 {
				nums[index - 1] = - nums[index - 1]
			}
		}
	}

	ans := 0
	for ; ans < leng; ans ++ {
		if nums[ans] > 0 {
			break
		}
	}

	return ans + 1
}
// @lc code=end

