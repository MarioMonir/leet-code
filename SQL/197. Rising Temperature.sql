-- Write your PostgreSQL query statement below
SELECT 
    weatherToday.id
FROM 
    weather weatherYesterday 
JOIN 
    weather weatherToday
ON 
    weatherToday.recordDate = weatherYesterday.recordDate + INTERVAL '1 DAY'
    AND weatherToday.temperature > weatherYesterday.temperature
