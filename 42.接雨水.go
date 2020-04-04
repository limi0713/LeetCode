/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (49.57%)
 * Likes:    1042
 * Dislikes: 0
 * Total Accepted:    75.2K
 * Total Submissions: 150.4K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 
 * 
 * 
 * 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢
 * Marcos 贡献此图。
 * 
 * 示例:
 * 
 * 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出: 6
 * 
 */

// @lc code=start

func findMaxIndex(height []int) int {

	if len(height) <= 0 {
		return -1
	}

	index := 1 

	for i:=2;i<len(height)-1;i++{
		if height[i] > height[index] {
			index = i 
		}
	}

	return index
}

func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}



func trap(height []int) int {

	if len(height) < 3{
		return 0
	}

	if len(height) == 3 {
		min := min(height[0],height[2])
		if min <=  height[1] {
			return 0
		}

		return min - height[1]
	}

	l := list.New()
	ans := 0

	for i := range height {
		if l.Len() == 0 {
			if height[i] > 0 {
				l.PushBack(i)
			}

			continue
		}

		for height[l.Back().Value.(int)] < height[i] {
			top := l.Back().Value.(int)
			l.Remove(l.Back())
			if l.Len() == 0 {
				break
			}

			preTop := l.Back().Value.(int)
			ans += (min(height[i],height[preTop]) - height[top]) * (i - preTop - 1)
		}

		l.PushBack(i)
	}

	return ans
	
}
// @lc code=end

