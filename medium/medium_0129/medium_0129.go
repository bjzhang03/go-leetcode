package medium_0129

func dfs(root *TreeNode, current int, save *int) {
	if root.Left == nil && root.Right == nil {
		current = current*10 + root.Val
		*save = *save + current
		current = current / 10
	} else {

		if root.Left != nil {
			current = current*10 + root.Val
			dfs(root.Left, current, save)
			current = current / 10
		}
		if root.Right != nil {
			current = current*10 + root.Val
			dfs(root.Right, current, save)
			current = current / 10
		}
	}

}

func sumNumbers(root *TreeNode) int {
	result := 0
	if root != nil {
		dfs(root, 0, &result)
	}
	return result
}
