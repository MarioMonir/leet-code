/*
 * @lc app=leetcode id=26 lang=javascript
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
  let k = 1;
  for (let i = 1; i < nums.length; i++) {
    if (nums[i] > nums[i - 1]) {
      nums[k] = nums[i];
      k += 1;
    }
  }
  return k;
};
// @lc code=end

// removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4]);
// removeDuplicates([1, 1, 2]);
// removeDuplicates([-1, 0, 0, 0, 0, 3, 3]);
