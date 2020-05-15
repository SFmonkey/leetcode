/**
 * initialize your data structure here.
 */
var MinStack = function() {
    this.arr = []
    this.minArr = []
};

/** 
 * @param {number} x
 * @return {void}
 */
MinStack.prototype.push = function(x) {
    this.arr.push(x)
    if(x < this.minArr[this.minArr.length-1] || this.minArr.length == 0) {
        this.minArr.push(x)
    }else{
        this.minArr.push(this.minArr[this.minArr.length-1])
    }
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function() {
    let temp = this.arr.splice(this.arr.length-1, 1)
    this.minArr.splice(this.minArr.length-1, 1)
    return temp
};

/**
 * @return {number}
 */
MinStack.prototype.top = function() {
    return this.arr[this.arr.length-1]
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function() {
    return this.minArr[this.minArr.length-1]
};

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(x)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */