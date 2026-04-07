from typing import List

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        sell_index, buy_index, profit = 1 ,0 , 0
        while sell_index < len(prices):
            if prices[sell_index] > prices[buy_index]:
                profit+=prices[sell_index] - prices[buy_index]
                buy_index = sell_index
            else:
                buy_index+=1

            sell_index += 1

        return profit
        

if __name__ == "__main__":
    solution = Solution()
    print(solution.maxProfit([7,1,5,3,6,4]))
    print(solution.maxProfit([1,2,3,4,5]))