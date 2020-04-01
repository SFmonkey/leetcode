/*
 * @lc app=leetcode.cn id=234 lang=javascript
 *
 * [234] 回文链表
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
 * @return {boolean}
 */
var isPalindrome = function(head) {
    if (!head) return true;
    let reverse = (pre, cur) => {
        if (!cur) return pre;
        const next = cur.next;
        cur.next = pre;
        pre = cur;
        cur = next;
        return reverse(pre, cur);
    }
    let fast = head.next;
    let slow = head;
    while(fast && fast.next) {
        fast = fast.next.next;
        slow = slow.next;
    }

    let next = slow.next;
    slow.next = null;

    let newStart = reverse(null, next);

    for(let p = head, newP = newStart; newP != null; p = p.next, newP = newP.next) {
        if (p.val != newP.val) return false;
    }

    return true;
};
// @lc code=end

