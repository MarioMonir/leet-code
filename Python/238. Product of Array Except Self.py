from typing import List

class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        answer = [1] * len(nums)
        last_prefix = 1 
        for i in range(len(nums)):
            answer[i] = last_prefix
            last_prefix *= nums[i]
        
        last_postfix = 1
        for i in range(len(nums)-1,-1,-1):
            answer[i] *= last_postfix
            last_postfix *= nums[i]

        return answer


if __name__ == "__main__":
    solution = Solution()
    print(solution.productExceptSelf([1, 2, 3, 4]))