package noOfIsland

func numIslands(grid [][]byte) int {

	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if string(grid[i][j]) == "1" {
				max++
				dfs(grid, i, j)
			}
		}
	}
	return max
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || string(grid[i][j]) == "0" {
		return
	}
	grid[i][j] = []byte("0")[0]
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
