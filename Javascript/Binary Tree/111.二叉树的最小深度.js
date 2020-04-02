/*
 * @lc app=leetcode.cn id=111 lang=javascript
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (41.85%)
 * Likes:    237
 * Dislikes: 0
 * Total Accepted:    62.9K
 * Total Submissions: 150.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 * 
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 返回它的最小深度  2.
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
var minDepth = function(root) {
    // 递归终止条件 
    if(root == null) return 0;

    if (root.left && root.right) {
        return Math.min(minDepth(root.left) + 1, minDepth(root.right)+1);
    } else if (root.left) {
        return minDepth(root.left) + 1;
    } else if (root.right) {
        return minDepth(root.right) + 1;
    } else {
        return 1;
    }
};

var minDepth = function(root) {
    if (!root) return 0;
    let queue = [root];
    let level = 0;

    while(queue.length) {
        let size = queue.length;
        while (size --) {
            let node = queue.shift();
            if (!node.left && !node.right) return level + 1;
            node.left && queue.push(node.left);
            node.right && queue.push(node.right);
        }
        level ++;
    }

    return level;
};
// @lc code=end

