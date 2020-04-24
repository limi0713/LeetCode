--
-- @lc app=leetcode.cn id=596 lang=mysql
--
-- [596] 超过5名学生的课
--
-- https://leetcode-cn.com/problems/classes-more-than-5-students/description/
--
-- database
-- Easy (38.99%)
-- Likes:    125
-- Dislikes: 0
-- Total Accepted:    24.8K
-- Total Submissions: 63.2K
-- Testcase Example:  '{"headers": {"courses": ["student", "class"]}, "rows": {"courses": [["A", "Math"], ["B", "English"], ["C", "Math"], ["D", "Biology"], ["E", "Math"], ["F", "Computer"], ["G", "Math"], ["H", "Math"], ["I", "Math"]]}}'
--
-- 有一个courses 表 ，有: student (学生) 和 class (课程)。
-- 
-- 请列出所有超过或等于5名学生的课。
-- 
-- 例如,表:
-- 
-- 
-- +---------+------------+
-- | student | class      |
-- +---------+------------+
-- | A       | Math       |
-- | B       | English    |
-- | C       | Math       |
-- | D       | Biology    |
-- | E       | Math       |
-- | F       | Computer   |
-- | G       | Math       |
-- | H       | Math       |
-- | I       | Math       |
-- +---------+------------+
-- 
-- 
-- 应该输出:
-- 
-- 
-- +---------+
-- | class   |
-- +---------+
-- | Math    |
-- +---------+
-- 
-- 
-- Note:
-- 学生在每个课中不应被重复计算。
-- 
--

-- @lc code=start
# Write your MySQL query statement below

select class from courses group by class having count(distinct student) >=5;

-- @lc code=end

