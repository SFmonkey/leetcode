/*
 * @lc app=leetcode.cn id=103 lang=javascript
 *
 * [103] 二叉树的锯齿形层次遍历
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
 * @return {number[][]}
 */
var zigzagLevelOrder = function(root) {
    if (!root) return [];
    let queue = [root];
    let res = [];
    let level = 0;
    while (queue.length) {
        res.push([]);
        let size = queue.length;
        while(size --) {
            let front = queue.shift();
            res[level].push(front.val);

            if (front.left) queue.push(front.left);
            if (front.right) queue.push(front.right);
        }
        if (level % 2) res[level].reverse()
        level ++;
    }

    return res;
};
// @lc code=end

