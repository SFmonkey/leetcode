class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    # 快慢指针判断链表是否有环
    def hasCycle(self, head: ListNode) -> bool:
        if head == None or head.next == None:
            return False

        fast = head
        slow = head

        while fast.next != None and fast.next.next != None:
            fast = fast.next.next
            slow = slow.next

            if fast == slow:
                return True

        return False
