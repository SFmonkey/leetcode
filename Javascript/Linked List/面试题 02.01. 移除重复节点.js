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
 * 用q表示p的父节点
 * 如果该节点已经存在则将p = p.next q.next = p
*/ 
var removeDuplicateNodes = function(head) {
    let q;
    let p = head;
    let obj = {}
    while(p){
        if(!obj[p.val]) {
            obj[p.val] = true
            p=p.next
            if(!q) {
                q= head
            }else{
                q = q.next
            }
        }else{
            p=p.next
            q.next = p
        } 
    }
    return head
};