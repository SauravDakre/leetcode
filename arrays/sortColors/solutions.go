package sortColors

func sortColors(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	s := 0
	e := len(nums) - 1
	mid := 0
	for mid <= e && s < e {
		if nums[mid] == 0 {
			nums[s], nums[mid] = nums[mid], nums[s]
			mid++
			s++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[mid], nums[e] = nums[e], nums[mid]
			// we are not incrementing mid, because when end will be at len(nums)-1 position and if we are swapping
			// we donot know what is present in last position, so in next iteration we will check what we swapped by
			// checking at mid position and moving that value to appropriate postion
			e--
		}
		// fmt.Println(nums,s,mid,e)
	}

}
