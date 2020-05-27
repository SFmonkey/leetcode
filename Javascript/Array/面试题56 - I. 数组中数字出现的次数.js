/**
 * @param {number[]} nums
 * @return {number[]}
 */
/**
 * 在obj中存储出现过一次的元素
 */
var singleNumbers = function(nums) {
    let obj = {}
    for(let i=0;i<nums.length;i++) {
        if(!obj[nums[i]]) {
            obj[nums[i]] = true
        }else{
            delete obj[nums[i]]
        }
    }
    let res = []
    for(let item in obj) {
        res.push(item)
    }
    return res
};