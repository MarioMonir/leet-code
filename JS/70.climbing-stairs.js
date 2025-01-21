/*
 * @lc app=leetcode id=70 lang=javascript
 *
 * [70] Climbing Stairs
 */

// @lc code=start
/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function (n) {
  // Fibonacci Number Sequence recursion
  //   function recursiveFibonacci(n) {
  //     if (n === 0 || n == 1) return n;
  //     else return recursiveFibonacci(n - 2) + recursiveFibonacci(n - 1);
  //   }

  // Fibonacci Number Seque by loop
  function fibonacci(n) {
    let a = 0;
    let b = 1;
    let fib = 1;

    for (let i = 2; i <= n; i++) {
      fib = a + b;
      a = b;
      b = fib;
    }
    return fib;
  }

  //   return recursiveFibonacci(n + 1);
  return fibonacci(n + 1);
};
// @lc code=end

// golden ratio = 1.618034
// , Fn = (Φ**n - (1-Φ)**n )/√5

let goldenRatio = 1.618034;
let fiboEquation = (n) => {
  return parseInt(
    ((goldenRatio ** n - (-goldenRatio) ** -n) / Math.sqrt(5)).toFixed(1)
  );
};

for (let i = 0; i < 10; i++) {
  console.log(i, fiboEquation(i), climbStairs(i));
}
