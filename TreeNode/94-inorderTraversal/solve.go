package main

/**
  题目链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 中序遍历： 左子树 -> 根结点 -> 右子树
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return res
}

func inorderTraversal_one(root *TreeNode) []int {
	var stack []*TreeNode
	res := make([]int, 0)
	for root != nil || len(stack) != 0{
		for root != nil{
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}


func main()  {

}