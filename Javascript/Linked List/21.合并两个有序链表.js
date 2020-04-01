/*
 * @lc app=leetcode.cn id=21 lang=javascript
 *
 * [21] 合并两个有序链表
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
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
// var mergeTwoLists = function(l1, l2) {
//     const merge = (l1, l2) => {
//         if (l1 == null) return l2;
//         if (l2 == null) return l1;

//         if (l1.val > l2.val) {
//             l2.next = merge(l1, l2.next);
//             return l2;
//         }else {
//             l1.next = merge(l1.next, l2);
//             return l1;
//         }

//     }

//     return merge(l1, l2);
// };
var mergeTwoLists = function(l1, l2) {
    if (l1 == null) return l2;
    if (l2 == null) return l1;

    let p = dummyHead = new ListNode();

    let p1 = l1, p2 =l2;

    while(p1 && p2) {
        if (p1.val > p2.val) {
            p.next = p2;
            p = p.next;
            p2 = p2.next;
        } else {
            p.next = p1;
            p = p.next;
            p1 = p1.next;
        }
    }

    if (p1) {
        p.next = p1;
    } else {
        p.next = p2;
    }

    return dummyHead.next;
}
// @lc code=end

