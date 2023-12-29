package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	root := NewTreeNode(4)
	data := []int{3, 5, 1, 2, 12, 7, 15}

	for _, v := range data {
		Insert(root, v)
	}
	fmt.Println(levelOrder(root))
}

// 100. Same Tree
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true // Если узел - лист и его значение равно targetSum

	}

	targetSum -= root.Val

	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

//101. Symmetric Tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(l *TreeNode, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == nil && r == nil
	}
	if l.Val != r.Val {
		return false
	}
	return check(l.Left, r.Right) && check(l.Right, r.Left)
}

// 199. Binary Tree Right Side View
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	preOrder2(root, &res, 0)
	return res
}

func preOrder2(node *TreeNode, res *[]int, level int) {
	if node == nil {
		return
	}
	if level == len(*res) {
		*res = append(*res, node.Val)
	}

	preOrder2(node.Right, res, level+1) // Идем сначала вправо, чтобы добавить самые правые значения первыми
	preOrder2(node.Left, res, level+1)
}

// 102. Binary Tree Level Order Traversal
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	preOrder(root, &res, 0)
	return res
}
func preOrder(node *TreeNode, res *[][]int, level int) {
	if node == nil {
		return
	}
	if level == len(*res) {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], node.Val)

	preOrder(node.Left, res, level+1)
	preOrder(node.Right, res, level+1)
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preorderHelp(root, &res)
	return res
}

func preorderHelp(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	*res = append(*res, node.Val)
	preorderHelp(node.Left, res)
	preorderHelp(node.Right, res)
}

///

func NewTreeNode(newVal int) *TreeNode {
	return &TreeNode{Val: newVal}
}

func Insert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return NewTreeNode(value)
	}

	if value < root.Val {
		root.Left = Insert(root.Left, value)
	} else {
		root.Right = Insert(root.Right, value)
	}
	return root
}

func PreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func InOrderTraversal(root *TreeNode) {
	if root != nil {
		InOrderTraversal(root.Left)
		fmt.Println(root.Val)
		InOrderTraversal(root.Right)
	}
}

func PostOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.Left)
	PostOrderTraversal(root.Right)
	fmt.Println(root.Val)
}

func minRoot(root *TreeNode, min *int) {
	if root == nil {
		return
	}
	*min = root.Val
	minRoot(root.Left, min)
}

func maxRoot(root *TreeNode, max *int) {
	if root == nil {
		return
	}
	maxRoot(root.Right, max)
	*max = root.Val
}

/*
-----PreOrderTraversal
4
1
2
5
12
7
-----InOrderTraversal
1
2
4
5
7
12
-----PostOrderTraversal
2
1
7
12
5
4

*/
