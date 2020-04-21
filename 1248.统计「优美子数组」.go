/*
 * @lc app=leetcode.cn id=1248 lang=golang
 *
 * [1248] 统计「优美子数组」
 *
 * https://leetcode-cn.com/problems/count-number-of-nice-subarrays/description/
 *
 * algorithms
 * Medium (47.86%)
 * Likes:    34
 * Dislikes: 0
 * Total Accepted:    4.8K
 * Total Submissions: 9.7K
 * Testcase Example:  '[1,1,2,1,1]\n3'
 *
 * 给你一个整数数组 nums 和一个整数 k。
 * 
 * 如果某个 连续 子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。
 * 
 * 请返回这个数组中「优美子数组」的数目。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [1,1,2,1,1], k = 3
 * 输出：2
 * 解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [2,4,6], k = 1
 * 输出：0
 * 解释：数列中不包含任何奇数，所以不存在优美子数组。
 * 
 * 
 * 示例 3：
 * 
 * 输入：nums = [2,2,2,1,2,2,1,2,2,2], k = 2
 * 输出：16
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 50000
 * 1 <= nums[i] <= 10^5
 * 1 <= k <= nums.length
 * 
 * 
 */

// @lc code=start
func numberOfSubarrays(nums []int, k int) int {
	oddNumIndex := make([]int, 0,len(nums))
	
	for i := range nums {
		if nums[i] % 2 == 1 {
			oddNumIndex = append(oddNumIndex,i)
		}
	}

	if len(oddNumIndex) < k {
		return 0
	}

	ans := 0
	for i:= k-1;i < len(oddNumIndex); i ++ {
		leftEven, rightEven := oddNumIndex[i-k+1], len(nums) - oddNumIndex[i] - 1
		if i - k >= 0 {
			leftEven = oddNumIndex[i-k+1] - oddNumIndex[i-k] - 1
		}

		if i + 1 < len(oddNumIndex) {
			rightEven = oddNumIndex[i+1] - oddNumIndex[i] - 1
		}

		ans += (leftEven+1) * (rightEven+1)
	}

	return ans
}
// @lc code=end

