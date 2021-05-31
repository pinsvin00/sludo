package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var games[] Game
var last_id = 0


func main() {
	game := new_game()
	games = append(games, game)

	fmt.Println("Initiated server...");
	r := mux.NewRouter()

	game.Moves[0][0] = 38
	game.Positions[0][0] = 55


	r.HandleFunc("/api/greet_user", greet).Methods("GET");
	r.HandleFunc("/api/join", register_player).Methods(http.MethodPost);
	r.HandleFunc("/api/process", regular_fetch).Methods("POST");
	r.Use(mux.CORSMethodMiddleware(r))
	log.Fatal(http.ListenAndServe(":9000", r))
}
