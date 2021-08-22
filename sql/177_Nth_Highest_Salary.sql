/* Write a SQL query to get the nth highest salary from the Employee table.
    https://leetcode.com/problems/nth-highest-salary  */

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
RETURN (
  SELECT IFNULL(
      (SELECT Salary
          FROM (SELECT SALARY, DENSE_RANK() OVER(ORDER BY Salary DESC) as s_rank FROM Employee) as sub1
          WHERE s_rank = N LIMIT 1),
      NULL)
);
END



// PRACTICE
SELECT IFNULL(
  (SELECT Salary
   FROM (SELECT SALARY, DENSE_RANK() OVER(ORDER BY Salary DESC) as s_rank FROM Employee) as sub1
         WHERE s_rank = N LIMIT 1),
   NULL)
