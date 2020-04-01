# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
	def addExtra(self, l, ADD, node):
		o_node = node
		while l != None:
			var = l.val
			if ADD:
				var += 1
				ADD = False
			if var >= 10:
				var -= 10
				ADD = True

			node.next = ListNode(var)
			node = node.next
			l = l.next

		if ADD:
			node.next = ListNode(1)

	def addTwoNumbers(self, l1, l2):
		"""
		:type l1: ListNode
		:type l2: ListNode
		:rtype: ListNode
		"""
		head = ListNode(0)
		node = head
		ADD = False
		while l1 != None and l2 != None:
			var = l1.val + l2.val
			if ADD:
				var += 1
				ADD = False
			if var >= 10:
				var -= 10
				ADD = True

			node.next = ListNode(var)
			node = node.next
			l1 = l1.next
			l2 = l2.next

		if l1 == None and l2 == None:
			if ADD:
				node.next = ListNode(1)
			return head.next
		elif l1 != None:
			self.addExtra(l1, ADD, node)
		else:
			self.addExtra(l2, ADD, node)

		return head.next
