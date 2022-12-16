--
-- @lc app=leetcode.cn id=183 lang=mysql
--
-- [183] 从不订购的客户
--
-- https://leetcode-cn.com/problems/customers-who-never-order/description/
--
-- database
-- Easy (66.16%)
-- Likes:    149
-- Dislikes: 0
-- Total Accepted:    56.6K
-- Total Submissions: 85.5K
-- Testcase Example:  '{"headers": {"Customers": ["Id", "Name"], "Orders": ["Id", "CustomerId"]}, "rows": {"Customers": [[1, "Joe"], [2, "Henry"], [3, "Sam"], [4, "Max"]], "Orders": [[1, 3], [2, 1]]}}'
--
-- 某网站包含两个表，Customers 表和 Orders 表。编写一个 SQL 查询，找出所有从不订购任何东西的客户。
-- 
-- Customers 表：
-- 
-- +----+-------+
-- | Id | Name  |
-- +----+-------+
-- | 1  | Joe   |
-- | 2  | Henry |
-- | 3  | Sam   |
-- | 4  | Max   |
-- +----+-------+
-- 
-- 
-- Orders 表：
-- 
-- +----+------------+
-- | Id | CustomerId |
-- +----+------------+
-- | 1  | 3          |
-- | 2  | 1          |
-- +----+------------+
-- 
-- 
-- 例如给定上述表格，你的查询应返回：
-- 
-- +-----------+
-- | Customers |
-- +-----------+
-- | Henry     |
-- | Max       |
-- +-----------+
-- 
-- 
--

-- @lc code=start
# Write your MySQL query statement below

select c.Name as Customers
from Customers as c
where c.Id not in
(
    select distinct CustomerId from orders
);

-- @lc code=end

