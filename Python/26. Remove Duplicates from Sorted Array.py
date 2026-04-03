from typing import List

class Solution:
    def removeDuplicates_brute_force(self, nums: List[int]) -> int:
        unique = 0
        visited_map = {}
        for i in range(len(nums)):
            if nums[i] in visited_map:
                nums[i] = 101
            else:
                visited_map[nums[i]] = True
                unique += 1
        nums.sort()
        return unique
        
    
    def removeDuplicates_memory_optimization(self, nums: List[int]) -> int:
        current = 101
        unique = 0
        for i in range(len(nums)):
            if nums[i] == current:
                nums[i] = 101
            else:
                current = nums[i]
                unique += 1

        nums.sort()
        return unique

    def removeDuplicates(self, nums: List[int]) -> int:
        step = 0
        for i in range(1, len(nums)):
            if nums[step] != nums[i]:
                step += 1
                nums[step] = nums[i]
        return step + 1
         




if __name__ == "__main__":
    solution = Solution()
    print(solution.removeDuplicates([1, 1, 2]))