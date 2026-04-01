package main

import (
	"fmt"
	"html/template"
	"net/http"
	"power4/gamelogic"
	"strconv"
)

var (
	grid          [][]int
	currentPlayer = 1
	winner        int
	player1       gamelogic.Player1
	player2       gamelogic.Player2
)

// Menu page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("home.html"))
	tmpl.Execute(w, nil)
}

// Page to take names
func startHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("start.html"))
		tmpl.Execute(w, nil)
		return
	}

	//POST : take the names
	name1 := r.FormValue("player1")
	name2 := r.FormValue("player2")

	//initializes the two players
	player1, player2 = gamelogic.InitPlayers(name1, name2)
	gamelogic.TokenChoice(&player1, &player2)

	//initializes the game
	grid = gamelogic.InitGrid()
	currentPlayer = 1
	winner = 0

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

// Game page
func gameHandler(w http.ResponseWriter, r *http.Request) {
	var message string

	if grid == nil {
		grid = gamelogic.InitGrid()
	}

	// if a player clicks on a column
	if r.Method == http.MethodPost && winner == 0 {
		col, _ := strconv.Atoi(r.FormValue("column"))
		if col >= 0 && col < 7 {
			playMove(col)
		}
	}

	// data for the template
	data := struct {
		Grid    [][]int
		Player  int
		Cols    []int
		Winner  int
		Player1 gamelogic.Player1
		Player2 gamelogic.Player2
		message string
	}{
		Grid:    grid,
		Player:  currentPlayer,
		Cols:    []int{0, 1, 2, 3, 4, 5, 6},
		Winner:  winner,
		Player1: player1,
		Player2: player2,
		message: message,
	}

	tmpl := template.Must(template.ParseFiles("jeu.html"))
	tmpl.Execute(w, data)

	// If victory, reinitialize grid and game
	if winner != 0 {
		grid = gamelogic.InitGrid()
		currentPlayer = 1
		winner = 0
	}
}

// Other pages

func developperHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("develop.html"))
	tmpl.Execute(w, nil)
}

func rulesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("rules.html"))
	tmpl.Execute(w, nil)
}

// GAME LOGIC
func playMove(columnChoice int) [][]int {
	var valid bool
	var message string

	if currentPlayer == 1 {
		grid, valid, message = gamelogic.PutToken1(grid, player1, columnChoice)
	} else {
		grid, valid, message = gamelogic.PutToken2(grid, player2, columnChoice)
	}

	if !valid {
		fmt.Println(message)
		return grid
	}

	if gamelogic.CheckWin(grid) {
		winner = currentPlayer
		fmt.Println(message)
		return grid
	}

	// changing turn
	if currentPlayer == 1 {
		currentPlayer = 2
	} else {
		currentPlayer = 1
	}
	fmt.Println(message)
	return grid
}

func getCurrentPlayerName() string {
	if currentPlayer == 1 {
		return player1.Name
	}
	return player2.Name
}

// Server launch
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", homeHandler)            // Page d'accueil
	http.HandleFunc("/start", startHandler)      // Page de saisie des noms
	http.HandleFunc("/game", gameHandler)        // Page du jeu
	http.HandleFunc("/page4", developperHandler) // Page développeur
	http.HandleFunc("/rules", rulesHandler)      // Rules page

	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
