
type Coord struct {
	r, c int
}

func islandsAndTreasure(grid [][]int) {
	isWater := func(coord Coord) bool {
		if coord.c < 0 || coord.c >= len(grid[0]) || coord.r < 0 || coord.r >= len(grid) || grid[coord.r][coord.c] == -1 {
			return true
		}
		return false
	}

	isChest := func(coord Coord) bool {
		if grid[coord.r][coord.c] == 0 {
			return true
		}
		return false
	}

	directions := []Coord{Coord{1, 0}, Coord{-1, 0}, Coord{0, 1}, Coord{0, -1}}
	var visited map[Coord]bool
	var distance map[Coord]int
	bfs := func(start Coord) int {
		q := []Coord{start}
		visited = make(map[Coord]bool)
		distance = make(map[Coord]int)
		visited[start] = true
		distance[start] = 0
		for len(q) > 0 {
			current := q[0]
			q = q[1:]
			if isChest(current) {
				return distance[current]
			}
			for _, d := range directions {
				neighbour := Coord{
					r: current.r + d.r,
					c: current.c + d.c,
				}

				if !isWater(neighbour) && !visited[neighbour] {
					q = append(q, neighbour)
					visited[neighbour] = true
					distance[neighbour] = distance[current] + 1
				}

			}
		}

		return grid[start.r][start.c]
	}

	// start := Coord{r: 0, c: 3}
	// grid[0][0] = bfs(start)

	for ridx, row := range grid {
		for cidx := range row {
			start := Coord{r: ridx, c: cidx}
			if !isWater(start) && !isChest(start) {
				res := bfs(start)
				fmt.Println(ridx, "-", cidx, " => ", res)
				grid[ridx][cidx] = res
			}

		}
	}
}
