/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 *
 * https://leetcode-cn.com/problems/insert-interval/description/
 *
 * algorithms
 * Hard (37.10%)
 * Likes:    128
 * Dislikes: 0
 * Total Accepted:    19.4K
 * Total Submissions: 52.3K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * 给出一个无重叠的 ，按照区间起始端点排序的区间列表。
 * 
 * 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
 * 
 * 示例 1:
 * 
 * 输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
 * 输出: [[1,5],[6,9]]
 * 
 * 
 * 示例 2:
 * 
 * 输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * 输出: [[1,2],[3,10],[12,16]]
 * 解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
 * 
 * 
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) <= 0 {
		return [][]int{newInterval}
	}

	if len(newInterval) <= 0 {
		return intervals
	}

	ls := list.New()
	hadInsert := false
	for i := 0; i < len(intervals); i ++ {
		
		if intervals[i][0] > newInterval[0] && !hadInsert {
			node := Node{
				start:newInterval[0],
				end:newInterval[1],
			}

			ls.PushBack(node)
			hadInsert = true
			i --
			continue
		}
		node := Node{
			start:intervals[i][0],
			end:intervals[i][1],
		}

		ls.PushBack(node)
	}

	if !hadInsert {
		node := Node{
			start:newInterval[0],
			end:newInterval[1],
		}

		ls.PushBack(node)
		hadInsert = true
	}

	ans := make([][]int, 0)
	ans = append(ans,[]int{ls.Front().Value.(Node).start,ls.Front().Value.(Node).end})
	ls.Remove(ls.Front())

	for ls.Len() > 0 {
		node := ls.Front().Value.(Node)

		if node.start <= ans[len(ans)-1][1] {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], node.end)
		} else {
			ans = append(ans,[]int{node.start, node.end})
		}
		ls.Remove(ls.Front())
	}

	return ans
}

type Node struct {
	start int
	end int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
// @lc code=end

