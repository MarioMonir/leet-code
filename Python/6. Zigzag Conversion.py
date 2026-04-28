class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows >= len(s) or numRows == 1:
            return s


        zigzag_arr = [""] * numRows
        array_index = 0 
        adder = 1 


        for i in range(len(s)):
            zigzag_arr[array_index] += s[i]
            array_index += adder

            if array_index == 0:
                adder = 1
            if array_index == numRows - 1:
                adder = -1

        return "".join(zigzag_arr)


if __name__ == "__main__":
    solution = Solution()
    print(solution.convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
    print(solution.convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
    print(solution.convert("A", 1) == "A")
    print(solution.convert("AB", 1) == "AB")
    print(solution.convert("AB", 2) == "AB")
    print(solution.convert("ABCDEF", 7) == "ABCDEF")
