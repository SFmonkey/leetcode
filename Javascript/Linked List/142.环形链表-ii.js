/*
 * @lc app=leetcode.cn id=142 lang=javascript
 *
 * [142] 环形链表 II
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
var detectCycle = function(head) {
    let dummyHead = fast = slow = new ListNode();

    dummyHead.next = head;
    if (fast.next == null ||  fast.next.next == null) return null;
    while(fast.next && fast.next.next) {
        fast = fast.next.next;
        slow = slow.next;
        if (fast == slow) {
            let p = dummyHead;
            while (p != slow) {
                p = p.next;
                slow = slow.next;
            }

            return p;
        }
    }

    return null;
};
// @lc code=end

