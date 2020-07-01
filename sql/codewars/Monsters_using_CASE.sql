/* https://www.codewars.com/kata/593ef0e98b90525e090000b9/train/sql */
/* https://www.codewars.com/kata/593ef0e98b90525e090000b9/train/sql */
SELECT th.id, th.heads, bh.legs, th.arms, bh.tails,
   CASE
      WHEN heads > arms OR tails > legs THEN 'BEAST'
      ELSE 'WEIRDO'
   END AS species
  FROM top_half th
  
INNER JOIN bottom_half bh ON th.id = bh.id
ORDER BY species


/* SOLUTION */
SELECT
  th.id,
  th.heads,
  bh.legs,
  th.arms,
  bh.tails,
  CASE WHEN th.heads > th.arms or bh.tails > bh.legs
       THEN 'BEAST'
       ELSE 'WEIRDO'
  END as species

FROM top_half th

INNER JOIN bottom_half bh
  on bh.id = th.id

ORDER BY species
