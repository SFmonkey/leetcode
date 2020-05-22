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
/**
 * 如果该节点为叶子节点统计该叶子节点的值
 * 如果非叶子节点需要统计其下面所有子树的值
 * 统计该值时判断该值出现的次数，若大于当前最大值则重新统计
 */
var findFrequentTreeSum = function(root) {
    let obj = {}
    let res = []
    let max = -Infinity
    var isInMax = function(val) {
        if(!obj[val]) {
            obj[val] = 1
        }else{
            obj[val]++
        }
        if(max < obj[val]) {
            max = obj[val]
            res = [val]
        }else if(max == obj[val]) {
            res.push(val)
        }
    }
    var getNum = function(tree) {
        if(!tree) {
            return null
        }
        if(!tree.left && !tree.right) {     //统计叶子节点
            isInMax(tree.val)
            return tree.val
        }
        let left = getNum(tree.left)
        let right = getNum(tree.right)
        let sum = right + left + tree.val   
        isInMax(sum)
        return sum
    }
    getNum(root)
    return res
};