package robhouse

func rob(nums []int) int {
	inc := 0
	exc := 0
	for i := 0; i < len(nums); i++ {
		t := maxf(exc+nums[i], inc)
		exc = inc
		inc = t

	}
	return inc
}

func robAnotherWay(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return maxf(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxf(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = maxf(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func maxf(x, y int) int {
	if x > y {
		return x
	}
	return y
}
