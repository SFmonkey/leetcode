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
/**
 * 如果为叶子节点返回为1
 * 如果当前节点不存在返回0 则表明当前节点为有一个孩子节点，所以直接取孩子节点的深度
*/
var minDepth = function(root) {
    if(!root) {
        return 0
    }
    if(!root.left && !root.right) {
        return 1
    }
    let left = minDepth(root.left)
    let right = minDepth(root.right)
    return (Math.min(left == 0 ? Infinity : left, right == 0 ? Infinity : right) + 1)
};