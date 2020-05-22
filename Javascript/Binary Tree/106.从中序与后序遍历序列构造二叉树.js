/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {number[]} inorder
 * @param {number[]} postorder
 * @return {TreeNode}
 */
/*
 * 后序排序中最后一个节点即为根节点
 * 在中序排序中找到根节点的位置，左边为他的左子树集合，右边为右子树集合
 * 中序排序中当前根节点在当前集合的位置(index)  即当前节点在中序排序中子树集合的右边界的下一个  (slice)方法不包括下一个
 * 当前子树集合分布应该为 length-1 根节点   0---index 左子树  index+1 --- length-1  为右子树
 * 所以左子树为后序排序为postorder.slice(0, index), 中序排序为inorder.slice(0, index)
 * 右子树后序排序为postorder.slice(index, postorder.length - 1)， 中序排序为 inorder.slice(index+1)
*/
var buildTree = function(inorder, postorder) {
    if(!inorder.length)  return null
    let val = postorder[postorder.length -1]
    let index = inorder.indexOf(val)
    let root = new TreeNode(val)
    root.right =buildTree(inorder.slice(index+1), postorder.slice(index, postorder.length - 1))
    root.left = buildTree(inorder.slice(0, index), postorder.slice(0, index))
    return root 
};