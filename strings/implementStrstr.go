package strings

/*
Implement strStr().
Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:
Input: haystack = "hello", needle = "ll"
Output: 2

Example 2:
Input: haystack = "aaaaa", needle = "bba"
Output: -1
*/

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
outer:
	for i, c := range haystack {
		if i+len(needle) <= len(haystack) && string(c) == string(needle[0]) && string(haystack[i+len(needle)-1]) == string(needle[len(needle)-1]) {
			for j, k := i, 0; k < len(needle); j, k = j+1, k+1 {
				if string(haystack[j]) != string(needle[k]) {
					continue outer
				}
			}
			return i
		}
	}
	return -1
}
