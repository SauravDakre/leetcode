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
