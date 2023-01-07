/*
 * @lc app=leetcode id=2469 lang=javascript
 *
 * [2469] Convert the Temperature
 */

// @lc code=start
/**
 * @param {number} celsius
 * @return {number[]}
 */
var convertTemperature = function (celsius) {
  return [(celsius + 273.15).toFixed(5), (celsius * 1.8 + 32.0).toFixed(5)];
};
// @lc code=end

// console.log(convertTemperature(122.11));
// console.log(convertTemperature(36.5));
