/*
 * @lc app=leetcode.cn id=120 lang=javascript
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode-cn.com/problems/triangle/description/
 *
 * algorithms
 * Medium (64.15%)
 * Likes:    352
 * Dislikes: 0
 * Total Accepted:    49.4K
 * Total Submissions: 76.9K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
 * 
 * 例如，给定三角形：
 * 
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 * 
 * 
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 * 
 * 说明：
 * 
 * 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
 * 
 */

// @lc code=start
/**
 * @param {number[][]} triangle
 * @return {number}
 */
var minimumTotal = function(triangle) {
    if (triangle.length < 1) return 0;

    if (triangle.length == 1) return triangle[0][0];

    let dp = [];

    for(let i = 0; i < triangle.length; i++) {
        dp[i] = new Array(triangle[i].length);
    }

    dp[0][0] = triangle[0][0];
    dp[1][1] = triangle[1][1] + triangle[0][0];
    dp[1][0] = triangle[1][0] + triangle[0][0];

    for(let i = 2; i < triangle.length; i++) {
        for (let j = 0; j < triangle[i].length; j ++) {
            if (j == 0) {
                dp[i][j] = dp[i-1][j] + triangle[i][j];
            } else if (j == triangle[i].length - 1) {
                dp[i][j] = dp[i-1][j-1] + triangle[i][j];
            } else {
                dp[i][j] = Math.min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j];
            }
        }
    }

    return Math.min.apply(null, dp[dp.length - 1]);
};
// @lc code=end

