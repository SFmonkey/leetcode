/*
 * @lc app=leetcode.cn id=114 lang=javascript
 *
 * [114] 二叉树展开为链表
 *
 * https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (68.13%)
 * Likes:    305
 * Dislikes: 0
 * Total Accepted:    34.2K
 * Total Submissions: 50K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * 给定一个二叉树，原地将它展开为链表。
 * 
 * 例如，给定二叉树
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   5
 * ⁠/ \   \
 * 3   4   6
 * 
 * 将其展开为：
 * 
 * 1
 * ⁠\
 * ⁠ 2
 * ⁠  \
 * ⁠   3
 * ⁠    \
 * ⁠     4
 * ⁠      \
 * ⁠       5
 * ⁠        \
 * ⁠         6
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
 * @return {void} Do not return anything, modify root in-place instead.
 */
var flatten = function(root) {
    if(root == null) return;
    flatten(root.left);
    flatten(root.right);
    if(root.left) {
        let p = root.left;
        while(p.right) {
            p = p.right;
        }
        p.right = root.right;
        root.right = root.left;
        root.left = null;
    }
};

var flatten = function(root) {
    if(root == null) return;
    let stack = [];
    let visited = new Set();
    let p = root;
    // 开始后序遍历
    while(stack.length || p) {
        while(p) {
            stack.push(p);
            p = p.left;
        }
        let node = stack[stack.length - 1];
        // 如果右孩子存在，而且右孩子未被访问
        if(node.right && !visited.has(node.right)) {
            p = node.right;
            visited.add(node.right);
        } else {
            if(node.left) {
                let p = node.left;
                while(p.right) {
                    p = p.right;
                }
                p.right = node.right;
                node.right = node.left;
                node.left = null;
            }
            stack.pop();
        }
    }
};
// @lc code=end

