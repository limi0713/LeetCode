/*
 * @lc app=leetcode.cn id=5403 lang=golang
 *
 * [5403] 有序矩阵中的第 k 个最小数组和
 *
 * https://leetcode-cn.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/description/
 *
 * algorithms
 * Hard (43.24%)
 * Likes:    6
 * Dislikes: 0
 * Total Accepted:    938
 * Total Submissions: 2K
 * Testcase Example:  '[[1,3,11],[2,4,6]]\n5'
 *
 * 给你一个 m * n 的矩阵 mat，以及一个整数 k ，矩阵中的每一行都以非递减的顺序排列。
 *
 * 你可以从每一行中选出 1 个元素形成一个数组。返回所有可能数组中的第 k 个 最小 数组和。
 *
 *
 *
 * 示例 1：
 *
 * 输入：mat = [[1,3,11],[2,4,6]], k = 5
 * 输出：7
 * 解释：从每一行中选出一个元素，前 k 个和最小的数组分别是：
 * [1,2], [1,4], [3,2], [3,4], [1,6]。其中第 5 个的和是 7 。
 *
 * 示例 2：
 *
 * 输入：mat = [[1,3,11],[2,4,6]], k = 9
 * 输出：17
 *
 *
 * 示例 3：
 *
 * 输入：mat = [[1,10,10],[1,4,5],[2,3,6]], k = 7
 * 输出：9
 * 解释：从每一行中选出一个元素，前 k 个和最小的数组分别是：
 * [1,1,2], [1,1,3], [1,4,2], [1,4,3], [1,1,6], [1,5,2], [1,5,3]。其中第 7 个的和是 9
 * 。
 *
 *
 * 示例 4：
 *
 * 输入：mat = [[1,1,10],[2,2,9]], k = 7
 * 输出：12
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == mat.length
 * n == mat.length[i]
 * 1 <= m, n <= 40
 * 1 <= k <= min(200, n ^ m)
 * 1 <= mat[i][j] <= 5000
 * mat[i] 是一个非递减数组
 *
 *
 */

package main

// @lc code=start
func kthSmallest(mat [][]int, k int) int {
	left, right := 0, 0
	for i := range mat {
		left += mat[i][0] // 最小和

		right += mat[i][len(mat[i])-1] // 最大和
	}

	min := left

	for left < right {
		mid := left + (right-left)/2 // 统计比mid小的和有多少个

		leftCount := 1 // 比mid小的个数，left比mid小，所以初始为1

		dfs(mat, &leftCount, mid, min, k, 0)

		if leftCount >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func dfs(mat [][]int, count *int, mid, min, k, index int) {
	if index >= len(mat) || *count > k || min > mid {
		return
	}

	// 深度优先遍历，先确定最后一行的元素，0:n-2行都选择0号元素，
	// 从1（因为0号元素已经初始化）开始遍历最后一行
	dfs(mat, count, mid, min, k, index+1)

	/*
		当前和满足条件<= mid时，遍历index行,表示将选择的index行元素进行替换，
		替换之后，扔满足要求，则替换下面的行
	*/
	for i := 1; i < len(mat[index]); i++ {
		if min+mat[index][i]-mat[index][0] <= mid { // 在第index行选择第i个元素，和小于mid时，调整index行之后的行元素
			*count += 1
			dfs(mat, count, mid, min+mat[index][i]-mat[index][0], k, index+1)
		} else { // 当index行出现一个元素，和大于mid时，则index行之后的元素都不需要考虑
			break
		}
	}
}

// @lc code=end

// func main() {
// 	fmt.Println(kthSmallest([][]int{{1, 3, 11}, {2, 4, 6}}, 5))
// }
