//给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[20,9],[15,7]]
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
// -100 <= Node.val <= 100
//
//
// Related Topics 树 广度优先搜索 二叉树 👍 706 👎 0

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

func zigzagLevelOrder(root *TreeNode) [][]int {
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
		if len(res)% 2 !=0{
			arr = zrev(arr)
		}
		res = append(res,arr)
	}
	return res

}

func zrev(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i],arr[j] = arr[j],arr[i]
	}
	return arr

}

//leetcode submit region end(Prohibit modification and deletion)
