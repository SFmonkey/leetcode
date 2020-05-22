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
 * 给每个节点添加一个标志位pos
 * 如果某个节点已经有pos标志位了  说明有环  返回该节点
 * 如果循环自己跳出，则返回null
*/
var detectCycle = function(head) {
    let p = head
    let i = 0
    while(p) {
        if(p.pos == undefined) {
            p.pos = i++
            p = p.next
        }else{
            return p
        }
    }   
    return null
};