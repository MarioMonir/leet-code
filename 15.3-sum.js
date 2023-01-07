/*
 * @lc app=leetcode id=15 lang=javascript
 *
 * [15] 3Sum
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function (nums) {
  if (nums.length < 3) return [];

  if (nums.length === 3) {
    const sum = nums.reduce((a, b) => a + b, 0);
    if (sum === 0) return [nums];
    return [];
  }

  let result = [];

  // sort the array
  nums.sort((a, b) => a - b);

  for (let i = 0; i < nums.length - 2; i++) {
    // Already Checked this element
    if (i == 0 || (i > 0 && nums[i] == nums[i - 1])) continue;

    // ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

    let lower = i - 1;
    let higher = nums.length - 1;

    // ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

    while (lower < higher) {
      if (nums[i] + nums[lower] + nums[higher] === 0) {
        result.push([nums[i], nums[lower], nums[higher]]);

        // Prevenet duplicates higher and lowers
        while (lower < higher && nums[lower] + nums[lower + 1]) lower++;
        while (lower < higher && nums[higher] + nums[higher - 1]) higher--;

        lower++;
        higher--;
      } else if (nums[i] + nums[lower] + nums[higher] > 0) {
        higher--;
      } else {
        lower++;
      }
    }

    // ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
  }
  return result;
};
// @lc code=end
