/*
 * @lc app=leetcode.cn id=124 lang=javascript
 *
 * [124] 二叉树中的最大路径和
 *
 * https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (39.73%)
 * Likes:    366
 * Dislikes: 0
 * Total Accepted:    31.5K
 * Total Submissions: 79.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个非空二叉树，返回其最大路径和。
 * 
 * 本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
 * 
 * 示例 1:
 * 
 * 输入: [1,2,3]
 * 
 * ⁠      1
 * ⁠     / \
 * ⁠    2   3
 * 
 * 输出: 6
 * 
 * 
 * 示例 2:
 * 
 * 输入: [-10,9,20,null,null,15,7]
 * 
 * -10
 * / \
 * 9  20
 * /  \
 * 15   7
 * 
 * 输出: 42
 * 
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
 * @return {number}
 */
var maxPathSum = function(root) {
    let res = [];

    let fn = (root) => {
        if (root == null) return 0;

        let leftVal = Math.max(fn(root.left), 0);
        let rightVal = Math.max(fn(root.right), 0);

        res.push(leftVal + rightVal + root.val);
        return Math.max(leftVal, rightVal) + root.val;
    }

    res.push(fn(root));
    return Math.max.apply(null, res);
};
var maxPathSum = function(root) {
    let maxsum = -Infinity;
    let fn = (root) => {
        if (root == null) return 0;

        let leftVal = Math.max(fn(root.left), 0);
        let rightVal = Math.max(fn(root.right), 0);
        let newVal = root.val + leftVal + rightVal;
        if (maxsum < newVal) maxsum = newVal

        return root.val + Math.max(rightVal, leftVal)
    }

    fn(root);
    return maxsum;
}
// @lc code=end

