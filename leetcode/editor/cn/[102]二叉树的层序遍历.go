//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
//
// Related Topics 树 广度优先搜索 二叉树 👍 1484 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int,0)
	if root==nil {
		return res
	}
	queue := make([]*TreeNode,0)
	queue = append(queue,root)
	for root!=nil && len(queue)>0{
		arr := make([]int,0)
		size := len(queue)
		for i:=0 ;i<size;i++{
			node := queue[0]
			if node != nil {
				arr = append(arr,node.Val)
			}
			if node.Left != nil {
				queue = append(queue,node.Left)
			}
			if node.Right!=nil {
				queue = append(queue,node.Right)
			}
			queue = queue[1:]
		}
		res = append(res,arr)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
