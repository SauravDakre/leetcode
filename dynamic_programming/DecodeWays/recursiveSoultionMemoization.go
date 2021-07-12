package DecodeWays

import "fmt"

func numDecodings(s string) int {
	m := make(map[string]bool)
	for i := 1; i <= 26; i++ {
		m[fmt.Sprintf("%d", i)] = true
	}
	// fmt.Println(m)
	mo := make(map[string]int)
	return ways(s, m, mo)
}

func ways(s string, m map[string]bool, mo map[string]int) int {
	// fmt.Println("=",s)
	if c, ok := mo[s]; ok {
		return c
	}
	if len(s) == 0 {
		return 1
	}
	x := 0
	y := 0
	if len(s) >= 1 {
		if _, ok := m[string(s[:1])]; ok {
			x = ways(s[1:], m, mo)
		}
	}
	if len(s) >= 2 {
		if _, ok := m[string(s[:2])]; ok {
			y = ways(s[2:], m, mo)
		}
	}
	mo[s] = x + y
	return x + y
}
