/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 *
 * https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/description/
 *
 * algorithms
 * Hard (37.14%)
 * Likes:    174
 * Dislikes: 0
 * Total Accepted:    12.9K
 * Total Submissions: 34.6K
 * Testcase Example:  '[5,2,6,1]'
 *
 * 给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于
 * nums[i] 的元素的数量。
 * 
 * 示例:
 * 
 * 输入: [5,2,6,1]
 * 输出: [2,1,1,0] 
 * 解释:
 * 5 的右侧有 2 个更小的元素 (2 和 1).
 * 2 的右侧仅有 1 个更小的元素 (1).
 * 6 的右侧有 1 个更小的元素 (1).
 * 1 的右侧有 0 个更小的元素.
 * 
 * 
 */

// @lc code=start
func countSmaller(nums []int) []int {
    if len(nums)<=0 {
        return nil
    }

    counts := make([]int,len(nums))

    for i:= 1;i<len(nums);i++{
        for j:=0;j<i;j++{
            if nums[j]>nums[i]{
                counts[j]++
            }
        }
    }

    return counts
}
// @lc code=end

