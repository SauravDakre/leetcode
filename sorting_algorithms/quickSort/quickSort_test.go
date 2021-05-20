package quickSort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	ar := []int{9, 8, 7, 6, 5}
	sort(ar)
	fmt.Println(ar)
}
