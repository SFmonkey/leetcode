/*
 * @lc app=leetcode.cn id=341 lang=javascript
 *
 * [341] 扁平化嵌套列表迭代器
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * function NestedInteger() {
 *
 *     Return true if this NestedInteger holds a single integer, rather than a nested list.
 *     @return {boolean}
 *     this.isInteger = function() {
 *         ...
 *     };
 *
 *     Return the single integer that this NestedInteger holds, if it holds a single integer
 *     Return null if this NestedInteger holds a nested list
 *     @return {integer}
 *     this.getInteger = function() {
 *         ...
 *     };
 *
 *     Return the nested list that this NestedInteger holds, if it holds a nested list
 *     Return null if this NestedInteger holds a single integer
 *     @return {NestedInteger[]}
 *     this.getList = function() {
 *         ...
 *     };
 * };
 */
/**
 * @constructor
 * @param {NestedInteger[]} nestedList
 */
// var NestedIterator = function(nestedList) {
//     this.list = [];
//     this.getList(nestedList);
// };


// NestedIterator.prototype.getList = function(nestedList) {
//     for (let i = 0; i < nestedList.length; i++) {
//         if (Object.prototype.toString.call(nestedList[i]) === '[object Array]') {
//             this.getList(nestedList[i]);
//         } else {
//             this.list.push(nestedList[i]);
//         } 
//     }
// };

// /**
//  * @this NestedIterator
//  * @returns {boolean}
//  */
// NestedIterator.prototype.hasNext = function() {
//     return this.list.length > 0;
// };

// /**
//  * @this NestedIterator
//  * @returns {integer}
//  */
// NestedIterator.prototype.next = function() {
//     return this.list.shift();
// };

/**
 * Your NestedIterator will be called like this:
 * var i = new NestedIterator(nestedList), a = [];
 * while (i.hasNext()) a.push(i.next());
*/
var NestedIterator = function(nestedList) {
    this.list = [];
    this.resetList(nestedList);
}; 
NestedIterator.prototype.resetList = function(arr) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i].isInteger())
            this.list.push(arr[i].getInteger())
        else
            this.resetList(arr[i].getList())
    }
}; 
NestedIterator.prototype.hasNext = function() {
    return this.list.length > 0;
}; 
NestedIterator.prototype.next = function() {
    return this.list.shift();
};
// @lc code=end

