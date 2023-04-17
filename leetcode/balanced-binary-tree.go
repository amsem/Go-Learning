func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ans := true
	var getDepth func(*TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node.Left == nil && node.Right == nil {
			return 0
		}
		l := 0
		r := 0
		if node.Left != nil {
			l = getDepth(node.Left) + 1
		}
		if node.Right != nil {
			r = getDepth(node.Right) + 1
		}

		if l > r {
			if l-r > 1 {
				ans = false
			}
			return l
		}

		if r-l > 1 {
			ans = false
		}
		return r
	}
	getDepth(root)
	return ans
}
