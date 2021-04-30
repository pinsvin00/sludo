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
	fmt.Println(games[0].Turn)
	// dsn := "root:ziomek22@tcp(127.0.0.1:3306)/ludo?charset=utf8mb4&parseTime=True&loc=Local";
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// fmt.Print(db)
	// fmt.Print(err)

	fmt.Println("Initiated server...");
	r := mux.NewRouter()

	r.HandleFunc("/api/greet_user", greet).Methods("GET");
	r.HandleFunc("/api/join", register_player).Methods(http.MethodPost);
	r.HandleFunc("/api/process", regular_fetch).Methods("POST");
	r.Use(mux.CORSMethodMiddleware(r))
	log.Fatal(http.ListenAndServe(":9000", r))
}
