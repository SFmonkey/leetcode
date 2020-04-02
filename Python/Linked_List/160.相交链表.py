# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
	def getIntersectionNode(self, headA, headB):
		"""
		:type head1, head1: ListNode
		:rtype: ListNode
		"""
		node = headA
		lenA = 0
		while node != None:
			node = node.next
			lenA += 1
		node = headB
		lenB = 0
		while node != None:
			node = node.next
			lenB += 1

		if lenA >= lenB:
			start = headA
			node = headB
			diff = lenA - lenB
		else:
			start = headB
			node = headA
			diff = lenB - lenA

		for i in range(diff):
			start = start.next

		while start != None:
			if start == node:
				return start
			start = start.next
			node = node.next

			