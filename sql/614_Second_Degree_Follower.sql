/* https://leetcode.com/problems/second-degree-follower */

WITH follower_count AS (
  SELECT followee, COUNT(DISTINCT follower) AS follower_count 
  FROM follow
  GROUP BY followee
)
    
SELECT DISTINCT f.follower, fc.follower_count AS num FROM follow f
LEFT JOIN follower_count fc ON f.follower = fc.followee
WHERE ISNULL(fc.follower_count) = FALSE
ORDER BY f.follower
