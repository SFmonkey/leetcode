/*
 * @lc app=leetcode.cn id=543 lang=javascript
 *
 * [543] 二叉树的直径
 *
 * https://leetcode-cn.com/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (49.08%)
 * Likes:    294
 * Dislikes: 0
 * Total Accepted:    41.7K
 * Total Submissions: 84.8K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
 * 
 * 
 * 
 * 示例 :
 * 给定二叉树
 * 
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \     
 * ⁠     4   5    
 * 
 * 
 * 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
 * 
 * 
 * 
 * 注意：两结点之间的路径长度是以它们之间边的数目表示。
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
var diameterOfBinaryTree = function(root) {
    let maxDepth = (root) => {
        if (root == null) return 0;
        return Math.max(maxDepth(root.left) + 1, maxDepth(root.right) + 1);
    }

    let fn = (root) => {
        if (root == null) return 0;
        let rootMax = maxDepth(root.left) + maxDepth(root.right);
        let childMax = Math.max(fn(root.left), fn(root.right));

        return Math.max(rootMax, childMax);
    }

    if (root == null) return 0;
    return fn(root);
};

var diameterOfBinaryTree = function(root) {
    let help = (node) => {
        if(node == null) return 0;

        let left = node.left ? help(node.left) + 1: 0;
        let right = node.right ? help(node.right) + 1: 0;

        let cur = left + right;

        if (max < cur) max = cur;

        return Math.max(left, right);
    }
    
    let max = 0;
    if (root == null) return 0;
    help(root);

    return max;
}
// @lc code=end

