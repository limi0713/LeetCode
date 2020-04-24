/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (40.90%)
 * Likes:    349
 * Dislikes: 0
 * Total Accepted:    69.6K
 * Total Submissions: 169.8K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 * 
 * 示例 1:
 * 
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 * 
 * 
 * 示例 2:
 * 
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 * 
 */

// @lc code=start
func merge(intervals [][]int) [][]int {

	if len(intervals) <= 0 {
		return [][]int{}
	}

	array := make(InterValArray, 0)

	handle(&array,intervals)

	sort.Sort(array)

	ans := make([][]int, 0)

	ans = append(ans, []int{array[0].start,array[0].end})

	for i := 1; i < len(array); i ++ {
		if array[i].start <= ans[len(ans)-1][1] {
			ans[len(ans)-1][1] = max(array[i].end,ans[len(ans)-1][1])
		}else {
			ans = append(ans, []int{array[i].start, array[i].end})
		}
	}

	return ans
}

func handle(array *InterValArray, intervals [][]int) {
	for i := range intervals {
		interval := InterVal {
			start: intervals[i][0],
			end : intervals[i][1],
		}

		*array = append(*array,interval)
	}
}

type InterVal struct {
	start int
	end int
}

type InterValArray []InterVal

func (a InterValArray) Len()int{
	return len(a)
}

func (a InterValArray)Less(i, j int) bool {
	return a[i].start < a[j].start
}

func (a InterValArray)Swap(i, j int){
	a[i], a[j] = a[j], a[i]
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

// @lc code=end

