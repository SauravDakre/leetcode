package main

import "fmt"

func main() {
	res := make([][]int, 0)
	tmp := make([]int, 0)

	combinations(&res, 0, tmp, []int{1, 2, 3, 4})
	fmt.Println(len(res), res)
}

func combinations(res *[][]int, i int, tmp, c []int) {
	if i == len(c) {
		*res = append(*res, tmp)
		// fmt.Println(res, i, tmp, c)
		return
	}
	// fmt.Println(res, i, tmp, c)
	combinations(res, i+1, tmp, c)
	tmp = append(tmp, c[i])
	combinations(res, i+1, tmp, c)
	tmp = tmp[:len(tmp)-1]
}
