package validsudoku

func isValidSudoku(board [][]byte) bool {
	lines := len(board)
	columns := len(board[0])
	columnElems := make([]map[byte]struct{}, columns)
	squareElems := make([]map[byte]struct{}, columns)
	for i := range columnElems {
		columnElems[i] = make(map[byte]struct{})
		squareElems[i] = make(map[byte]struct{})
	}

	var checkSquare int
	for i := 0; i < lines; i++ {
		lineElems := make(map[byte]struct{})
		for j := 0; j < columns; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := lineElems[board[i][j]]; ok {
				return false
			}
			if _, ok := columnElems[j][board[i][j]]; ok {
				return false
			}
			if i < 3 {
				if j < 3 {
					checkSquare = 0
				} else if j > 5 {
					checkSquare = 2
				} else {
					checkSquare = 1
				}
			} else if i > 5 {
				if j < 3 {
					checkSquare = 6
				} else if j > 5 {
					checkSquare = 8
				} else {
					checkSquare = 7
				}
			} else {
				if j < 3 {
					checkSquare = 3
				} else if j > 5 {
					checkSquare = 5
				} else {
					checkSquare = 4
				}
			}
			if _, ok := squareElems[checkSquare][board[i][j]]; ok {
				return false
			}

			lineElems[board[i][j]] = struct{}{}
			columnElems[j][board[i][j]] = struct{}{}
			squareElems[checkSquare][board[i][j]] = struct{}{}
		}
	}
	return true
}
