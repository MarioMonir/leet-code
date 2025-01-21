/*
 * @lc app=leetcode id=110 lang=javascript
 *
 * [110] Balanced Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isBalanced = function (root) {};
// @lc code=end

function TreeNode(value = 0, left = null, right = null) {
  return {
    value,
    left,
    right,
  };
}

let node5 = TreeNode(7, null, null);
let node4 = TreeNode(15, null, null);
let node3 = TreeNode(20, node4, node5);
let node2 = TreeNode(9, null, null);
let root = TreeNode(3, node2, node3);

// return;
function dfs(root) {
  function parseSubTree(node) {
    if (!node) return;
    console.log(node.value);
    parseSubTree(node.left);
    parseSubTree(node.right);
  }
  return parseSubTree(root);
}

console.log("DFS");
bfs(root);

// ============================================================

function bfs(root) {
  function parseSubTree(node) {
    if (!node) return;
    parseSubTree(node.left);
    parseSubTree(node.right);
    console.log(node.value);
  }
  return parseSubTree(root);
  x;
}

console.log("BFS");

dfs(root);
