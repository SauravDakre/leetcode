package productExceptSelf

func productExceptSelf(nums []int) []int {
	op := make([]int, len(nums))
	i := 1
	op[0] = nums[0]
	for i = 1; i < len(nums); i++ {
		op[i] = op[i-1] * nums[i]
	}
	// fmt.Println(op)
	i--
	op[i] = op[i-1]
	// fmt.Println(op)
	product := 1 * nums[i]
	i--
	for ; i > 0; i-- {
		op[i] = product * op[i-1]
		product *= nums[i]
	}
	op[i] = product
	return op
}
