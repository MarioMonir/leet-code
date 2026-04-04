from typing import List

class Solution:
    def hIndex(self, citations: List[int]) -> int:
        citations.sort(reverse=True)
        h = 0
        for i in range(len(citations)):
            if citations[i] >= i+1:
                h+=1
        return h


if __name__ == "__main__":
    solution = Solution()
    print(solution.hIndex([3,0,6,1,5]))
    print(solution.hIndex([1,3,1]))