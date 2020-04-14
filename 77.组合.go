/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (73.44%)
 * Likes:    245
 * Dislikes: 0
 * Total Accepted:    44.6K
 * Total Submissions: 60.7K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 * 
 * 示例:
 * 
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 * 
 */

// @lc code=start
func combine(n int, k int) [][]int {
	if  n < k {
		return [][]int{}
	}

	ans := make([][]int,0)

	dfs(&ans,n,k,[]int{})

	return ans
}

func dfs(ans *[][]int, n,k int, cur []int) {
	if len(cur) == k {
		*ans = append(*ans, cur)
		return
	}

	start := 1
	if len(cur) > 0 {
		start = cur[len(cur)-1] + 1
	}
	for ;start<=n;start ++ {
		temp := make([]int,0)
		temp = append(temp, cur...)
		temp = append(temp, start)

		dfs(ans,n,k,temp)
	}
}
// @lc code=end

