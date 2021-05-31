package kthpermutation

import "fmt"

func getPermutation(n int, k int) string {
	dig := make([]int, n)
	for i := 1; i <= n; i++ {
		dig[i-1] = i
	}
	m := make(map[int]int)
	res := make([]int, 0)
	for n > 0 {
		n--

		if len(dig) == 1 {
			res = append(res, dig[0])
			break
		}

		m[n] = fact(n)

		index := k / m[n]
		// edge scenario
		if k%m[n] == 0 {
			index -= 1
		}
		fmt.Println(n, k, m[n], res, index, dig)
		res = append(res, dig[index])
		k = k - m[n]*index
		dig = append(dig[:index], dig[index+1:]...)

	}
	// fmt.Println(res)
	rs := ""
	for _, r := range res {
		rs += fmt.Sprintf("%d", r)
	}
	return rs
}

func fact(n int) int {
	s := 1
	for n > 0 {
		s *= n
		n--
	}
	return s
}
