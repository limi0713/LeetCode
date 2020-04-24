/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 *
 * https://leetcode-cn.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (38.55%)
 * Likes:    547
 * Dislikes: 0
 * Total Accepted:    78.7K
 * Total Submissions: 203.9K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 * 
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 
 * 判断你是否能够到达最后一个位置。
 * 
 * 示例 1:
 * 
 * 输入: [2,3,1,1,4]
 * 输出: true
 * 解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
 * 
 * 
 * 示例 2:
 * 
 * 输入: [3,2,1,0,4]
 * 输出: false
 * 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
 * 
 * 
 */

// @lc code=start
func canJump(nums []int) bool {
	if len(nums) <= 0 {
		return true
	}

	ans := make([]bool,len(nums))

	ans[0] = true

	for i:=0;i<len(ans);i++{
		if !ans[i] {
			continue
		}

		for j := 1; j<=nums[i];j ++ {
			if i+j < len(ans){
				ans[i+j] = true
			}
		}

		if ans[len(nums)-1] {
			return true
		}
	}

	return ans[len(nums) - 1]
}
// @lc code=end

