from typing import List

class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        count = 0
        isFirstWordVisited = False 

        for i in range(len(s)-1,-1,-1):
            if s[i] == " " and isFirstWordVisited:
                return count
            
            if s[i] == " " and not isFirstWordVisited:
                continue

            if not isFirstWordVisited:
                isFirstWordVisited = True

            count += 1

        return count

   

if __name__ == "__main__":
    solution = Solution()
    print(solution.lengthOfLastWord("Hello World"))
    print(solution.lengthOfLastWord("   fly me   to   the moon  "))
    print(solution.lengthOfLastWord("luffy is still joyboy"))