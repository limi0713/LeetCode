/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 *
 * https://leetcode-cn.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (33.06%)
 * Likes:    433
 * Dislikes: 0
 * Total Accepted:    52.4K
 * Total Submissions: 158.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
 * 
 * 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
 * 
 * 必须原地修改，只允许使用额外常数空间。
 * 
 * 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 * 
 */

// @lc code=start

func handlerMax(nums,maxNum []int)[]int{

	if len(nums)==0 {
		return []int{}
	}

	leng := len(nums)
	// maxNum = make([]int,leng)

	maxNum[leng-1] = leng - 1

	for i := len(maxNum)-2;i>=0;i--{
		if nums[i] > nums[maxNum[i+1]] {
			maxNum[i] = i
		}else {
			maxNum[i] = maxNum[i+1]
		}
	}

	return maxNum
}

func findMin(nums []int,start int,low int)int{
	ans := start
	for i := start+1;i<len(nums);i++{
		if nums[i] < nums[ans] && nums[i] > low{
			ans = i
		}
	}

	return ans
}

func nextPermutation(nums []int)  {
	if len(nums) <= 0 {
		return
	}

	maxNum := make([]int,len(nums))

	handlerMax(nums,maxNum)

	leng := len(nums)

	for i := leng -2; i >= 0; i--{
		if nums[i] >=nums[maxNum[i]] {
			continue
		}

		minIndex := findMin(nums,i+1,nums[i])

		nums[i],nums[minIndex] = nums[minIndex],nums[i]

		tempNums := nums[i+1:]
		sort.Ints(tempNums)

		nums = append(nums[:i+1],tempNums...)
		return
	}

	sort.Ints(nums)
}
// @lc code=end

