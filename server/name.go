package server

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name1, name2 string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Player 1, what's your name?")
	for {
		name1, _ = reader.ReadString('\n')
		name1 = strings.TrimSpace(name1)

		if IsAlpha(name1) {
			break
		}
		fmt.Println("Letters only")
	}

	fmt.Println("Player 2, what's your name?")
	for {
		name2, _ = reader.ReadString('\n')
		name2 = strings.TrimSpace(name2)

		if IsAlpha(name2) {
			break
		}
		fmt.Println("Letters only")
	}

	player1, player2 := InitPlayers(name1, name2)
	TokenChoice(&player1, &player2)

	grid := InitGrid()
	GameLoop(player1, player2, grid)
}