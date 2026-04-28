class Solution:
    def reverseWords(self, s: str) -> str:
        return " ".join(reversed(s.strip().split()))



if __name__ == "__main__":
    solution = Solution()
    print(solution.reverseWords("the sky is blue") == "blue is sky the" )
    print(solution.reverseWords("  hello world  ") == "world hello")
    print(solution.reverseWords("a good   example") == "example good a")