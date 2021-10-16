package uniquePermutations

func permuteUnique(nums []int) [][]int {
	count := make(map[int]int)
	for _, t := range nums {
		if v, ok := count[t]; ok {
			count[t] = v + 1
		} else {
			count[t] = 1
		}
	}
	res := make([][]int, 0)
	tmp := make([]int, 0)
	generate(&res, tmp, count, len(nums))
	return res
}

func generate(res *[][]int, tmp []int, count map[int]int, n int) {
	if len(tmp) == n {
		(*res) = append((*res), tmp)
		return
	}

	for k, v := range count {
		if v == 0 {
			continue
		} else {
			count[k] = v - 1
		}
		tmp = append(tmp, k)
		generate(res, tmp, count, n)
		tmp = (tmp)[:len(tmp)-1]
		count[k] = v
	}
}
