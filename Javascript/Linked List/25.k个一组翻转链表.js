/*
 * @lc app=leetcode.cn id=25 lang=javascript
 *
 * [25] k个一组翻转链表
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
 * @param {number} k
 * @return {ListNode}
 */
// var reverseKGroup = function(head, k) {
//     if (!head) return null;

//     let reverse = (a, b) => {
//         let pre;
//         let cur = a;
//         let next = b;
//         while(cur != b) {
//             next = cur.next;
//             cur.next = pre;
//             pre = cur;
//             cur = next;
//         }

//         return pre;
//     }

//     let a = head;
//     let b = head;

//     for(let i = 0; i < k; i ++) {
//         if (!b) return b;
//         b = b.next;
//     }

//     let newHead = reverse(a, b);
//     a.next = reverseKGroup(b, k);

//     return newHead;
// };
var reverseKGroup = function(head, k) {
    let count = 0;

    for (let p = head; p != null; p = p.next) {
        count ++;
    }

    if (count < k ) return head;

    let loopCount = Math.floor(count / k);
    let p = dummyHead = new ListNode();
    
    dummyHead.next = head;

    for (let i = 0; i < loopCount; i++) {
        let pre = null,cur = p.next;

        for (let j = 0; j < k; j++) {
            let next = cur.next;
            cur.next = pre;
            pre = cur;
            cur = next;
        }

        let start = p.next;

        p.next = pre;
        start.next = cur;
        p = start;
    }

    return dummyHead.next;


}
// @lc code=end

