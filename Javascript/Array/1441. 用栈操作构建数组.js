/**
 * @param {number[]} target
 * @param {number} n
 * @return {string[]}
 */
/**
 * 在1-n中寻找target中的数字 i，当找到所有target的中的所有元素则返回
 * 当i在target中，则需要进行Push操作
 * 当i没有在target中则需要进行Push Pop操作
 */
var buildArray = function(target, n) {
    let res = []
    let index = 0
    for(let i=0;i<n;i++) {
        if(index == target.length) {
            return res
        }
        if(target[index] == i+1) {
            res.push("Push")
            index++
        }else{
            res.push("Push", "Pop")
        }
    }
    return res
};