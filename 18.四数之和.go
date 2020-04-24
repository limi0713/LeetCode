/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (37.35%)
 * Likes:    423
 * Dislikes: 0
 * Total Accepted:    67.4K
 * Total Submissions: 180.3K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
 * + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 * 
 * 注意：
 * 
 * 答案中不可以包含重复的四元组。
 * 
 * 示例：
 * 
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 * 
 * 满足要求的四元组集合为：
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 * 
 * 
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)

	ans := make([][]int,0)

	for i:=0;i<len(nums)-3;i++{
		for j:=i+1;j<len(nums)-2;j++{
			sum := target - nums[i] - nums[j]

			l,r := j+1,len(nums)-1
			for l<r{
				if nums[l]+nums[r] > sum {
					r--
				}else if nums [l] + nums[r] < sum {
					l++
				}else {
					ansT := []int{nums[i],nums[j],nums[l],nums[r]}
					ans = append(ans,ansT)
					l++
					r--
				}
			}
		}
	}

	ansMap := make(map[string]([]int))

	for i := 0;i<len(ans);i++ {
		
		str := fmt.Sprintf("%d%d%d%d",ans[i][0],ans[i][1],ans[i][2],ans[i][3])

		ansMap[str] = ans[i]
	}

	res := make([][]int,0)

	for _,v := range ansMap {
		res = append(res,v)
	}

	return res

}

func isSame(a,b []int)bool{

	if len(a)!= len(b){
		return false
	}

	for i := 0;i < len(a);i++{
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// @lc code=end

