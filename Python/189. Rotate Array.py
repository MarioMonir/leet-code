from typing import List

class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        for i in range(k):
            last = nums[-1]
            for j in range(len(nums)-1, 0, -1):
                nums[j] = nums[j-1]
            nums[0] = last


        

if __name__ == "__main__":
    solution = Solution()
    nums = [1,2,3,4,5,6,7]
    k = 3
    solution.rotate(nums, k)
    # [5, 6, 7, 1, 2, 3, 4]
    print(nums)

    nums = [-1,-100,3,99]
    k = 2
    solution.rotate(nums, k)
    # [3, 99, -1, -100]
    print(nums)