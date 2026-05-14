func pacificAtlantic(heights [][]int) [][]int {
    // 1. Define the struct (can be inside or outside)
    type P struct {
        r, c int
    }

    rows := len(heights)
    if rows == 0 {
        return [][]int{}
    }
    cols := len(heights[0])

    // 2. Correct function variable assignment
    flowable := func(next P, curr P) bool {
        // Bounds check is MANDATORY before accessing heights
        if next.r < 0 || next.r >= rows || next.c < 0 || next.c >= cols {
            return false
        }
        return heights[next.r][next.c] <= heights[curr.r][curr.c]
    }

    directions := []P{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    var pacific, atlantic bool
    
    // 3. Declare DFS variable first to allow recursion
    var dfs func(point P, vis map[P]bool)
    dfs = func(point P, vis map[P]bool) {
        vis[point] = true

        // Check boundaries for ocean reach
        if point.r == 0 || point.c == 0 {
            pacific = true
        }
        if point.r == rows-1 || point.c == cols-1 {
            atlantic = true
        }

        // Optimization: stop if both found
        if pacific && atlantic {
            return
        }

        for _, d := range directions {
            next := P{r: point.r + d.r, c: point.c + d.c}
            // Check flowable and if NOT visited
            if flowable(next, point) && !vis[next] {
                dfs(next, vis)
            }
        }
    }

    result := [][]int{}
    for rr := 0; rr < rows; rr++ {
        for cc := 0; cc < cols; cc++ {
            pacific = false
            atlantic = false
            vis := make(map[P]bool)
            dfs(P{rr, cc}, vis)
            
            if pacific && atlantic {
                result = append(result, []int{rr, cc})
            }
        }
    }
    return result
}