/**
 * 使用stacks 存储当前key和value 还有他出现的次数  times
 * 使用keys存储次数相同的集合  即一个键值对存入，将会保留到 keys[0]里面 当执行一次get操作时 他就会被放入 keys[1]中
 * 
 * 执行get时判断stack中是否含有该值如果有则返回，并且将他在keys 中的集合向下挪一位
 * 如果没有返回负一
 * 
 * 执行put时，如果该值在stacks中存储过 直接将其重新赋值，并且将他在keys 中的集合向下挪一位
 * 如果没有
 *      是否等于最大缓存capacity，
 *          如果等于移除出现次数最少并且最久未使用的
 *      在keys[0]中push该值，stacks中赋值
 */
/**
 * @param {number} capacity
 */
var LFUCache = function(capacity) {
    this.keys = {0: []}
    this.stacks = {}
    this.capacity = capacity
};

/** 
 * @param {number} key
 * @return {number}
 */
LFUCache.prototype.get = function(key) {
    if(this.stacks[key] == undefined) {
        return -1
    }else{
        this.changeIndex(key)
        return this.stacks[key].value
    }
};

/** 
 * @param {number} key 
 * @param {number} value
 * @return {void}
 */
LFUCache.prototype.put = function(key, value) {
    if(this.capacity == 0) {
        return
    }
    if(this.stacks[key] == undefined) {
		if(Object.keys(this.stacks).length == this.capacity) {
			this.remove()
		}
        this.keys[0].push(key)
        this.stacks[key] = {value: value, times:0}
    }else{
        this.stacks[key].value = value
        this.changeIndex(key)
    }
};

LFUCache.prototype.changeIndex = function(key) {
    let times = ++this.stacks[key].times
    let index = this.keys[times -1].indexOf(key)
    let temp = this.keys[times -1].splice(index, 1)
    if(!this.keys[times]) {
        this.keys[times] = []
    }
    this.keys[times].push(key)
}
LFUCache.prototype.remove = function() {
	for(let item in this.keys) {
		if(this.keys[item].length > 0) {
			let temp = this.keys[item].shift()
			delete this.stacks[temp]
			return
		}
	}	
}
/**
 * Your LFUCache object will be instantiated and called as such:
 * var obj = new LFUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */