/*
 * @lc app=leetcode.cn id=24 lang=javascript
 *
 * [24] 两两交换链表中的节点
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
// var swapPairs = function(head) {
//     if (head == null || head.next == null) return head; 
//     let dummyHead = p = new ListNode();
//     let node1, node2;

//     dummyHead.next = head;

//     while((node1 = p.next) && (node2 = p.next.next)) {
//         node1.next = node2.next;
//         node2.next = node1;
//         p.next = node2;
//         p = node1;
//     }

//     return dummyHead.next;
// };
var swapPairs = function(head) {
    if (!head || !head.next) return head;

    let next = head.next;

    head.next = swapPairs(next.next);
    next.next = head;

    return next;
}
// @lc code=end

