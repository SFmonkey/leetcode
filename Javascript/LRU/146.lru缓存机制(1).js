/*
 * @lc app=leetcode.cn id=146 lang=javascript
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (46.59%)
 * Likes:    515
 * Dislikes: 0
 * Total Accepted:    47.7K
 * Total Submissions: 101.4K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 * 
 * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果密钥已经存在，则变更其数据值；如果密钥不存在，则插入该组「密钥/数据值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 * 
 * 
 * 
 * 进阶:
 * 
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 * 
 * 
 * 
 * 示例:
 * 
 * LRUCache cache = new LRUCache( 2  );
 * 
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得密钥 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得密钥 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 * 
 * 
 */

// @lc code=start
/**
 * 使用obj来存储存入的键值对，arr来存储访问顺序(最久未被访问的元素在第一)， capacity 缓存容量
 * 进行get请求时，如果没有找到该值则返回-1
 * 若找到该值则将该值移到数组的最后一个元素，被返回该值
 * 进行put操作时，如果当前值已经存在则将该值重新赋值，并将其移到数组的最后
 * 若没有该值，比较当前数组的大小是否等于capacity缓存容量
 *      若等于则将数组中第一个元素(即最久未被访问的元素)移除，同时删除obj里面的key
 *      若不等于则直接往数组里面进行push操作(即该元素是最近访问), obj[key] = value
 */
/**
 * @param {number} capacity
 */
var LRUCache = function(capacity) {
    this.obj = {}
    this.arr = []
    this.capacity = capacity
};

/** 
 * @param {number} key
 * @return {number}
 */
LRUCache.prototype.get = function(key) {
    if(this.obj[key] != undefined) {
        this.changeIndex(key)
        return this.obj[key]
    }else{
        return -1
    }
};

/** 
 * @param {number} key 
 * @param {number} value
 * @return {void}
 */
LRUCache.prototype.put = function(key, value) {
    if(this.obj[key] != undefined) {
        this.obj[key] = value
        this.changeIndex(key)
    }else{
        if(this.arr.length == this.capacity) {
            let temp = this.arr.shift()
            delete this.obj[temp]
        }
        this.obj[key] = value
        this.arr.push(key)
    }
};
/**
 * @param {number} key 
 * 该方法将将key移动到数组最后一位(即该元素为最后被访问)
 */
LRUCache.prototype.changeIndex = function(key) {
    let index = this.arr.indexOf(key)
    let temp = this.arr.splice(index, 1)
    this.arr.push(temp[0])
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */
// @lc code=end