package hammingWeight

func hammingWeight(num uint32) int {
	c := 0
	for num != 0 {
		c++
		num &= (num - 1)
	}
	return c
}
