/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
/**
 * 滑动窗口问题
 * 当连续n个数字和小于k时，将右指针向右移一位，left到right的数字相乘都小于k，所以cout += right-left+1
 * 当大于k时mul除掉num[left],左指针向右移一位
 */
var numSubarrayProductLessThanK = function(nums, k) {
    let left = right = 0;
    let mul = 1;
    let count = 0
    for(;left < nums.length && right < nums.length;) {
        let temp = mul * nums[right]
        if(temp < k) {
            count+= right-left+1
            mul = temp
            right++
        }else{
            mul/=nums[left]
            left++
        }
    }
    return count
};