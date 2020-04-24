/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 *
 * https://leetcode-cn.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Medium (48.53%)
 * Likes:    217
 * Dislikes: 0
 * Total Accepted:    29.4K
 * Total Submissions: 60.6K
 * Testcase Example:  '3\n3'
 *
 * 给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
 * 
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 * 
 * 
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 * 
 * 
 * 给定 n 和 k，返回第 k 个排列。
 * 
 * 说明：
 * 
 * 
 * 给定 n 的范围是 [1, 9]。
 * 给定 k 的范围是[1,  n!]。
 * 
 * 
 * 示例 1:
 * 
 * 输入: n = 3, k = 3
 * 输出: "213"
 * 
 * 
 * 示例 2:
 * 
 * 输入: n = 4, k = 9
 * 输出: "2314"
 * 
 * 
 */

// @lc code=start

type MyList struct {
	ls *list.List
}

func NewMyList() *MyList {
	return &MyList{
		ls : list.New(),
	}
}

func prepareList(n int) *MyList {

	myLs := NewMyList()
	for i :=1;i<=n;i++{
		myLs.ls.PushBack(i)
	}

	return myLs
}

func (m *MyList)getNum(index int) int {
	node := m.ls.Front()
	for i:=0;i<index;i++{
		node = node.Next()
	}
	
	m.ls.Remove(node)

	return node.Value.(int)
}

func handle(jc map[int]int, n int){

	prod := 1

	for i:=1;i<=n;i++{
		prod *= i
		jc[i] = prod
	}
}

func getPermutation(n int, k int) string {

	myLs := prepareList(n)

	jc := make(map[int]int)

	handle(jc, n)

	ans := ""
	p := n
	for m:=k-1;len(ans) < n - 1; {
		index:= m/jc[p-1]
		m = m%jc[p-1]
		ans += strconv.Itoa(myLs.getNum(index))
		p -= 1
	}

	ans += strconv.Itoa(myLs.getNum(0))

	return ans
}
// @lc code=end

