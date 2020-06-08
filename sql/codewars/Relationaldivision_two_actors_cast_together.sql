# https://www.codewars.com/kata/
#  5th Kyu
SELECT title FROM film
WHERE film_id IN (SELECT film_id FROM film_actor WHERE actor_id=105) AND
      film_id IN (SELECT film_id FROM film_actor WHERE actor_id=122)
ORDER BY title;
