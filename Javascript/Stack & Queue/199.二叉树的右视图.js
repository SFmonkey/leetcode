/*
 * @lc app=leetcode.cn id=199 lang=javascript
 *
 * [199] 二叉树的右视图
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
var rightSideView = function(root) {
    if (!root) return [];
    let queue = [root];
    let res = [];

    while(queue.length) {
        let size = queue.length;
        res.push(queue[0].val);
        while(size --) {
            let front = queue.shift();

            if (front.right) queue.push(front.right);
            if (front.left) queue.push(front.left);
        }
    }

    return res;
};
// @lc code=end

