-- https://leetcode.com/problems/rising-temperature/
SELECT w.id
FROM Weather w
INNER JOIN Weather w2 ON DATE_ADD(w2.recordDate, INTERVAL 1 DAY)=w.recordDate
where w.Temperature>w2.Temperature



--FASTER
WITH old_temp AS (
  SELECT id,
    LAG(recordDate, 1) OVER(ORDER BY recordDate) AS last_date,
    LAG(Temperature, 1) OVER (ORDER BY recordDate) AS last_temp
  FROM Weather
)
SELECT w.id FROM Weather w
INNER JOIN old_temp lg ON w.id = lg.id
WHERE w.Temperature > lg.last_temp and DATEDIFF(w.recordDate, lg.last_date)=1
