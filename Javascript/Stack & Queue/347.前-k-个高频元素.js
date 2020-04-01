/*
 * @lc app=leetcode.cn id=347 lang=javascript
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
class Heap {
    constructor() {
        this.data = [];
    }

    insert(item) {
        this.data.push(item);
        let idx = this.data.length - 1;
        let fatherIdx = Math.floor((idx - 1) / 2);

        while(fatherIdx >= 0) {
            if (this.data[idx].val > this.data[fatherIdx].val) {
                let temp = this.data[idx];
                this.data[idx] = this.data[fatherIdx];
                this.data[fatherIdx] = temp;
            }

            idx =fatherIdx;
            fatherIdx = Math.floor((idx -1 ) / 2 );
        }
    }

    deleting() {
        if (this.data.length == 1) {
            return this.data.pop();
        }

        let idx = 0;
        let val = this.data[idx];
        this.data[idx] = this.data.pop();

        while(idx < this.data.length) {
            let left = 2 * idx + 1;
            let right = 2 * idx + 2;
            let select = left;

            if (right < this.data.length) {
                select = (this.data[left].val < this.data[right].val) ? right : left;
            }

            if (select < this.data.length && this.data[idx].val < this.data[select].val) {
                let temp = this.data[idx];
                this.data[idx] = this.data[select];
                this.data[select] = temp;
            }
            idx = select;
        }
        return val;
    }
}

const topKFrequent=(nums,k)=>{
    let map=new Map(),res=[];
    for(let i=0;i<nums.length;i++){
        if(map.has(nums[i])){
            map.set(nums[i],map.get(nums[i])+1);
        }else{
            map.set(nums[i],1);
        }
    }
    let h=new Heap();
    for(let [key,val] of map.entries()){
        h.insert({key,val});
    }
    while(res.length<k){
        res.push(h.deleting().key);
    }
    return res;
};

// @lc code=end

