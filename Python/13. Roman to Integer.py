from typing import List

class Solution:
    def romanToInt(self, s: str) -> int:
        roman_to_int = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000
        }
        total = 0

        for i in range(len(s)):
            if i+1<len(s) and roman_to_int[s[i]] < roman_to_int[s[i+1]]:
                total-=roman_to_int[s[i]]
            else:
                total+=roman_to_int[s[i]]
        return total

if __name__ == "__main__":
    solution = Solution()
    print("III == 3", solution.romanToInt("III") == 3) # 3
    print("IV == 4", solution.romanToInt("IV") == 4) # 4
    print("IX == 9", solution.romanToInt("IX") == 9) # 9
    print("LVIII == 58", solution.romanToInt("LVIII") == 58) # 58
    print("MCMXCIV == 1994", solution.romanToInt("MCMXCIV") == 1994) # 1994