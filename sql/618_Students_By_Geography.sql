/* https://leetcode.com/problems/students-report-by-geography */

WITH America AS (
  SELECT ROW_NUMBER() OVER (ORDER BY name) AS rw
       , name from student WHERE continent='America'
  )
, Asia AS (
  SELECT ROW_NUMBER() OVER (ORDER BY name) AS rw
       , name from student WHERE continent='Asia'
  )
, Europe AS (
  SELECT ROW_NUMBER() OVER (ORDER BY name) AS rw
       , name from student WHERE continent='Europe'
)
  
SELECT am.name AS America,
  aa.name as Asia,
  e.name as Europe
FROM America am
LEFT JOIN Asia aa ON aa.rw = am.rw
LEFT JOIN Europe e on e.rw = am.rw
