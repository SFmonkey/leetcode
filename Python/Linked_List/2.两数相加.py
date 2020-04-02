# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
	def addTwoNumbers(self, l1, l2):
		"""
		:type l1: ListNode
		:type l2: ListNode
		:rtype: ListNode
		"""
		head = ListNode(0)
		node = head
		ADD = 0
		while l1 != None or l2 != None or ADD != 0:
			if l1 != None:
				ADD += l1.val
				l1 = l1.next
			if l2 != None:
				ADD += l2.val
				l2 = l2.next

			node.next = ListNode(ADD%10)
			if ADD >= 10:
				ADD = int(ADD/10)

		return head.next

