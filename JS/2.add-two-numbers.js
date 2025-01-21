/*
 * @lc app=leetcode id=2 lang=javascript
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
function ListNode(val, next) {
  this.val = val === undefined ? 0 : val;
  this.next = next === undefined ? null : next;
}

/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {
  let sums = new ListNode();
  let ptr = sums;
  let carrier = 0;

  while (l1 || l2 || carrier) {
    let v1 = l1?.val || 0;
    let v2 = l2?.val || 0;

    // Add a new digit
    let sum = v1 + v2 + carrier;
    carrier = parseInt(sum / 10);
    unit = sum % 10;
    ptr.next = new ListNode(unit);

    // Increment the pointers positions
    ptr = ptr.next;
    l1 = l1?.next;
    l2 = l2?.next;
  }

  // console.dir({ sums }, { depth: Infinity });
  return sums.next;
};

// @lc code=end
