/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 *
 * https://leetcode-cn.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (59.59%)
 * Likes:    200
 * Dislikes: 0
 * Total Accepted:    29.7K
 * Total Submissions: 49.9K
 * Testcase Example:  '[1,2,2]'
 *
 * 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 * 
 * 说明：解集不能包含重复的子集。
 * 
 * 示例:
 * 
 * 输入: [1,2,2]
 * 输出:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 * 
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int,0)
	sort.Ints(nums)
	backtrack(&ans,nums,[]int{},0)
	return ans
}

func backtrack(ans *[][]int, nums, cur []int, index int) {
	temp := make([]int,len(cur))
	copy(temp,cur)
	*ans = append(*ans, temp)

	for i:=index; i<len(nums);i++{
		if i > index && nums[i] == nums[i-1] {
			continue
		}

		temp := make([]int,len(cur))
		copy(temp,cur)
		temp = append(temp,nums[i])
		backtrack(ans,nums,temp,i+1)
	}
}

// @lc code=end

