
-- 500ms
SELECT Name AS Customers FROM CUSTOMERS
WHERE Id NOT IN
  (SELECT DISTINCT CustomerId from Orders)

