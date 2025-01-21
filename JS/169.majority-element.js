/*
 * @lc app=leetcode id=169 lang=javascript
 *
 * [169] Majority Element
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function (nums) {
  // - - - - - - - - - - - - - - - - - - -

  //   // Brute force algorithm
  //   function bruteForce(nums) {
  //     if (nums.length === 1) return nums[0];

  //     let hashCount = {};
  //     for (let i = 0; i < nums.length; i++) {
  //       if (hashCount.hasOwnProperty(nums[i])) {
  //         hashCount[nums[i]]++;

  //         // found the majority element
  //         if (hashCount[nums[i]] >= nums.length / 2) {
  //           return nums[i];
  //         }
  //       } else {
  //         hashCount[nums[i]] = 1;
  //       }
  //     }
  //   }

  // - - - - - - - - - - - - - - - - - - -

  // Boyer–Moore majority vote algorithm
  function boyerMooreMajorityVote(nums) {
    let major = nums[0];
    let count = 1;

    for (let i = 1; i < nums.length; i++) {
      if (count === 0) {
        major = nums[i];
      }

      if (nums[i] === major) {
        count++;
      } else {
        count--;
      }
    }

    return major;
  }

  // - - - - - - - - - - - - - - - - - - -

  //   return bruteForce(nums);
  return boyerMooreMajorityVote(nums);
};
// @lc code=end
