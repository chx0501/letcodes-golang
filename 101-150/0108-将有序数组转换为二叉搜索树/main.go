package _108_将有序数组转换为二叉搜索树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 二叉搜索树: 中序: 左根右

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0{
		return nil
	}
	rootIndex := len(nums) / 2
	return &TreeNode{
		Val:   nums[rootIndex],
		Left:  sortedArrayToBST(nums[:rootIndex]),
		Right: sortedArrayToBST(nums[rootIndex+1:]),
	}
}
