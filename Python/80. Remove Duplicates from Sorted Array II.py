from typing import List

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        huge = (10**4)+1
        visited_value = huge 
        visited_times = 0
        k = len(nums)
        
        for i in range(len(nums)):
            if nums[i] == visited_value:
                visited_times += 1
                visited_value = nums[i]

                if visited_times > 2:
                    nums[i] = huge
                    k -=1

            else:
                visited_times = 1
                visited_value = nums[i]

        nums.sort()
        print(nums)
        return k


if __name__ == "__main__":
    solution = Solution()
    print(solution.removeDuplicates([1, 1, 1, 2, 2, 3]))