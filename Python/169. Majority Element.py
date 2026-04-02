from typing import List

class Solution:
    def majorityElement(self, nums: List[int]) -> int:
       countMap = {}
       majority = nums[0]

       for i in range(len(nums)):
        if nums[i] in countMap:
            countMap[nums[i]] += 1
        else:
            countMap[nums[i]] = 1

       
        for key, value in countMap.items():
            if value > countMap[majority]:
                majority = key

       return majority



if __name__ == "__main__":
    solution = Solution()
    print(solution.majorityElement([3, 2, 3]))