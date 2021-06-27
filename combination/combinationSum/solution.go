package combinationSum

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0)
	gen(candidates, target, 0, &tmp, &res)
	return res
}

func gen(c []int, t, i int, tmp *[]int, res *[][]int) {
	if t == 0 {
		r := make([]int, len(*tmp))
		copy(r, *tmp)
		*res = append(*res, r)

		return
	}

	if i >= len(c) || t < 0 {
		return
	}
	for ; i < len(c); i++ {

		*tmp = append(*tmp, c[i])
		gen(c, t-c[i], i, tmp, res)
		*tmp = (*tmp)[:len(*tmp)-1]
	}

}
