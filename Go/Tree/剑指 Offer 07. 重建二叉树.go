package Tree

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 前序的第一个节点是根节点
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	// 在中序中找到根节点所在位置 => i
	i := 0
	for ; i<len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	// 中序遍历中根节点左边的节点都是左子树 即：左子树长度可以确定，对应前序中的范围为 => 1:len(inorder[:i])+1
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	// 中序遍历中根节点右边的节点都是右子树 即：右子树长度可以确定，对应前序中的范围为 => len(inorder[:i])+1:
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
