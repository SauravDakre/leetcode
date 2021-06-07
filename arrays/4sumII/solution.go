package sumII

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			if cnt, ok := m[n1+n2]; ok {
				m[n1+n2] = cnt + 1
			} else {
				m[n1+n2] = 1
			}
		}
	}
	count := 0
	for _, n1 := range nums3 {
		for _, n2 := range nums4 {
			if cnt, ok := m[-(n1 + n2)]; ok {
				count += cnt
			}
		}
	}
	return count
}
