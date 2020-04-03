/*
 * @lc app=leetcode.cn id=257 lang=javascript
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (63.03%)
 * Likes:    229
 * Dislikes: 0
 * Total Accepted:    30.8K
 * Total Submissions: 48.8K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 
 * 输入:
 * 
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 * 
 * 输出: ["1->2->5", "1->3"]
 * 
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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
 * @return {string[]}
 */
var binaryTreePaths = function(root) {
    let res = [];
    let path = [];
    let dfs = (root) => {
        if (root == null) return;

        path.push(root);
        dfs(root.left);
        dfs(root.right);

        if (!root.left && !root.right) {
            res.push(path.map(v => v.val).join('->')); 
        }

        path.pop();
    }

    dfs(root);

    return res;
};

var binaryTreePaths = function(root) {
    if (root == null) return [];

    let stack = [];
    let p = root;
    let res = [];
    let set = new Set();
    while(stack.length || p) {
        while(p) {
            stack.push(p);
            p = p.left;
        }

        let node = stack[stack.length - 1];

        if (!node.left && !node.right) {
            res.push(stack.map(v => v.val).join('->'));
        }

        if (node.right && !set.has(node.right)) {
            p = node.right;
            set.add(node.right);
        }else{
            stack.pop();
        }
    }

    return res;
}
// @lc code=end

