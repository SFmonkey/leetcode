class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

def detectCycle(head: ListNode) -> ListNode:
    if head == None or head.next == None:
        return None
    slow, fast = head, head
    while fast != None and fast.next.next != None:
        slow = slow.next
        fast = fast.next.next

        if slow == fast:
            fast = head
            ind = 0
            while fast != slow:
                fast = fast.next
                slow = slow.next
                ind += 1
            return ind
    return None
