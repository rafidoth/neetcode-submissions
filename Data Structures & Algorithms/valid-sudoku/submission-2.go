
type nothing struct{}

const (
	ss = 3 // small square side size
)


func isValidRow(board [][]byte, r int) bool {
	set := make(map[byte]nothing)
	for _, val := range board[r] {	
		if val == '.' {
			continue
		}
		if _, exists := set[val]; exists {
				return false
		}
		set[val] = nothing{}
	}
	return true
}

func isValidCol(board [][]byte, c int) bool {
	set := make(map[byte]nothing)
	for _, row := range board {	
		val := row[c]
		if val == '.' {
			continue
		}
		if _, exists := set[val]; exists {
				return false
		}
		set[val] = nothing{}
		
	}
	return true
}




func isValidSmallSquare(board [][]byte, r,c int) bool {
	set := make(map[byte]nothing)
	endR, endC := r+ss, c+ss
	for _, row := range board[r:endR] {
		for _, val := range row[c:endC] {
			if val == '.' {
				continue
			}
			if _, exists := set[val]; exists {
				return false
			}
			set[val] = nothing{}
		} 
	}
	return true
}


func isValidSudoku(board [][]byte) bool {
	// input is static 9x9 board so creating 
	// start points manually make sense

	// since 9X9 Square -> 9X(3x3) squares
	// all starting points aare
	// 00, 03, 06
	// 30, 33, 36
	// 60, 63, 66

	starts := [][2]int{
		{0,0}, {0,3}, {0,6},
		{3,0}, {3,3}, {3,6},
		{6,0}, {6,3}, {6,6},
	}


	for _, x := range starts {
		if !isValidSmallSquare(board, x[0], x[1]) {
			return false
		}
	}

	// small squares are vallid already at this point
	fmt.Println(1)

	for x := range 9 {
		if !isValidRow(board, x) {
			return false
		}
	}
	fmt.Println(2)


	// small squares and rows are valid
	for x := range 9 {
		if !isValidCol(board, x){
			return false
		}
	}
	fmt.Println(3)
	return true
}


