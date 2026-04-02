from typing import List

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        buy_index, sell_index, profit = 0, 1, 0

        while sell_index < len(prices):
            if prices[sell_index] > prices[buy_index]:
                profit = max(profit, prices[sell_index] - prices[buy_index])
            else:
                buy_index = sell_index
            sell_index += 1

        return profit



if __name__ == "__main__":
    solution = Solution()
    print(solution.maxProfit([7,1,5,3,6,4]))
    print(solution.maxProfit([7,6,4,3,1]))
    