/*
 * @lc app=leetcode id=206 lang=javascript
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function (head) {
  if (!head || !head.next) return head;

  ptr1 = head;
  ptr2 = head.next;

  head.next = null;

  while (ptr2) {
    tempPointer = ptr1;
    ptr1 = ptr2;
    ptr2 = ptr2.next;
    ptr1.next = tempPointer;
  }

  return ptr1;
};
// @lc code=end
