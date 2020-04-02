/*
 * @lc app=leetcode.cn id=94 lang=javascript
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (70.75%)
 * Likes:    442
 * Dislikes: 0
 * Total Accepted:    131.5K
 * Total Submissions: 185.7K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
 * 
 * 示例:
 * 
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 * 
 * 输出: [1,3,2]
 * 
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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
 * @return {number[]}
 */
var inorderTraversal = function(root) {
    let res = [];

    let tarverse = (root) => {
        if (root == null) return;
        tarverse(root.left);
        res.push(root.val);
        tarverse(root.right);
    }

    tarverse(root);
    return res;
};

var inorderTraversal = function(root) {
    if (!root) return [];
    
    let res = [], stack = [];
    let p = root;

    while(stack.length || p) {
        while(p) {
            stack.push(p);
            p = p.left;
        }

        let node = stack.pop();
        res.push(node.val);
        p = node.right;
    }

    return res;
}
// @lc code=end

