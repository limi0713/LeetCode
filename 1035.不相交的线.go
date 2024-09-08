/*
 * @lc app=leetcode.cn id=1035 lang=golang
 *
 * [1035] 不相交的线
 */

// @lc code=start
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)
	dp := make([][]int, l1)
	for i := 0; i < l1; i++ {
		dp[i] = make([]int, l2)
	}

	for i := 0; i < l1; i++ {
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
		} else if i > 0 {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j := 0; j < l2; j++ {
		if nums1[0] == nums2[j] {
			dp[0][j] = 1
		} else if j > 0 {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[l1-1][l2-1]
}

// @lc code=end

