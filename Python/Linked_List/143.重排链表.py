# Definition for single-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverse(self, node):
        if node.next == None:
            return node
        last = self.reverse(node.next)
        node.next.next = node
        node.next = None

        return last

    def reorderList(self, head: ListNode) -> None:
        if head == None or head.next == None:
            return
        # find the middle position
        slow = head
        fast = head
        while fast.next and fast.next.next:
            slow = slow.next
            fast = fast.next.next
        slow = slow.next

        temp = head
        while head.next != slow:
            head = head.next
        head.next = None
        head = temp

        # reverse the linked list
        re_head = self.reverse(slow)

        # combine two linked list
        cur = head
        while cur and re_head:
            temp = cur.next
            cur.next = re_head
            re_head = re_head.next
            cur.next.next = temp
            cur = cur.next.next




