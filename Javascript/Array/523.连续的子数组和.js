/*
 * @lc app=leetcode.cn id=523 lang=javascript
 *
 * [523] 连续的子数组和
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var checkSubarraySum = function(nums, k) {
    let sum = 0;
    let obj = {0: -1}
    k = Math.abs(k)
    for(let i=0;i<nums.length;i++) {
        sum += nums[i]
        if(k != 0) {
            sum = sum % k
        }
        if(obj[sum] !== undefined ) {
            if((i - obj[sum]) >1) {
                return true
            }
        }else{
            obj[sum] = i
        }
    }
    return false
};
// @lc code=end

