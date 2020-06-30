/* 6 kyu: SQL Basics: Simple HAVING
    https://www.codewars.com/kata/58164ddf890632ce00000220/train/sql */

SELECT age, COUNT(*) AS total_people FROM people AS p1
GROUP BY age
HAVING COUNT(*) >= 10


