package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
翻转一棵二叉树。
示例：

输入：

		 4
	   /   \
	  2     7
	 / \   / \
	1   3 6   9
输出：
		 4
	   /   \
	  7     2
	 / \   / \
	9   6 3   1
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	return root
}

/*
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
示例：给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

func getMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := getMaxDepth(root.Left)
	rightMax := getMaxDepth(root.Right)
	if leftMax > rightMax {
		return leftMax + 1
	} else {
		return rightMax + 1
	}
}

func main() {

}
