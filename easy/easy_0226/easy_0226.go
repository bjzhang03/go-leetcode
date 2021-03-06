package easy_0226

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		result := &TreeNode{root.Val, nil, nil}
		result.Left = invertTree(root.Right)
		result.Right = invertTree(root.Left)
		return result
	}
	return nil
}
