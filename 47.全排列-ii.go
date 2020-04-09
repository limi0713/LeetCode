/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (57.84%)
 * Likes:    258
 * Dislikes: 0
 * Total Accepted:    49.7K
 * Total Submissions: 85.8K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列，返回所有不重复的全排列。
 * 
 * 示例:
 * 
 * 输入: [1,1,2]
 * 输出:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 * 
 */

// @lc code=start
func dfs(ans *[][]int, nums, cur []int, hadAdd map[int]bool){
	if len(cur) == len(nums) {
		*ans = append(*ans, cur)
		return
	}

	for i := range nums {

		// nums[i-1]一定比nums[i]先搜索到，
		// 当搜索起点为i时，表示起点i-1已被搜索过
		// nums[i]==nums[i-1]则，nums[i]作为起点和nums[i-1]重复，剪枝
		if  _, ok := hadAdd[i-1]; i > 0 && nums[i] == nums[i-1] && !ok{
			continue
		}

		if _, ok := hadAdd[i]; ok {
			continue
		}

		temp := make([]int, 0)
		temp = append(temp, cur...)
		temp = append(temp, nums[i])

		hadAdd[i] = true

		dfs(ans, nums, temp, hadAdd)
		
		delete(hadAdd, i)
	}

}

func permuteUnique(nums []int) [][]int {

	sort.Ints(nums)

	ans := make([][]int, 0)

	dfs(&ans, nums, []int{}, make(map[int]bool))

	return ans
}
// @lc code=end

