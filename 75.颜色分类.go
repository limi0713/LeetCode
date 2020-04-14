/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int)  {
	if len(nums) <= 0 {
		return
	}

	right0,left2 := 0, len(nums)-1
	for i := 0; i < len(nums) && i <= left2;i ++{
		if nums[i] == 0 {
			nums[i], nums[right0] = nums[right0], nums[i]
			right0 ++
		}else if nums[i] == 2 {
			nums[i], nums[left2] = nums[left2], nums[i]
			left2 --
			i -- // 将后面没有处理的数交换到前面，需要回退一步
		}
	}
}
// @lc code=end

