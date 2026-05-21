
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]struct{}, 9)
	cols := make([]map[byte]struct{}, 9)
	squares := make([]map[byte]struct{}, 9)

	for i := 0; i < 9; i++ {
        rows[i] = make(map[byte]struct{})
        cols[i] = make(map[byte]struct{})
        squares[i] = make(map[byte]struct{})
    }


	for i, row := range board {
		for j , val := range row {
			if val == '.' {
				continue
			}

			if _, exists := rows[i][val] ; exists {
				return false
			}
			rows[i][val] = struct{}{}
			
			if _, exists := cols[j][val] ; exists {
				return false
			}
			cols[j][val] = struct{}{}

			row := i/3
			col := j/3
			sqIdx := row*3 + col
			if _, exists := squares[sqIdx][val]; exists {
				return false
			}
			squares[sqIdx][val] = struct{}{}
		}
	}
	return true


}
