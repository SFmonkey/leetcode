/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {number[]} preorder
 * @param {number[]} inorder
 * @return {TreeNode}
 */
/*
 * 前序排序中第一节点就是根节点
 * 在中序排序中找到根节点的位置，左边为他的左子树集合，右边为右子树集合
 * 中序排序中当前根节点在当前集合的位置(index) "+1"  即当前子树的集合长度
 * 当前子树集合分布应该为 0 根节点   1---index+1 左子树  index+1 --- length-1  为右子树
 * 所以左子树为前序排序为preorder.slice(1, index+1), 中序排序为inorder.slice(0, index)
 * 右子树前序排序为preorder.slice(index +1)， 中序排序为 inorder.slice(index+1))
*/
var buildTree = function(preorder, inorder) {
    if(!preorder.length) {
        return null
    }
    let index = inorder.indexOf(preorder[0]);
    let tree = new TreeNode(preorder[0]);
    tree.left = buildTree(preorder.slice(1, index+1),inorder.slice(0, index))   
    tree.right = buildTree(preorder.slice(index +1),inorder.slice(index+1))
    return tree
};