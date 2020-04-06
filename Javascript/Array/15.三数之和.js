/*
 * @lc app=leetcode.cn id=15 lang=javascript
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (26.23%)
 * Likes:    1967
 * Dislikes: 0
 * Total Accepted:    191.2K
 * Total Submissions: 726.8K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？请你找出所有满足条件且不重复的三元组。
 * 
 * 注意：答案中不可以包含重复的三元组。
 * 
 * 
 * 
 * 示例：
 * 
 * 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 * 
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 * 
 * 
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function(nums) {
    let res = [];
    nums.sort((a, b) => a - b);

    for (let i = 0; i < nums.length; i ++) {
        let target = 0 - nums[i];

        let l = i + 1;
        let r = nums.length - 1;

        if (nums[i] > 0) {
            break;
        }

        if (i == 0 || nums[i] != nums[i - 1]) {
            while(l < r) {
                if (nums[l] + nums[r] == target) {
                    res.push([nums[i], nums[l], nums[r]]);
                    while (l < r && nums[l] == nums[l + 1]) l++;
                    while (l < r && nums[r] == nums[r - 1]) r--;
                    l++;
                    r--;
                } else if (nums[l] + nums[r] < target) {
                    l ++;
                } else {
                    r --;
                }
            }
        }
    }

    return res;
};
// @lc code=end

