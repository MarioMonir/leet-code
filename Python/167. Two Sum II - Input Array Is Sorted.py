from typing import List

class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        start = 0
        end = len(numbers) - 1

        while start < end :
            sum = numbers[start] + numbers[end]

            if sum == target :
                return  [start+1,end+1]

            if sum > target :
                end-=1
            else:
                start +=1
    
        return []

    
if __name__ == "__main__":
    solution = Solution()
    print(solution.twoSum([2,7,11,15], 9)==[1,2])
    print(solution.twoSum([2,3,4], 6)==[1,3])
    print(solution.twoSum([-1,0], -1)==[1,2])
    print(solution.twoSum([0,0,3,4], 0)==[1,2])
    print(solution.twoSum([0,0,3,4], 9)==[])