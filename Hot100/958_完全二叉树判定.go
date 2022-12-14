package Hot100

func isCompleteTree(root *TreeNode) bool {
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 && q[0] != nil {
		q = append(q, q[0].Left, q[0].Right)
		q = q[1:]
	}

	for _, v := range q {
		if v != nil {
			return false
		}
	}

	return true
}
