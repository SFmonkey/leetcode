/*
 * @lc app=leetcode.cn id=206 lang=javascript
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function(head) {
    let reverse = function(pre, cur){
        if(!cur) return pre;
        // 保存 next 节点
        let next = cur.next;
        cur.next = pre;
        return reverse(cur, next);
      }
      return reverse(null, head);
};
// @lc code=end

