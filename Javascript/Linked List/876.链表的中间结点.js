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
/*
 * 使用两个指针操作链表，p 单步移动  q双步移动
 * 当q到达链表末端p在中间
*/
var middleNode = function(head) {
    let p= head;
    let q = head;
    while(q != null) {
        if(!q.next) {
            return p
        }
        p=p.next
        q=q.next.next;
    }
    return p
};