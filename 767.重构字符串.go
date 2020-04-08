/*
 * @lc app=leetcode.cn id=767 lang=golang
 *
 * [767] 重构字符串
 *
 * https://leetcode-cn.com/problems/reorganize-string/description/
 *
 * algorithms
 * Medium (39.69%)
 * Likes:    63
 * Dislikes: 0
 * Total Accepted:    5.5K
 * Total Submissions: 13.8K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。
 * 
 * 若可行，输出任意可行的结果。若不可行，返回空字符串。
 * 
 * 示例 1:
 * 
 * 
 * 输入: S = "aab"
 * 输出: "aba"
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: S = "aaab"
 * 输出: ""
 * 
 * 
 * 注意:
 * 
 * 
 * S 只包含小写字母并且长度在[1, 500]区间内。
 * 
 * 
 */

// @lc code=start
type ChNode struct {
	ch byte
	count int
}

type ChCount []*ChNode

func (c ChCount)Len() int {
	return len(c)
}

func (c ChCount)Less(i, j int) bool {
	return c[i].count > c[j].count
}

func (c ChCount)Swap(i,j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *ChCount) Push(item interface {}){
	*c = append(*c, item.(*ChNode))
}

func (c *ChCount) Pop() interface {} {
	item := (*c)[len(*c)-1]
	(*c) = (*c)[:len(*c)-1]
	return item
}

func (c *ChCount)Update(item *ChNode) {
	item.count -- 
}

func reorganizeString(S string) string {
	if len(S) <= 0 {
		return ""
	}
	
	chMap := make(map[byte]int)

	arr := make([]byte,len(S))

	for i := range S {
		chMap[S[i]] = chMap[S[i]] + 1
		arr[i] = '.'
	}

	chNodes := make(ChCount,0)
	for k, v := range chMap {
		item := &ChNode {
			ch : k,
			count : v,
		}
		chNodes = append(chNodes,item)
		delete(chMap,k)
	}

	heap.Init(&chNodes)

	ans := ""
	for len(chNodes) > 0 {
		item1 := heap.Pop(&chNodes).(*ChNode)
		if len(ans) <= 0 || ans[len(ans)-1] != item1.ch {
			ans += string(item1.ch)
			item1.count -- 
			if item1.count > 0{
				heap.Push(&chNodes,item1)
			}
			continue
		}

		if len(chNodes) <= 0 {
			break
		}

		item2 := heap.Pop(&chNodes).(*ChNode)
		ans += string(item2.ch)
		item2.count --
		if item2.count > 0 {
			heap.Push(&chNodes,item2)
		}
		heap.Push(&chNodes,item1)
	}

	if len(ans) == len(S) {
		return ans
	}

	return ""

}
// @lc code=end

