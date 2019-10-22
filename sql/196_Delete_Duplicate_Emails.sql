// 196. Delete Duplicate Emails
//   Keep only unique emails based on its smallest Id.

DELETE FROM Person
WHERE Id NOT IN
    (SELECT a.minId FROM
            (SELECT MIN(Id) AS minId FROM Person GROUP BY Email)
	         AS a);
		 
