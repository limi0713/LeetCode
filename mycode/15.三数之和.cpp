/*
 * @lc app=leetcode.cn id=15 lang=cpp
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (23.36%)
 * Likes:    1269
 * Dislikes: 0
 * Total Accepted:    84.2K
 * Total Submissions: 355.6K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？找出所有满足条件且不重复的三元组。
 * 
 * 注意：答案中不可以包含重复的三元组。
 * 
 * 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 * 
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 * 
 * 
 */
class Solution
{
public:
    vector<vector<int>> threeSum(vector<int> &nums)
    {
        map<int, int> num_map;
        auto it = nums.begin();
        for (; it != nums.end(); it++)
        {
            num_map[*it] = *it;
        }
        vector<vector<int>> res;
        auto it_map = num_map.begin();
        for (; it_map != num_map.end(); it_map++)
        {
            auto it_next = it_map;
            for (it_next++; it_next != num_map.end(); it_next++)
            {
                auto it_temp = num_map.find(0 - (it_map->second + it_next->second));
                if (it_temp != num_map.end())
                {
                    vector<int> re;
                    re.push_back(it_map->second);
                    re.push_back(it_next->second);
                    re.push_back(it_temp->second);
                    res.push_back(re);
                }
            }
        }
        return res;
    }
};
