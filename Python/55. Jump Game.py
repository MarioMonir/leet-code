from typing import List

class Solution:
    def canJump(self, nums: List[int]) -> bool:
        max_reach = 0

        for i in range(len(nums)-1):
            if i > max_reach:
                return False
            max_reach = max(max_reach, i + nums[i])

        return max_reach >= len(nums)-1

      

if __name__ == "__main__":
    solution = Solution()
    print(solution.canJump([2,3,1,1,4])) # True
    print(solution.canJump([3,2,1,0,4])) # False
    print(solution.canJump([0])) # True
    print(solution.canJump([0,2,3])) # False
    print(solution.canJump([2,5,0,0])) # True
    print(solution.canJump([2,0,0])) # True
