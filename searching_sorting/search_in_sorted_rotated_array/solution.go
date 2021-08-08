package search_in_sorted_rotated_array

func search(nums []int, target int) int {
	return searchb(nums, 0, len(nums)-1, target)
}

func searchb(nums []int, l, r, target int) int {
	if l > r {
		return -1
	}

	mid := (l + r) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[l] <= nums[mid] {
		if target >= nums[l] && nums[mid] >= target {
			return searchb(nums, l, mid-1, target)
		}
		return searchb(nums, mid+1, r, target)
	}

	if target >= nums[mid] && target <= nums[r] {
		return searchb(nums, mid+1, r, target)
	}

	return searchb(nums, l, mid-1, target)
}
