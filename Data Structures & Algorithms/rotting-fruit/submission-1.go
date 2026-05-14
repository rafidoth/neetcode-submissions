type RC struct {
	r,c int
}



func orangesRotting(grid [][]int) int {
	validRC:= func (r,c int) bool {
		if r <0 || r >= len(grid)|| c <0 || c >= len(grid[0]){
			return false
		}
		return true
	}
	isEmpty := func (r,c int) bool {
		if validRC(r,c) {
			return grid[r][c] == 0
		}
		return true
	}
	freshCount := 0
    for i := range grid {
		for j := range grid[i]{
			if grid[i][j] == 1{
				freshCount++
			}
		}
	}
	directions := []RC{RC{1,0}, RC{-1,0}, RC{0,1}, RC{0,-1}}
	time := make(map[RC]int)
	visited := make(map[RC]bool)
	queue := []RC{}
	bfs:= func(){
		for _,val := range queue {
			visited[val] = true
			time[val] = 0
		}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			visited[curr] = true
			for _, val := range directions {
				n:= RC{
					r : curr.r + val.r,
					c : curr.c + val.c,
				}
				if !isEmpty(n.r, n.c) && !visited[n] {
					queue = append(queue, n)
					val, ok := time[n]
					if ok {
						time[n]= min(val, time[curr]+1)
					}else{
						time[n] = time[curr]+1
					}
					
				}
			}
		}
	}


	for i := range grid {
		for j := range grid[i]{
			rc := RC{
				r : i,
				c : j,
			}
			if grid[i][j] == 2{
				queue = append(queue, rc)
			}
		}
	}
	bfs()
	res := 0
	freshGotRottenCount := 0
	for _,val:= range time{
		res = max(res,val)
		if val > 0 {
			freshGotRottenCount++
		}
	}
	fmt.Println("fresh ", freshCount, " GotRotten ", freshGotRottenCount)
	if freshCount > 0 && freshCount-freshGotRottenCount != 0{
		return -1
	}

	return res
}
