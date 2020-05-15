/*
 * @lc app=leetcode.cn id=14 lang=javascript
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (36.72%)
 * Likes:    974
 * Dislikes: 0
 * Total Accepted:    235K
 * Total Submissions: 634.7K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 * 
 * 如果不存在公共前缀，返回空字符串 ""。
 * 
 * 示例 1:
 * 
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 * 
 * 
 * 示例 2:
 * 
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 * 
 * 
 * 说明:
 * 
 * 所有输入只包含小写字母 a-z 。
 * 
 */

// @lc code=start
/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    if (strs.length === 0 ) return "";

    let prefix = strs[0];
    for(let i = 1; i < strs.length; i++) {
        let j = 0
        for(; j < strs[i].length && j < prefix.length; j++) {
            if (strs[i][j] !== prefix[j]) {
                break;
            }
        }

        prefix = prefix.substring(0, j);

        if (prefix === '') {
            return '';
        }
    }

    return prefix;
};
// @lc code=end

