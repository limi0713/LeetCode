/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (77.14%)
 * Likes:    534
 * Dislikes: 0
 * Total Accepted:    78.9K
 * Total Submissions: 102.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 * 
 * 说明：解集不能包含重复的子集。
 * 
 * 示例:
 * 
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 * 
 */

// @lc code=start
func subsets(nums []int) [][]int {
	if len(nums) <= 0 {
		return [][]int{[]int{}}
	}

	sort.Ints(nums)

	ans := make([][]int,0,int(math.Pow(2, float64(len(nums)))))

	cur := make([]int, 0, len(nums))

	backtrack(&ans, nums, cur, 0)

	return ans
}

func backtrack(ans *[][]int, nums,cur []int, index int){
	temp := make([]int,len(cur))
	copy(temp, cur)
	*ans = append(*ans, temp)

	for i := index; i < len(nums); i ++ {
		leng := len(cur)

		cur = append(cur, nums[i])

		backtrack(ans,nums,cur,i+1)

		cur = cur[:leng]
	}
}
// @lc code=end

