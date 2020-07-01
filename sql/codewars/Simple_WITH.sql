/* */

WITH special_sales AS (SELECT DISTINCT department_id from sales WHERE price > 90.)
SELECT id, name FROM departments
WHERE id IN (SELECT department_id FROM special_sales);


/* ALTERNATE SOLUTION - INNER JOIN */
WITH special_sales AS (
  SELECT DISTINCT department_id
    FROM sales
      WHERE price > 90.0
      )
SELECT id, name FROM departments
INNER JOIN special_sales
  ON special_sales.department_id  = departments.id;
