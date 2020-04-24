/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (39.31%)
 * Likes:    371
 * Dislikes: 0
 * Total Accepted:    77.8K
 * Total Submissions: 198.1K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 * 
 * 你的算法时间复杂度必须是 O(log n) 级别。
 * 
 * 如果数组中不存在目标值，返回 [-1, -1]。
 * 
 * 示例 1:
 * 
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: [3,4]
 * 
 * 示例 2:
 * 
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: [-1,-1]
 * 
 */

// @lc code=start
func searchRangeStartEnd(nums []int,target,start,end int)[]int{
	if end < start {
		return []int{-1,-1}
	}

	if end == start {
		if nums[start] == target {
			return []int{start,end}
		}
		return []int{-1,-1}
	}
	
	mid := (end+start) / 2 

	if nums[mid] < target {
		return searchRangeStartEnd(nums, target, mid + 1, end)
	}else if nums[mid] > target {
		return searchRangeStartEnd(nums, target, start, mid - 1)
	}else {
		leftAns := searchRangeStartEnd(nums, target, start, mid)
		rightAns := searchRangeStartEnd(nums, target, mid+1, end)

		ans := make([]int,0)
		for i:=0;i<len(leftAns);i ++{
			if leftAns[i] != -1 {
				ans = append(ans,leftAns[i])
			}
		}

		for i:=0;i<len(rightAns);i ++{
			if rightAns[i] != -1 {
				ans = append(ans,rightAns[i])
			}
		}

		if len(ans) == 0 {
			return []int{-1,-1}
		}

		return []int{ans[0],ans[len(ans)-1]}
	}

}

func searchRange(nums []int, target int) []int {

	return searchRangeStartEnd(nums,target,0,len(nums)-1)

}
// @lc code=end

