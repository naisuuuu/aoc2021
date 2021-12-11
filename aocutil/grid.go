package aocutil

var diagonal = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
var adjacent = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func inBounds(grid [][]int, r, c int) bool {
	if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) {
		return true
	}
	return false
}

func Neighbors(grid [][]int, row, col int) [][2]int {
	var out [][2]int
	for _, c := range adjacent {
		nr, nc := row+c[0], col+c[1]
		if inBounds(grid, nr, nc) {
			out = append(out, [2]int{nr, nc})
		}
	}
	return out
}

func NeighborsDiag(grid [][]int, row, col int) [][2]int {
	var out [][2]int
	for _, c := range diagonal {
		nr, nc := row+c[0], col+c[1]
		if inBounds(grid, nr, nc) {
			out = append(out, [2]int{nr, nc})
		}
	}
	return out
}
