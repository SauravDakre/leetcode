package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(ar []int, l, r, k int) int {
	t := partition(ar, l, r)
	if t == k {
		return ar[t]
	} else if t > k {
		return quickSort(ar, l, t-1, k)
	} else {
		return quickSort(ar, t+1, r, k)
	}

}

func partition(ar []int, l, r int) int {
	pivot := ar[r]
	i := l - 1
	for j := l; j < r; j++ {
		if ar[j] < pivot {
			i++
			ar[i], ar[j] = ar[j], ar[i]
		}
	}
	ar[i+1], ar[r] = ar[r], ar[i+1]
	return i + 1
}

func findKthSmallest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, k-1)
}

func main() {
	ar := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(ar, 2))
	fmt.Println(findKthSmallest(ar, 2))
}
