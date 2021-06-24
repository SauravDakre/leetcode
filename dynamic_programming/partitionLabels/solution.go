package partitionLabels

func partitionLabels(s string) []int {
	var last [26]int
	res := make([]int, 0)
	for i, t := range s {
		last[int(rune(t)-rune('a'))] = i
	}
	j, anchor := 0, 0

	for i := 0; i < len(s); i++ {
		if j < last[int(rune(s[i])-rune('a'))] {
			j = last[int(rune(s[i])-rune('a'))]
		}
		if j == i {
			res = append(res, (i - anchor + 1))
			anchor = i + 1
		}
	}
	return res
}
