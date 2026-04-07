class Solution:
    def intToRoman(self, num: int) -> str:
        symbols = [
            ('I',1),
            ('IV', 4),
            ('V', 5),
            ('IX', 9),
            ('X', 10),
            ('XL', 40),
            ('L', 50),
            ('XC', 90),
            ('C', 100),
            ('CD', 400),
            ('D', 500),
            ('CM', 900),
            ('M', 1000),
        ]

        decimal=""

        for roman_symbol , decimal_value in reversed(symbols):
            if num // decimal_value:
                count = num // decimal_value
                decimal += (roman_symbol * count)
                num = num % decimal_value

        return decimal

if __name__ == "__main__":
    solution = Solution()
    print(solution.intToRoman(3) == "III")
    print(solution.intToRoman(4) == "IV")
    print(solution.intToRoman(9) == "IX")
    print(solution.intToRoman(58) == "LVIII")
    print(solution.intToRoman(1994) == "MCMXCIV")
