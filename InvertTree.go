package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root

}
func printPreOrder(root *TreeNode, level int) {
	if root == nil {
		return
	}

	printPreOrder(root.Right, level+1)

	for i := 0; i < level; i++ {
		fmt.Print("    ")

	}
	fmt.Println(root.Val)
	printPreOrder(root.Left, level+1)
	// rep 1 -> 5 != nil. rep 2 -> 7 != nil. print 7.

}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println("Before Invert : \n")
	printPreOrder(root, 0)

	invertTree(root)

	fmt.Println("\n After Invert: \n")
	printPreOrder(root, 0)
}
