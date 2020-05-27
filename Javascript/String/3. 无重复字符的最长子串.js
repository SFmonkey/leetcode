/**
 * @param {string} s
 * @return {number}
 */
/**
 * 滑动窗口
 * 当s[right]没有在strObj中或者strObj[s[right]] < left 当前strObj中存储的重复元素不在当前计算的字符串中，计算则现在为当前的最大长度
 * 如果出现过将left置为上次改重复元素的上一位
 * 
 * 注意：left指向的是当前需要计算的字符串的前一位
 */
var lengthOfLongestSubstring = function(s) {
    let left = -1;
    let right = 0;
    let res = 0
    let strObj = {}
    for(;left < s.length && right < s.length;) {
        if(strObj[s[right]] == undefined || strObj[s[right]] < left) {
            res = Math.max(res,  right-left)
        }else{
            left = strObj[s[right]]
        }
        strObj[s[right]] = right
        right++
    }
    return res
};