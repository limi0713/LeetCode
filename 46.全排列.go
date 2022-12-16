/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (74.58%)
 * Likes:    591
 * Dislikes: 0
 * Total Accepted:    98.7K
 * Total Submissions: 132.3K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
 * 
 * 示例:
 * 
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 * 
 */

// @lc code=start
func permute(nums []int) [][]int {
	if len(nums) <= 0 {
		return [][]int{}
	}

	ans := make([][]int, 0)

	dfs(&ans,nums,[]int{},map[int]bool{})

	return ans
}

func dfs(ans *[][]int, nums, cur []int, hadAdd map[int]bool){
	if len(cur) == len(nums) {
		*ans = append(*ans,cur)
		return
	}

	for i := 0; i < len(nums); i ++ {
		if _, ok := hadAdd[i]; ok {
			continue
		}

		temp := make([]int,0)
		temp = append(temp,cur...)
		temp = append(temp,nums[i])

		hadAdd[i] = true

		dfs(ans,nums,temp,hadAdd)

		delete(hadAdd,i)
	}
}

// @lc code=end

func permute(nums []int) [][]int {
    ans := make([][]int,0)
    
    flag := make([]bool,len(nums))
    
    backtrack(&ans,nums,[]int{},&flag)
    
    return ans
}

func backtrack(ans *[][]int,nums,cur []int,flag *[]bool){
    if len(cur) == len(nums) {
        temp :=cur[:]
        *ans = append(*ans, temp)
        return
    }
    
    for i := 0;i<len(nums);i++{
        if (*flag)[i] {
            continue
        }
        
        (*flag)[i] = true
        
        temp :=cur[:]
        temp = append(temp,nums[i])
        backtrack(ans,nums,temp,flag)
        
        (*flag)[i] = false
    }
}