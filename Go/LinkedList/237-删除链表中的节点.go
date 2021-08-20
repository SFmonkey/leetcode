package LinkedList

// 有点奇葩，给的是一个非末尾的节点，也就是不知道前一个节点；把下一个节点的内容复制到当前节点，然后从链表中删除下一个节点
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
