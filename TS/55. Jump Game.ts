function canJump(nums: number[]): boolean {
  let max_reach = 0;

  for (let i = 0; i < nums.length - 1; i++) {
    if (i > max_reach) return false;
    max_reach = Math.max(max_reach, i + nums[i]);
  }

  return max_reach >= nums.length - 1;
}

console.log(canJump([2, 3, 1, 1, 4]) == true); // true
console.log(canJump([3, 2, 1, 0, 4]) == false); // false
console.log(canJump([0]) == true); // true
console.log(canJump([0, 2, 3]) == false); // false
console.log(canJump([2, 5, 0, 0]) == true); // true
console.log(canJump([2, 0, 0]) == true); // true
