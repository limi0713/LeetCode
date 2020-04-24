--
-- @lc app=leetcode.cn id=197 lang=mysql
--
-- [197] 上升的温度
--

-- @lc code=start
# Write your MySQL query statement below

select tt.Id as id from Weather t, Weather tt 
where DATEDIFF(tt.RecordDate,t.RecordDate)=1 and tt.Temperature > t.Temperature;

-- @lc code=end

