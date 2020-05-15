/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
/*
 * 前缀和思路
 * 将数组前n次和 通过 {sum: count} 方式存储  =》 sum 为前n次和，count为sum出现的次数
 * 将问题转化为 求sum - k 出现的次数
*/
var subarraySum = function(nums, k) {
    let obj = {0: 1}
    let sum = 0
    let count = 0
    for(let i=0;i<nums.length;i++) {
        sum+= nums[i]
        if(obj[sum-k]) {
            count+= obj[sum-k]
        }
        if(obj[sum]) {
            obj[sum] ++ 
        }else{
            obj[sum] = 1
        }
    }
    return count
};