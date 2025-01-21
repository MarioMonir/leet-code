--
-- @lc app=leetcode id=1587 lang=mysql
--
-- [1587] Bank Account Summary II
--

-- @lc code=start
# Write your MySQL query statement below
SELECT name , sum(amount) as balance
FROM Users 
INNER JOIN Transactions ON Users.account = Transactions.account
GROUP BY Users.name
HAVING balance > 10000

-- @lc code=end

