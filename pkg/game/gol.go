package game

var dir = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	//	{0, 0},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func getNeighborCount(state [][]int, y, x int) int {
	count := 0
	for _, d := range dir {
		row := y + d[0]
		col := x + d[1]

		if row < 0 || row >= len(state) {
			continue
		}
		if col < 0 || col >= len(state[row]) {
			continue
		}

		if state[row][col] == 1 {
			count++
		}
	}
	return count
}

func PlayRound(state [][]int) [][]int {
	out := make([][]int, len(state))
	for i := range state {
		out[i] = make([]int, len(state[i]))
	}

	for y, rows := range state {
		for x := range rows {
			out[y][x] = state[y][x]

			count := getNeighborCount(state, y, x)
			if count < 2 {
				out[y][x] = 0
			}
			if count > 3 {
				out[y][x] = 0
			}
			if count == 3 {
				out[y][x] = 1
			}
		}
	}
	return out
}

func ToString(state [][]int) string {
	out := ""
	for _, row := range state {
		for _, cell := range row {
			if cell == 1 {
				out += "#"
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}
