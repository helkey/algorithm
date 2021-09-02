/* Find employees who have the highest salary in each of the departments
     https://leetcode.com/problems/department-highest-salary/submissions/ */

-- WITHOUT Window Functions
WITH max_sal AS (
  Select DepartmentId, MAX(Salary) Salary
    FROM Employee
      GROUP BY 1
      )

SELECT d.Name Department, e.Name Employee, e.Salary
FROM Employee e
INNER JOIN Department d ON d.Id=e.DepartmentId
INNER JOIN max_sal mx ON mx.DepartmentId=e.DepartmentId
WHERE e.Salary=mx.Salary
ORDER BY 1,2


-- WITH Window Functions
SELECT d.Name as Department, r.Name as Employee, r.Salary FROM
  (SELECT
       DepartmentId,
       Name,
       Salary,
       RANK() OVER(PARTITION BY DepartmentId ORDER BY Salary DESC) AS rnk
   FROM Employee) r
INNER JOIN Department d ON r.DepartmentId = d.Id
WHERE rnk = 1
