/*
 * @lc app=leetcode.cn id=56 lang=javascript
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (40.81%)
 * Likes:    428
 * Dislikes: 0
 * Total Accepted:    98.4K
 * Total Submissions: 230.8K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 * 
 * 示例 1:
 * 
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 * 
 * 
 * 示例 2:
 * 
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 * 
 */

// @lc code=start
/**
 * @param {number[][]} intervals
 * @return {number[][]}
 */
var merge = function(intervals) {
    if (!intervals.length) {
        return [];
    }

    let res = [];
        
    intervals.sort((a, b) => {
      return a[0] - b[0];
    });  
    
    let idx = intervals[0];
    
    for (let i = 1; i < intervals.length; i++) {
      let cur = intervals[i]
      if (idx[1] >= cur[0]) {
        idx[1] = Math.max(idx[1], cur[1]); 
      } else {
        res.push(idx);
        idx = cur;
      }
    }
    
    res.push(idx);
    return res;
};
// @lc code=end

