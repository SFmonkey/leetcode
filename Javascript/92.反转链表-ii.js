/*
 * @lc app=leetcode.cn id=92 lang=javascript
 *
 * [92] 反转链表 II
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
 * @param {number} m
 * @param {number} n
 * @return {ListNode}
 */
// var reverseBetween = function(head, m, n) {
//     let p = dummyHead = new ListNode();
//     let front,tail, pre, cur;
//     p.next = head;

//     for(let i = 0; i < m - 1; i ++) {
//         p = p.next;
//     }

//     front = p;

//     pre = tail = p.next;
//     cur = pre.next;

//     for(let i = 0; i < n - m; i ++) {
//         let next = cur.next;
//         cur.next = pre;
//         pre = cur;
//         cur = next;
//     }
    
//     front.next = pre;
//     tail.next = cur;

//     return dummyHead.next;
// };
var reverseBetween = function(head, m, n) {
    let nextTail = null;
    let reverseN = (head, n) => {
        if (n == 1) {
            nextTail = head.next;
            return head;
        }
        let last = reverseN(head.next, n-1);
        head.next.next = head;
        head.next = nextTail;
        return last;
    }
    if (m == 1) {
        return reverseN(head, n);
    }
    head.next = reverseBetween(head.next, m -1, n-1);
    return head;
}
// @lc code=end

