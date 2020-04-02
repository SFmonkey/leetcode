/*
 * @lc app=leetcode.cn id=145 lang=javascript
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (70.78%)
 * Likes:    263
 * Dislikes: 0
 * Total Accepted:    63.7K
 * Total Submissions: 89.9K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 后序 遍历。
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
 * 输出: [3,2,1]
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
var postorderTraversal = function(root) {
    let res = [];

    let tarverse = (root) => {
        if (!root) return;
        tarverse(root.left);
        tarverse(root.right);
        res.push(root.val);
    }

    tarverse(root);
    return res;
};
var postorderTraversal = function(root) {
    if (!root) return [];
    let res = [], stack = [];
    let visited = new Set();
    let p = root;


    while(stack.length || p) {
        while(p) {
            stack.push(p);
            p = p.left;
        }

        let node = stack[stack.length - 1];

        if(node.right && !visited.has(node.right)) {
            p = node.right;
            visited.add(node.right);
        } else {
            res.push(node.val);
            stack.pop();
        }
    }

    return res;
}
// @lc code=end

