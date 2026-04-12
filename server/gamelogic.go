package server

import (
	"fmt"
)

type Player1 struct {
	Name  string
	Wins  int
	Color int
}

type Player2 struct {
	Name  string
	Wins  int
	Color int
}

// Initializes the two players
func InitPlayers(name1 string, name2 string) (Player1, Player2) {
	p1 := Player1{
		Name:  name1,
		Wins:  0,
		Color: 1, // Red
	}
	p2 := Player2{
		Name:  name2,
		Wins:  0,
		Color: 2, // Yellow
	}
	return p1, p2
}

// Colors for each player
func TokenChoice(p1 *Player1, p2 *Player2) {
	p1.Color = 1
	p2.Color = 2
}

// Initializes a 6x7 grid
func InitGrid() [][]int {
	grid := make([][]int, 6)
	for i := range grid {
		grid[i] = make([]int, 7)
	}
	return grid
}

// Put the token of Player 1 in a given column
func PutToken1(grid [][]int, p1 Player1, columnChoice int) ([][]int, bool, string) {
	if columnChoice < 0 || columnChoice >= 7 {
		return grid, false, fmt.Sprintf("invalid column")
	}
	for i := len(grid) - 1; i >= 0; i-- {
		if grid[i][columnChoice] == 0 {
			grid[i][columnChoice] = p1.Color
			return grid, true, fmt.Sprintf("Player 1 played in column %d", columnChoice+1)
		}
	}
	return grid, false, fmt.Sprintf("Column %d full!", columnChoice+1) // Column full
}

// Put the token of Player 2 in a given column
func PutToken2(grid [][]int, p2 Player2, columnChoice int) ([][]int, bool, string) {
	if columnChoice < 0 || columnChoice >= 7 {
		return grid, false, fmt.Sprintf("invalid column")
	}
	for i := len(grid) - 1; i >= 0; i-- {
		if grid[i][columnChoice] == 0 {
			grid[i][columnChoice] = p2.Color
			return grid, true, fmt.Sprintf("Player 2 played in column %d", columnChoice+1)
		}
	}
	return grid, false, fmt.Sprintf("Column %d full!", columnChoice+1) // Column full
}

// GAME LOOP (NOT USED)
// (handled via HTTP events)
func GameLoop(player1 Player1, player2 Player2, grid [][]int) [][]int {
	for !GridFull(grid) {
		// In web version, this is replaced by button clicks
		if CheckWin(grid) {
			return grid
		}
	}
	return grid
}

// Verifies if any player wins (4 tokens horizontally, vertically or diagonally)
func CheckWin(grid [][]int) bool {
	rows := len(grid)
	columns := len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			player := grid[r][c]
			if player == 0 {
				continue
			}

			// Horizontal
			if c+3 < columns &&
				grid[r][c+1] == player &&
				grid[r][c+2] == player &&
				grid[r][c+3] == player {
				return true
			}

			// Vertical
			if r+3 < rows &&
				grid[r+1][c] == player &&
				grid[r+2][c] == player &&
				grid[r+3][c] == player {
				return true
			}

			// Diagonal right-down
			if r+3 < rows && c+3 < columns &&
				grid[r+1][c+1] == player &&
				grid[r+2][c+2] == player &&
				grid[r+3][c+3] == player {
				return true
			}

			// Diagonal left-down
			if r+3 < rows && c-3 >= 0 &&
				grid[r+1][c-1] == player &&
				grid[r+2][c-2] == player &&
				grid[r+3][c-3] == player {
				return true
			}
		}
	}
	return false
}

func playMove(columnChoice int) [][]int {
	var valid bool

	if currentPlayer == 1 {
		grid, valid, _ = PutToken1(grid, player1, columnChoice)
	} else {
		grid, valid, _ = PutToken2(grid, player2, columnChoice)
	}

	if !valid {
		return grid
	}

	if CheckWin(grid) {
		winner = currentPlayer
		return grid
	}

	// changer de joueur
	if currentPlayer == 1 {
		currentPlayer = 2
	} else {
		currentPlayer = 1
	}

	return grid
}
