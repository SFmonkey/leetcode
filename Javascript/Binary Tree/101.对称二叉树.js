/*
 * @lc app=leetcode.cn id=101 lang=javascript
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (50.60%)
 * Likes:    693
 * Dislikes: 0
 * Total Accepted:    116.4K
 * Total Submissions: 229.9K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 * 
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 * 
 * 
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 * 
 * 
 * 说明:
 * 
 * 如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
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
 * @return {boolean}
 */
var isSymmetric = function(root) {
    if (!root) return true;
    let queue = [root.left, root.right];
    let node1, node2;

    while(queue.length) {
        node1 = queue.shift();
        node2 = queue.shift();

        if (!node1 == !node2) continue;

        if (!node1 || !node2 || node1.val !== node2.val)  return false;

        queue.push(node1.left);
        queue.push(node2.right);
        queue.push(node1.right);
        queue.push(node2.left);
    }

    return true;
};

var isSymmetric = function(root) {
    let fn = (node1, node2) => {
        if (!node1 && !node2) return true;

        if(!node1 || !node2 || node1.val !== node2.val) return false;

        return fn(node1.left, node2.right) && fn(node2.left, node1.right);
    }
    if (root == null) return true;
    return fn(root.left, root.right);
}
// @lc code=end

