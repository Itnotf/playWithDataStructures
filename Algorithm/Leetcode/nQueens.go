package Leetcode

func solveNQueens(n int) [][]string {
	var res [][]string
	board := initBoard(n)
	nQueen(n, &res, board, 0)
	return res
}

func initBoard(n int) [][]byte {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	return board
}

func nQueen(n int, res *[][]string, board [][]byte, row int) {
	if row == n {
		//board 是一个解，返回
		var boardStr []string
		for _, b := range board {
			boardStr = append(boardStr, string(b))
		}
		*res = append(*res, boardStr)
		return
	}

	for col := 0; col < n; col++ {
		if !isValidQ(board, row, col) {
			continue
		}
		board[row][col] = 'Q'
		nQueen(n, res, board, row+1)
		board[row][col] = '.'
	}
}

func isValidQ(board [][]byte, row int, col int) bool {
	// col
	// j-i=col-row
	// j+i=col+row
	n := len(board)
	for i := 0; i < row; i++ {
		for j := 0; j < n; j++ {
			if j == col && board[i][j] == 'Q' {
				return false
			}
			if j-i == col-row && board[i][j] == 'Q' {
				return false
			}
			if j+i == col+row && board[i][j] == 'Q' {
				return false
			}
		}
	}
	return true
}
