/*
 * @lc app=leetcode.cn id=225 lang=javascript
 *
 * [225] 用队列实现栈
 */

// @lc code=start
/**
 * Initialize your data structure here.
 */
var MyStack = function() {
    this.queue1 = [];
    this.queue2 = [];
};

/**
 * Push element x onto stack. 
 * @param {number} x
 * @return {void}
 */
MyStack.prototype.push = function(x) {
    this.queue1.push(x);
};

/**
 * Removes the element on top of the stack and returns that element.
 * @return {number}
 */
MyStack.prototype.pop = function() {
    if (this.queue1.length !== 1) this.transfrom();

    return this.queue1.shift();
};

/**
 * Get the top element.
 * @return {number}
 */
MyStack.prototype.top = function() {
    if (this.queue1.length !== 1) this.transfrom();

    return this.queue1[0];
};

/**
 * Returns whether the stack is empty.
 * @return {boolean}
 */
MyStack.prototype.empty = function() {
    return !this.queue1.length && !this.queue2.length
};

MyStack.prototype.transfrom = function() {
    while (this.queue1.length > 1) {
        this.queue2.push(this.queue1.shift());
    }

    if (!this.queue1.length && this.queue2.length) {
        while (this.queue2.length > 1) {
            this.queue1.push(this.queue2.shift());
        }

        let temp = this.queue1;
        this.queue1 = this.queue2;
        this.queue2 = temp;
    }
};
/**
 * Your MyStack object will be instantiated and called as such:
 * var obj = new MyStack()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.empty()
 */
// @lc code=end

