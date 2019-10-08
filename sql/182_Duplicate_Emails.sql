List all emails that appear as duplicates
*** SELECT Email
FROM Person
GROUP BY Email
HAVING COUNT(Email) > 1;
