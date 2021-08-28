-- https://leetcode.com/problems/trips-and-users/submissions */

SELECT 
  t.Request_at AS Day,  
  ROUND(AVG(CASE WHEN Status!='completed' THEN 1 ELSE 0 END), 2 ) AS 'Cancellation Rate'
FROM Trips t
JOIN Users d ON d.Users_Id = t.Driver_ID
JOIN Users c ON c.Users_Id = t.Client_ID
WHERE t.Request_at>='2013-10-01' AND t.Request_at<='2013-10-03'
  AND d.Banned='No' AND c.Banned='No'
GROUP BY 1
