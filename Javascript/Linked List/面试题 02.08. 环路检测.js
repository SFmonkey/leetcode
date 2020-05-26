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
/**
 * 快慢指针方法
 * 以数组值为指针将数组类比为链表
 * fast 每次走两步，slow每次走一步
 * 当第一次fast == slow时，将快指针置为0
 * 快指针和慢指针每次移动一位当再次相遇时即为重复值
 * 
 * 注：第一次相遇时，设环长为m， 非环部分为n, 相遇在k处， 则快指针走了 n+B*m+k   慢指针走了 n+A*m+k
 * 快指针 - 慢指针 = (B-A) * m
 * 所以快慢指针都走了环长的倍数，所以当fast到起始点时走非环部分n时，慢指针也就刚好走到环的重复位置
 * 即此时快慢指针指向的即为环的起始值
 */
var detectCycle = function(head) {
    let slow = head;
    let fast = head;
    while(fast && fast.next && fast.next.next) {
        fast = fast.next.next;
        slow = slow.next;
        if(slow == fast) {
            fast = head
            while(slow !== fast){
                fast = fast.next;
                slow = slow.next;
            }
            return fast
        }
    }  
    return null
};