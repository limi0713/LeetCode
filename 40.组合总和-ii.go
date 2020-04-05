/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (61.00%)
 * Likes:    231
 * Dislikes: 0
 * Total Accepted:    49.2K
 * Total Submissions: 80.6K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * 
 * candidates 中的每个数字在每个组合中只能使用一次。
 * 
 * 说明：
 * 
 * 
 * 所有数字（包括目标数）都是正整数。
 * 解集不能包含重复的组合。 
 * 
 * 
 * 示例 1:
 * 
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 所求解集为:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 * 
 * 
 * 示例 2:
 * 
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 所求解集为:
 * [
 * [1,2,2],
 * [5]
 * ]
 * 
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {

	if len(candidates) <= 0{
		return [][]int{}
	}

	sort.Ints(candidates)

	index := 0 
	for ; index < len(candidates); index ++ {
		if candidates[index] > target {
			break
		}
	}

	candidates = candidates[:index]

	if len(candidates) <= 0{
		return [][]int{}
	}

	ans := make([][]int, 0)

	solve(&ans, candidates, []int{}, 0, 0, target)
	
	return ans
}

func solve(ans *[][]int, candidates, cur []int, start, curSum, target int){

	if curSum == target {
		*ans = append(*ans,cur)
		return
	}

	if curSum > target {
		return
	}

	for i := start; i < len(candidates); {

		temp := make([]int,0)
		temp = append(temp,cur...)
		temp = append(temp,candidates[i])

		solve(ans, candidates, temp, i+1, curSum + candidates[i], target)

		i++
		for ; i < len(candidates); i ++ {
			if candidates[i-1] == candidates[i] {
				continue
			}
			break
		}
	}
}
// @lc code=end

