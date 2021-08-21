package search2dMatrix

func searchMatrix(matrix [][]int, target int) bool {
	x := 0
	y := len(matrix[0]) - 1
	for x < len(matrix) && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

// using binary search
func searchMatrixViaBinarySearch(matrix [][]int, target int) bool {
	s := 0
	e := len(matrix) - 1
	for s <= e {

		mid := (s + e) / 2
		if target <= matrix[mid][len(matrix[0])-1] && target >= matrix[mid][0] {
			return binSearch1D(matrix[mid], target)
		}
		if target > matrix[mid][0] {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	return false
}

func binSearch1D(arr []int, k int) bool {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == k {
			return true
		}
		if k > arr[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
