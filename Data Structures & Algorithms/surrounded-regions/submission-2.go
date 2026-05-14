func solve(board [][]byte) {
	type P struct {
		r,c int
	}
	rows := len(board)
	cols := len(board[0])

	bordered := func(p P)bool {
		if p.r == 0 || p.r == rows -1 || p.c == 0 || p.c == cols-1 {
			return true
		}
		return false
	}
	cross := func(p P)bool {
		if board[p.r][p.c] == 'X' {
			return true
		}
		return false
	}
	directions := [][]int{
		{1,0}, {-1,0}, {0, 1}, {0, -1},
	}
	vis := make(map[P]bool)
	var dfs func(p P)
	dfs = func (p P){
		if p.r < 0 || p.c < 0 || p.r >=rows || p.c >= cols {
			return
		}
		if cross(p) {
			return
		}
		vis[p] = true
		for _, d := range directions{
			next := P{r:p.r + d[0], c : p.c + d[1]}
			if !vis[next] {
				dfs(next)
			}
		} 
	}

	
	for rid , row := range board {
		for cid, val := range row{
			point := P{r : rid, c : cid}
			if val == 'O' && bordered(point) {
				dfs(point)
			}
		}
	}

	for rid , row := range board {
		for cid, _ := range row{
			point := P{r : rid, c : cid}
			if !cross(point) && !vis[point]{
				board[rid][cid] = 'X'
			} 
		}
	}
	

}
