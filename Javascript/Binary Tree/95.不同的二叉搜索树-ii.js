/*
 * @lc app=leetcode.cn id=95 lang=javascript
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (62.33%)
 * Likes:    338
 * Dislikes: 0
 * Total Accepted:    25.7K
 * Total Submissions: 41.3K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
 * 
 * 示例:
 * 
 * 输入: 3
 * 输出:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * 解释:
 * 以上的输出对应以下 5 种不同结构的二叉搜索树：
 * 
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
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
 * @param {number} n
 * @return {TreeNode[]}
 */
var generateTrees = function(n) {
    let fn = (start, end) => {
        if (start > end) return [null];
        if (start === end) return [new TreeNode(start)];

        let res = [];

        for(let i = start; i <= end; i++) {
            let leftNodes = fn(start, i - 1);
            let rightNodes = fn(i + 1, end);

            for(let j = 0; j < leftNodes.length; j++) {
                for(let k = 0; k < rightNodes.length; k++) {
                    let root = new TreeNode(i);

                    root.left = leftNodes[j];
                    root.right = rightNodes[k];

                    res.push(root);
                }
            }
        }

        return res;
    }

    if (n == 0) return [];
    return fn(1, n);
};

var generateTrees = function(n) {
    let clone = (node, offset) => {
        if(node == null) return null;
        let newnode = new TreeNode(node.val + offset);
        newnode.left = clone(node.left, offset);
        newnode.right = clone(node.right, offset);
        return newnode;
    }

    if (n == 0) return [];
    let dp = [];
    dp[0] = [null];

    for(let i = 1; i <= n; i++) {
        dp[i] = [];
        for(let j = 1; j <= i; j++) {
            for(let leftNode of dp[j - 1]) {
                for(let rightNode of dp[i - j]) {
                    let node = new TreeNode(j);

                    node.left = leftNode;
                    node.right = clone(rightNode, j);
                    dp[i].push(node);
                }
            }
        }
    }

    return dp[n];
}
// @lc code=end

