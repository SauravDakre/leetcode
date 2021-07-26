package sorted_array_to_bst

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return bst(nums, 0, len(nums))
}

func bst(nums []int, s, e int) *TreeNode {
	if s >= e {
		return nil
	}
	mid := s + (e-s)/2
	t := new(TreeNode)
	t.Val = nums[mid]
	t.Left = bst(nums, s, mid)
	t.Right = bst(nums, mid+1, e)
	return t
}
