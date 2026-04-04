import random

class RandomizedSet:

    def __init__(self):
        self.value_map = {}

    def insert(self, val: int) -> bool:
        if val in self.value_map:
            return False

        self.value_map[val] = 1
        return True

    def remove(self, val: int) -> bool:
        if val not in self.value_map:
            return False

        del self.value_map[val]
        return True

        
    def getRandom(self) -> int:
        return random.choice(list(self.value_map.keys()))
        


if __name__ == "__main__":
    randomizedSet = RandomizedSet()
    # ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
    #  [[], [1], [2], [2], [], [1], [2], []]

    # Output
    # [null, true, false, true, 2, true, false, 2]    
    print(randomizedSet.insert(1))
    print(randomizedSet.remove(2))
    print(randomizedSet.insert(2))
    print(randomizedSet.getRandom())
    print(randomizedSet.remove(1))
    print(randomizedSet.insert(2))
    print(randomizedSet.getRandom())