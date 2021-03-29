/*
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.
*/
func binarySearch(nums []int, target int) int {
	l, mid, r := 0, len(nums)/2, len(nums)-1
	for l <= r {
		mid = (l + r) / 2
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else if target == nums[mid] {
			return mid
		}
	}
	return -1
}
