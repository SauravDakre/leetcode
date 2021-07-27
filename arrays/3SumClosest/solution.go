func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	const MaxUint = ^uint(0)
	// fmt.Println(nums)
	diff := int(MaxUint >> 1)

	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {

			sum := nums[i] + nums[j] + nums[k]
			if abs(target-sum) < abs(diff) {
				diff = target - sum
			}

			if sum > target {
				k--
			} else {
				j++
			}
			// fmt.Println(i,j,k, sum, closestSum)
		}
	}
	return target - diff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// [1,1,-1,-1,3]
// -1
// ans: -1