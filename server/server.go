package server

import (
	"fmt"
	"net/http"
	"os"
)

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback en local
	}

	// fichiers statiques
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("web/img"))))

	// routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/page4", developperHandler)
	http.HandleFunc("/rules", rulesHandler)

	fmt.Println("Serveur lancé sur http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
