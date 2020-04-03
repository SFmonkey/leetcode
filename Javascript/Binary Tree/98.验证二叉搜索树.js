/*
 * @lc app=leetcode.cn id=98 lang=javascript
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (29.56%)
 * Likes:    482
 * Dislikes: 0
 * Total Accepted:    86.3K
 * Total Submissions: 291.7K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 * 
 * 假设一个二叉搜索树具有如下特征：
 * 
 * 
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
 * 
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
var isValidBST = function(root) {
    let queue = [];
    let dfs = (root) => {
        if (root == null) return;
        dfs(root.left);
        queue.push(root.val);
        dfs(root.right);
    }

    dfs(root);

    for (let i =0; i< queue.length-1; i++){
        if (queue[i] >= queue[i+1]) return false
    }

    return true;
};

var isValidBST = function(root) {
    const help = (node, max, min) => {
        if(node == null) return true;
        if(node.val >= max || node.val <= min) return false;
        // 左孩子更新上界，右孩子更新下界，相当于边界要求越来越苛刻
        return help(node.left, node.val, min)
                && help(node.right, max, node.val);
    }
    return help(root, Number.MAX_SAFE_INTEGER, Number.MIN_SAFE_INTEGER);
};

var isValidBST = function(root) {
    if (root == null) return true;
    let stack = [root];
    let min = Number.MIN_SAFE_INTEGER;
    let max = Number.MAX_SAFE_INTEGER;

    root.max = max; root.min = min;

    while(stack.length) {
        let node = stack.pop();
        if(node.val <= node.min || node.val >= node.max)
        return false;

        if(node.left) {
            stack.push(node.left);
            // 更新上下界
            node.left.max = node.val;
            node.left.min = node.min;
        }
        if(node.right) {
            stack.push(node.right);
            // 更新上下界
            node.right.max = node.max;
            node.right.min = node.val;
        }
    }

    return true;
}
// @lc code=end

