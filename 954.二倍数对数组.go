/*
 * @lc app=leetcode.cn id=954 lang=golang
 *
 * [954] 二倍数对数组
 *
 * https://leetcode-cn.com/problems/array-of-doubled-pairs/description/
 *
 * algorithms
 * Medium (27.51%)
 * Likes:    18
 * Dislikes: 0
 * Total Accepted:    2.6K
 * Total Submissions: 9.6K
 * Testcase Example:  '[3,1,3,6]'
 *
 * 给定一个长度为偶数的整数数组 A，只有对 A 进行重组后可以满足 “对于每个 0 <= i < len(A) / 2，都有 A[2 * i + 1] =
 * 2 * A[2 * i]” 时，返回 true；否则，返回 false。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：[3,1,3,6]
 * 输出：false
 * 
 * 
 * 示例 2：
 * 
 * 输入：[2,1,2,6]
 * 输出：false
 * 
 * 
 * 示例 3：
 * 
 * 输入：[4,-2,2,-4]
 * 输出：true
 * 解释：我们可以用 [-2,-4] 和 [2,4] 这两组组成 [-2,-4,2,4] 或是 [2,4,-2,-4]
 * 
 * 示例 4：
 * 
 * 输入：[1,2,4,16,8,4]
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= A.length <= 30000
 * A.length 为偶数
 * -100000 <= A[i] <= 100000
 * 
 * 
 */

// @lc code=start
func sort(arr []int){
	if len(arr) <= 0 {
		return
	}

	ls := list.New()
	ls.PushBack(arr[0])

	for i := 1; i < len(arr); i ++ {
		end := ls.Back()

		for end != nil && math.Abs(float64(end.Value.(int))) > math.Abs(float64(arr[i])) {
			end = end.Prev()
		}

		if end == nil {
			ls.PushFront(arr[i])
		} else {
			ls.InsertAfter(arr[i], end)
		}
	}

	start := ls.Front()
	for i := 0; i < len(arr); i ++ {
		arr[i] = start.Value.(int)
		start = start.Next()
	}
}

func canReorderDoubled(A []int) bool {
	sort(A)

	numsMap := make(map[int]int)

	for i := range A {
		numsMap[A[i]] = numsMap[A[i]] + 1
	}

	for i := range A {
		if v, ok := numsMap[A[i]]; ok && v > 0 {

			if v2, ok2 := numsMap[2 * A[i]]; ok2 && v2 > 0 {
				numsMap[2*A[i]] --
				numsMap[A[i]] -- 
			}else {
				return false
			}
		} 
	}

	return true
}
// @lc code=end

