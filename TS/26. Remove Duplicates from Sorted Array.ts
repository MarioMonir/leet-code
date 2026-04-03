function removeDuplicates(nums: number[]): number {
  let unique: number = 0;
  let visitedMap: { [key: number]: boolean } = {};

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] in visitedMap) {
      nums[i] = 101;
    } else {
      visitedMap[nums[i]] = true;
      unique++;
    }
  }

  // ascending order
  nums = nums.sort((a, b) => a - b);

  return unique;
}

console.log(removeDuplicates([1, 1, 2]));
