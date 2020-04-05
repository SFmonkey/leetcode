/*
 * @lc app=leetcode.cn id=1 lang=javascript
 *
 * [1] 两数之和
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (45.74%)
 * Total Accepted:    334.3K
 * Total Submissions: 730.9K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 * 
 * 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
 * 
 * 示例:
 * 
 * 给定 nums = [2, 7, 11, 15], target = 9
 * 
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 * 
 * 
 */
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    for(let i = 0; i < nums.length; i ++) {
        const idx = nums.indexOf(target - nums[i])
        if (idx > -1 && idx !== i) {
            return [i, idx]
        }
    }

    return [];
};

var twoSum = function(nums, target) {
    let map = new Map();

    for (let i = 0; i < nums.length; i++) {
        const val = nums[i];
        
        map.set(val, i);
    }

    for (let i = 0; i < nums.length; i++) {
        const val = nums[i];
        
        if(map.has(target - val) && map.get(target - val) !== i) {
            return [i, map.get(target - val)]
        }
    }

    return [];
}
