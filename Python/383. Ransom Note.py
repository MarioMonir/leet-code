class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        hash = {}

        # store all ransomNote letters in the hash
        for i in range(len(ransomNote)):
            if ransomNote[i] in hash:
                hash[ransomNote[i]] += 1
            else:
                hash[ransomNote[i]] = 1

        # check if magazine can construct ransomNote
        for i in range(len(magazine)):
            if magazine[i] in hash:
                if hash[magazine[i]] == 1:
                    del hash[magazine[i]]
                else:
                    hash[magazine[i]] -= 1

        return len(hash.values()) == 0



if __name__ == "__main__":
    solution = Solution()
    print(solution.canConstruct("a", "b"))
    print(solution.canConstruct("aa", "ab"))
    print(solution.canConstruct("aa", "aab"))