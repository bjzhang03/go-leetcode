package medium_0200

func dfs(grid [][]byte, save *[][]bool, row int, col int) {
	if grid[row][col] == '1' && (*save)[row][col] == false {
		(*save)[row][col] = true
		if row+1 < len(grid) {
			dfs(grid, save, row+1, col)
		}
		if row-1 >= 0 {
			dfs(grid, save, row-1, col)
		}
		if col+1 < len(grid[0]) {
			dfs(grid, save, row, col+1)
		}
		if col-1 >= 0 {
			dfs(grid, save, row, col-1)
		}
	}
}

func numIslands(grid [][]byte) int {
	result := 0
	if len(grid) > 0 {
		save := make([][]bool, len(grid))
		for i := 0; i < len(save); i++ {
			save[i] = make([]bool, len(grid[0]))
		}
		//fmt.Println(save)
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == '1' && save[i][j] == false {
					result = result + 1
					dfs(grid, &save, i, j)
				}
			}
		}
	}
	return result
}
