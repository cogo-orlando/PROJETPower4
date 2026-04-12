package server

import (
	"html/template"
	"net/http"
	"strconv"
)

// ───── TEMPLATE ─────
func renderTemplate(w http.ResponseWriter, file string, data interface{}) {
	tmpl, err := template.ParseFiles("web/html/" + file)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmpl.Execute(w, data)
}

// ───── VARIABLES ─────
var (
	grid          [][]int
	currentPlayer = 1
	winner        int
	player1       Player1
	player2       Player2
)

// ───── HANDLERS ─────

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", nil)
}

func developperHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "develop.html", nil)
}

func rulesHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "rules.html", nil)
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "start.html", nil)
		return
	}

	name1 := r.FormValue("player1")
	name2 := r.FormValue("player2")

	player1, player2 = InitPlayers(name1, name2)
	TokenChoice(&player1, &player2)

	grid = InitGrid()
	currentPlayer = 1
	winner = 0

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	var message string

	if grid == nil {
		http.Redirect(w, r, "/start", http.StatusSeeOther)
		return
	}

	colStr := r.FormValue("column")

	if colStr != "" {
		col, err := strconv.Atoi(colStr)
		if err != nil || col < 0 || col >= 7 {
			http.Error(w, "colonne invalide", http.StatusBadRequest)
			return
		}
		playMove(col)
	}

	data := struct {
		Grid    [][]int
		Player  int
		Cols    []int
		Winner  int
		Player1 Player1
		Player2 Player2
		Message string
	}{
		Grid:    grid,
		Player:  currentPlayer,
		Cols:    []int{0, 1, 2, 3, 4, 5, 6},
		Winner:  winner,
		Player1: player1,
		Player2: player2,
		Message: message,
	}

	renderTemplate(w, "jeu.html", data)

	if winner != 0 {
		grid = InitGrid()
		currentPlayer = 1
		winner = 0
	}
}
