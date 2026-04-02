from typing import List

class Solution:
    def reverse_in_place(self, nums:List[int], startIndex: int, endIndex: int)->None:
        if endIndex >= len(nums):
            return

        while startIndex < endIndex:
            nums[startIndex],nums[endIndex] = nums[endIndex],nums[startIndex]
            startIndex+=1
            endIndex-=1

    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        k = k % len(nums)
        self.reverse_in_place(nums,0,len(nums)-1)
        self.reverse_in_place(nums,0,k-1)
        self.reverse_in_place(nums,k,len(nums)-1)



        

if __name__ == "__main__":
    solution = Solution()
    nums = [1,2,3,4,5,6,7]
    k = 3
    solution.rotate(nums, k)
    print(nums)
    # [5, 6, 7, 1, 2, 3, 4]

    nums = [-1,-100,3,99]
    k = 2
    solution.rotate(nums, k)
    print(nums)
    # [3, 99, -1, -100]


    nums=[-1]
    solution.rotate(nums, 2)
    print(nums)
    # [-1]