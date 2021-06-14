package DecodeWays

import "fmt"

func numDecodings(s string) int {
	m := make(map[string]string)
	a := rune('A')
	for i := 0; i < 26; i++ {
		m[fmt.Sprintf("%d", i+1)] = string(a + rune(i))
	}
	// fmt.Println(m)
	dp := make([]int, len(s)+2)

	return combMemoization(s, 0, m, dp)
	// cnt := 0
	// comb(s,0,m,&cnt)
	// return cnt
	// return comb(s,0,m)
}

func combMemoization(s string, i int, m map[string]string, dp []int) int {
	if i == len(s) {
		return 1
	}
	if dp[i] != 0 {
		return dp[i]
	}
	var x, y int
	if i+1 <= len(s) {
		if _, ok := m[s[i:i+1]]; ok {
			dp[i+1] = combMemoization(s, i+1, m, dp)
			x = dp[i+1]
		}
	}
	if i+2 <= len(s) {
		// fmt.Println("s[i:i+2]", s[i:i+2])
		if _, ok := m[s[i:i+2]]; ok {
			fmt.Println("s[i:i+2]", s[i:i+2])
			dp[i+2] = combMemoization(s, i+2, m, dp)
			y = dp[i+2]
		}
	}
	// fmt.Println(s, i, dp)
	return x + y
}

func combRecursion2(s string, i int, m map[string]string) int {
	fmt.Println(s, i)
	if i == len(s) {
		return 1
	}
	var x, y int
	if i+1 <= len(s) {
		if _, ok := m[s[i:i+1]]; ok {
			x = combRecursion2(s, i+1, m)
		}
	}
	if i+2 <= len(s) {
		if _, ok := m[s[i:i+2]]; ok {
			y = combRecursion2(s, i+2, m)
		}
	}
	return x + y
}

func combRecursion1(s string, i int, m map[string]string, cnt *int) {
	fmt.Println(s, i, *cnt)
	if i == len(s) {
		*cnt++
		return
	}
	if i+1 <= len(s) {
		if _, ok := m[s[i:i+1]]; ok {
			combRecursion1(s, i+1, m, cnt)
		}
	}
	if i+2 <= len(s) {
		if _, ok := m[s[i:i+2]]; ok {
			combRecursion1(s, i+2, m, cnt)
		}
	}
}
