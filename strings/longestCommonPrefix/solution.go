package longestCommonPrefix

import "strings"

func longestCommonPrefix(strs []string) string {
	minl := 99999
	for _, s := range strs {
		if minl > len(s) {
			minl = len(s)
		}
	}
	l := 0
	r := minl
	for l <= r {
		mid := (l + r) / 2
		if isCommon(strs, mid) {

			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return strs[0][0 : (l+r)/2]
}

func isCommon(strs []string, m int) bool {
	s := strs[0][0:m]
	for _, t := range strs {
		if !strings.HasPrefix(t, s) {
			return false
		}
	}
	return true
}
