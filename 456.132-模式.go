/*
 * @lc app=leetcode.cn id=456 lang=golang
 *
 * [456] 132模式
 *
 * https://leetcode-cn.com/problems/132-pattern/description/
 *
 * algorithms
 * Medium (26.32%)
 * Likes:    143
 * Dislikes: 0
 * Total Accepted:    7.3K
 * Total Submissions: 27.8K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个整数序列：a1, a2, ..., an，一个132模式的子序列 ai, aj, ak 被定义为：当 i < j < k 时，ai < ak <
 * aj。设计一个算法，当给定有 n 个数字的序列时，验证这个序列中是否含有132模式的子序列。
 * 
 * 注意：n 的值小于15000。
 * 
 * 示例1:
 * 
 * 
 * 输入: [1, 2, 3, 4]
 * 
 * 输出: False
 * 
 * 解释: 序列中不存在132模式的子序列。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: [3, 1, 4, 2]
 * 
 * 输出: True
 * 
 * 解释: 序列中有 1 个132模式的子序列： [1, 4, 2].
 * 
 * 
 * 示例 3:
 * 
 * 
 * 输入: [-1, 3, 2, 0]
 * 
 * 输出: True
 * 
 * 解释: 序列中有 3 个132模式的的子序列: [-1, 3, 2], [-1, 3, 0] 和 [-1, 2, 0].
 * 
 * 
 */

// @lc code=start
func handleNums(nums, flag []int){
	if len(nums) == 0 || len(flag) == 0 {
		return
	}

	flag[0] = nums[0]

	min := 0
	for i := range nums {
		flag[i] = nums[min]

		if i > min && nums[i] < nums[min] {
			min = i
		}
	}

	return
}

func find(nums []int, start int, min int) int {

	for i := start; i < len(nums); i ++ {
		if nums[i] > min && nums[i] < nums[start - 1]{
			return i
		}
	}

	return -1
}

func find132pattern(nums []int) bool {

	flag := make([]int,len(nums))

	handleNums(nums, flag)

	for i := 1; i < len(nums) - 1; i ++ {
		if flag[i] >= nums[i] {
			continue
		}

		if index := find(nums,i+1,flag[i]); index >= 0 {
			return true
		}
	}

	return false
}
// @lc code=end

