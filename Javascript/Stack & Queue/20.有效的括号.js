/*
 * @lc app=leetcode.cn id=20 lang=javascript
 *
 * [20] 有效的括号
 */

// @lc code=start
/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    let stack = [];
    for(let i = 0; i < s.length; i++) {
        let ch = s.charAt(i);
        if (ch === '(' || ch === '{' || ch === '[') stack.push(ch);



        if (ch === ')' && stack.pop(ch) !== '(') return false;
        if (ch === '}' && stack.pop(ch) !== '{') return false;
        if (ch === ']' && stack.pop(ch) !== '[') return false;
    }

    return stack.length === 0;
};
// @lc code=end

