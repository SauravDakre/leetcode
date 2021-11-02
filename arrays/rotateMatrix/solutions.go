package rotateMatrix

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {
		k, l := 0, len(matrix[i])-1
		for k < l {
			matrix[i][k], matrix[i][l] = matrix[i][l], matrix[i][k]
			k++
			l--
		}
	}

}
