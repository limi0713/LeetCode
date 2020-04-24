--
-- @lc app=leetcode.cn id=181 lang=mysql
--
-- [181] 超过经理收入的员工
--

-- @lc code=start
# Write your MySQL query statement below

select a.name as employee from employee a 
where a.Salary>(
    select b.Salary from employee b 
    where b.id = a.managerid);

-- @lc code=end

