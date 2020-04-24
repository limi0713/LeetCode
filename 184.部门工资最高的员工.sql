--
-- @lc app=leetcode.cn id=184 lang=mysql
--
-- [184] 部门工资最高的员工
--
-- https://leetcode-cn.com/problems/department-highest-salary/description/
--
-- database
-- Medium (44.75%)
-- Likes:    230
-- Dislikes: 0
-- Total Accepted:    34.6K
-- Total Submissions: 77.1K
-- Testcase Example:  '{"headers": {"Employee": ["Id", "Name", "Salary", "DepartmentId"], "Department": ["Id", "Name"]}, "rows": {"Employee": [[1, "Joe", 70000, 1], [2, "Jim", 90000, 1], [3, "Henry", 80000, 2], [4, "Sam", 60000, 2], [5, "Max", 90000, 1]], "Department": [[1, "IT"], [2, "Sales"]]}}'
--
-- Employee 表包含所有员工信息，每个员工有其对应的 Id, salary 和 department Id。
-- 
-- +----+-------+--------+--------------+
-- | Id | Name  | Salary | DepartmentId |
-- +----+-------+--------+--------------+
-- | 1  | Joe   | 70000  | 1            |
-- | 2  | Henry | 80000  | 2            |
-- | 3  | Sam   | 60000  | 2            |
-- | 4  | Max   | 90000  | 1            |
-- +----+-------+--------+--------------+
-- 
-- 
-- Department 表包含公司所有部门的信息。
-- 
-- +----+----------+
-- | Id | Name     |
-- +----+----------+
-- | 1  | IT       |
-- | 2  | Sales    |
-- +----+----------+
-- 
-- 
-- 编写一个 SQL 查询，找出每个部门工资最高的员工。例如，根据上述给定的表格，Max 在 IT 部门有最高工资，Henry 在 Sales
-- 部门有最高工资。
-- 
-- +------------+----------+--------+
-- | Department | Employee | Salary |
-- +------------+----------+--------+
-- | IT         | Max      | 90000  |
-- | Sales      | Henry    | 80000  |
-- +------------+----------+--------+
-- 
-- 
--

-- @lc code=start
# Write your MySQL query statement below

select
    Department.Name as Department,
    Employee.Name as Employee,
    Employee.Salary as Salary
from
    Employee
    join
    Department
    on Employee.DepartmentId = Department.Id
where 
    (Employee.DepartmentId,Employee.Salary)
    in
    (
        select DepartmentId,Max(Salary)
        from Employee
        group by DepartmentId
    );
-- @lc code=end

