package quickSort

func sort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(ar []int, l, r int) {
	if l < r {
		p := partition(ar, l, r)
		quickSort(ar, l, p-1)
		quickSort(ar, p+1, r)
	}
}

func partition(ar []int, l, r int) int {
	i := l - 1
	pivot := ar[r]
	for j := l; j < r; j++ {
		if ar[j] < pivot {
			i++
			ar[i], ar[j] = ar[j], ar[i]
		}
	}
	ar[i+1], ar[r] = ar[r], ar[i+1]
	return i + 1
}
