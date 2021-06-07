package longestConsecutiveSequence

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}

	ls := 0
	for _, n := range nums {
		if _, ok := m[n-1]; !ok {
			cn := n
			steak := 1
			for {
				if _, ok := m[cn+1]; !ok {
					break
				}
				cn = cn + 1
				steak += 1
				delete(m, cn)
			}
			if ls < steak {
				ls = steak
			}
		}
	}
	return ls
}
