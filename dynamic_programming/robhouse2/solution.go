package robhouse2

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(rob1(nums[0:len(nums)-1]), rob1(nums[1:len(nums)]))
}

func rob1(nums []int) int {
	inc := 0
	exc := 0
	for _, r := range nums {

		t := max(exc+r, inc)
		exc, inc = inc, t
		// fmt.Println(inc, exc)
	}
	return max(exc, inc)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
