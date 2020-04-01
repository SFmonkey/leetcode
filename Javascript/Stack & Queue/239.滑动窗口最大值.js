/*
 * @lc app=leetcode.cn id=239 lang=javascript
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var maxSlidingWindow = function(nums, k) {
    if (nums.length == 0 || !k) return [];
    let window = [], res = [];
    for(let i = 0; i < nums.length; i ++) {
        // 看最大值的 index 是否还在window内
        if (window[0] !== undefined && window[0] <= i - k) window.shift()

        while(nums[window[window.length - 1]] <= nums[i]) window.pop();

        window.push(i);

        if (i >= k - 1) res.push(nums[window[0]])
    }

    return res;
};
// @lc code=end

