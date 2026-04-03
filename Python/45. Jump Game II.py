from typing import List

class Solution:
    def jump(self, nums: List[int]) -> int:
        max_reach = 0
        jumps = 0
        current_end = 0
        for i in range(len(nums)-1):
            max_reach = max(max_reach, i + nums[i])

            if current_end == i:
                jumps += 1
                current_end = max_reach

            
        return jumps
    

if __name__ == "__main__":
    solution = Solution()
    print(solution.jump([2,3,1,1,4])) # 2
    print(solution.jump([2,3,0,1,4])) # 2
    print(solution.jump([1,2,1,1,1])) # 3
    print(solution.jump([1,2,1,2,1,9])) # 3