package main

func main() {

}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if del(root) {
		return nil
	}
	return root
}

func del(root *TreeNode) bool {

	if root.Left != nil && del(root.Left) {
		root.Left = nil
	}

	if root.Right != nil && del(root.Right) {
		root.Right = nil
	}
	return root.Right == nil && root.Left == nil && root.Val == 0
}
