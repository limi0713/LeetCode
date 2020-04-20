/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (39.28%)
 * Likes:    535
 * Dislikes: 0
 * Total Accepted:    40.6K
 * Total Submissions: 103.5K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 * 
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 * 
 * 
 * 
 * 
 * 
 * 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
 * 
 * 
 * 
 * 
 * 
 * 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入: [2,1,5,6,2,3]
 * 输出: 10
 * 
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	return handle(heights,0,len(heights)-1)
}

func handle(heights []int, start, end int) int {
	if start > end {
		return 0
	}

	if start == end {
		return heights[start]
	}

	minIndex := findMin(heights,start,end)
	area := heights[minIndex] * (end - start + 1)

	return max(max(area, handle(heights,start, minIndex - 1)),handle(heights,minIndex + 1, end))
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func findMin(heights []int, start, end int) int {

	minIndex := start
	for i := start + 1;i <= end; i ++ {
		if heights[i] < heights[minIndex] {
			minIndex = i
		}
	}

	return minIndex
}
// @lc code=end

