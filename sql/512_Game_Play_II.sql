

WITH first_event AS (
  SELECT player_id, MIN(event_date) as first_date FROM Activity
  GROUP BY player_id)

SELECT a.player_id, a.device_id FROM Activity a
  WHERE a.event_date =  (SELECT first_date FROM first_event WHERE a.player_id = first_event.player_id)
