package main

import (
	"bufio"
	"fmt"
	"os"
	"power4/gamelogic"
	"power4/utils"
	"strings"
)

var (
	player1Name string
	player2Name string
)

func main() {
	var name1 string
	var name2 string
	reader := bufio.NewReader(os.Stdin) // reads the input
	fmt.Println("Player 1, what's your name?")
	for {
		name1, _ = reader.ReadString('\n')
		name1 = strings.TrimSpace(name1)

		if utils.IsAlpha(name1) {
			break
		} else {
			fmt.Println("Your name must be composed of letters only")
		}
	}
	fmt.Println("Player 2, what's your name?")
	for {
		name2, _ = reader.ReadString('\n')
		name2 = strings.TrimSpace(name2)

		if utils.IsAlpha(name2) {
			break
		} else {
			fmt.Println("Your name must be composed of letters only")
		}
	}
	var choiceMenu string
	fmt.Println("Begin?")
	fmt.Println("1. START")
	fmt.Println("2. QUIT")
	fmt.Scanln(&choiceMenu)
	player1, player2 := gamelogic.InitPlayers(name1, name2)
	switch choiceMenu {
	case "1":
		gamelogic.TokenChoice(&player1, &player2)
	case "2":
		return
	default:
		fmt.Println("Invalid choice")

	}
	player1, player2 = gamelogic.InitPlayers(name1, name2)
	grid := gamelogic.InitGrid()
	gamelogic.GameLoop(player1, player2, grid)
}
