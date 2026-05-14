
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


	directions := []Coord{Coord{1, 0}, Coord{-1, 0}, Coord{0, 1}, Coord{0, -1}}
	var visited map[Coord]bool
	msBfs := func(q []Coord) {
		visited = make(map[Coord]bool)
        for _, val := range q {
            visited[val] = true
            grid[val.r][val.c] = 0

        }
		for len(q) > 0 {
			current := q[0]
			q = q[1:]
			for _, d := range directions {
				n := Coord{
					r: current.r + d.r,
					c: current.c + d.c,
				}

				if !isWater(n) && !visited[n] {
					q = append(q, n)
					visited[n] = true
                    grid[n.r][n.c] = min(grid[n.r][n.c], grid[current.r][current.c]+1)
				}
			}
		}
	}

	// start := Coord{r: 0, c: 3}
	// grid[0][0] = bfs(start)
    q := []Coord{}
	for ridx, row := range grid {
		for cidx, val := range row {
			current := Coord{r: ridx, c: cidx}
			if val == 0 {
                q = append(q, current)
            }
		}
	}
    msBfs(q)
}
