/*
 * @lc app=leetcode.cn id=144 lang=javascript
 *
 * [144] 二叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
// var preorderTraversal = function(root) {
//     let arr = [];

//     let tarverse = (root) => {
//         if (root == null) return;

//         arr.push(root.val);
//         tarverse(root.left);
//         tarverse(root.right);
//     }

//     tarverse(root);
//     return arr;
// };
var preorderTraversal = function(root) {
    if (!root) return [];
    let arr = [];
    let stack = [root];

    while(stack.length) {
        let node = stack.pop();

        arr.push(node.val);
        node.right && stack.push(node.right);
        node.left && stack.push(node.left);
    }

    return arr;
}
// @lc code=end

