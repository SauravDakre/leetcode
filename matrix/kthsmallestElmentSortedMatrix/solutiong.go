package kthsmallestElmentSortedMatrix

func kthSmallest(m [][]int, k int) int {
	l := m[0][0]
	r := m[len(m)-1][len(m[0])-1]
	for l < r {
		mid := (l + r) / 2
		c := count(m, mid)
		// fmt.Println(l, r, mid, c)
		if c >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l

}

func count(m [][]int, t int) int {
	i := 0
	j := len(m) - 1
	c := 0
	for ; i < len(m); i++ {
		for ; j >= 0; j-- {
			if m[i][j] <= t {
				c += (j + 1)
				break
			}
		}
	}
	return c
}
