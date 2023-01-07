/*
 * @lc app=leetcode id=258 lang=javascript
 *
 * [258] Add Digits
 */

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

/**
 *
 *  // Brute Force Solution
 *
 * var addDigits = function (num) {
 *    let sum = num;
 *    while (sum.toString().length > 1) {
 *      sum = sum
 *        .toString()
 *        .split("")
 *        .reduce((a, b) => +a + +b, 0);
 *      }
 *    return sum;
 *  };
 */

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

/**
 *  // The Mathmitcal Solution
 *  // Digital Root Theory
 *
 *
 * The digital root (also repeated digital sum) of a
 * natural number in a given radix is the (single digit)
 * value obtained by an iterative process of summing digits,
 * on each iteration using the result from the previous
 * iteration to compute a digit sum. The process continues
 * until a single-digit number is reached. For example,
 * in base 10, the digital root of the number 12345 is 6
 * because the sum of the digits in the number
 * is 1 + 2 + 3 + 4 + 5 = 15, then the addition process
 * is repeated again for the resulting number 15,
 * so that the sum of 1 + 5 equals 6,
 * which is the digital root of that number. In base 10,
 * this is equivalent to taking the remainder upon division by 9
 * (except when the digital root is 9,
 * where the remainder upon division by 9 will be 0),
 * which allows it to be used as a divisibility rule.
 *
 *
 * var addDigits = function (num) {
 *     return num && num % 9 === 0 ? 9 : num % 9;
 *   };
 *
 *
 */

// @lc code=start
/**
 * @param {number} num
 * @return {number}
 */
var addDigits = function (num) {
  return num && num % 9 === 0 ? 9 : num % 9;
};

// @lc code=end
