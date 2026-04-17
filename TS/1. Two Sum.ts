function twoSum(nums: number[], target: number): number[] {
  const hashMap: { [key: number]: number } = {};

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] in hashMap) {
      return [hashMap[nums[i]], i];
    }

    hashMap[target - nums[i]] = i;
  }

  return [];
}

console.log(twoSum([2, 7, 11, 15], 9));
console.log(twoSum([3, 2, 4], 6));
