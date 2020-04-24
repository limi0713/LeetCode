/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (68.69%)
 * Likes:    589
 * Dislikes: 0
 * Total Accepted:    79.6K
 * Total Submissions: 115.9K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * 
 * candidates 中的数字可以无限制重复被选取。
 * 
 * 说明：
 * 
 * 
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。 
 * 
 * 
 * 示例 1:
 * 
 * 输入: candidates = [2,3,6,7], target = 7,
 * 所求解集为:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 * 
 * 
 * 示例 2:
 * 
 * 输入: candidates = [2,3,5], target = 8,
 * 所求解集为:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 * 
 */

// @lc code=start
var (
	ans = make([][]int,0)
)

func solve(candidates []int,target,curSum int, curAns []int){

	if target < candidates[0] {
		return
	}

	if curSum == target {
		ans = append(ans,curAns)
	}else if curSum > target{
		return
	}

	for i := range candidates {
		if len(curAns) > 0 && candidates[i] < curAns[len(curAns)-1] {
			continue
		}

		temp := make([]int,0)
		temp = append(temp,curAns...)
		temp = append(temp,candidates[i])
		curSum += candidates[i]
		solve(candidates,target,curSum,temp)
		curSum -= candidates[i]		
	}

}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) <= 0 {
		return [][]int{}
	}

	sort.Ints(candidates)

	if len(ans) > 0 {
		ans = append(ans[:0],ans[len(ans)-1:len(ans)-1]...)
	}

	solve(candidates,target,0,[]int{})

	return ans
} 
// @lc code=end

