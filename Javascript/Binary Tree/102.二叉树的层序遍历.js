/*
 * @lc app=leetcode.cn id=102 lang=javascript
 *
 * [102] 二叉树的层序遍历
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
var levelOrder = function(root) {
    if (!root) return [];
    let res = [];
    let level = 0;
    let queue = [root];
    while(queue.length) {
        res.push([]);
        let size = queue.length;
        while (size --) {
            let node = queue.shift();
            res[level].push(node.val);
            if (node.left) queue.push(node.left);
            if (node.right) queue.push(node.right);
        }
        level ++;
    }

    return res;
};
// @lc code=end

