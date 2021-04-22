package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)


func all_players_READY(game * Game) bool{
	fmt.Println(game.Player_status)
	if game.Pi == 1 {
		return false
	}
	for i := 0; i < game.Pi; i++ {
		if game.Player_status[i] == NOT_READY{
			return false
		}
	}
	return true
}

func initial_position_of(player int) int{
	return 16 + player * 10
}

func first_free_field(game * Game,player int) int {
	bottom := player*4
	free := [...]bool{true, true, true, true}
	cap := 15
	for i := 0 ; i < 4; i++ {
		position := game.Positions[player][i]
		if position <= cap{
			position -= bottom
			free[position] = false
		}
	}
	for i := 0; i < 4; i++ {
		if free[i]{
			return bottom + i
		}
	}
	return 0
}

func check_pawns(game * Game, position int){
	player := game.Turn
	for i := 0; i < 4; i++ {
		if i == player{
			continue
		}
		for j := 0; j < 4; j++ {
			if game.Positions[i][j] == position {
				fmt.Printf("Collision between players %v and %v", player, i);
				fmt.Printf("Position of desroyed pawn %v", position);
				game.Positions[i][j] = first_free_field(game ,i)
				game.moves[i][j] = 0
			}
		}
	}
}

func check_passings(last_position int, position int, player int) bool{
	initial := initial_position_of(player) // 16, 26, 36, 46
	if last_position > initial && position <= initial{ //jesli gosc przeszedl initial nie dziala dla przebitki
		return true
	}
	return false
}

func first_free_house(game * Game,player int) int {
	bottom := player * 4 + (56) + game.pawns_in_house[player]
	return bottom
}

func check_player_ready(game * Game, request_dict map[string]string){
	for i := 0 ; i < game.Pi; i++ {
		if game.guids[i] == request_dict["key"]{
			game.Player_status[i], _ = strconv.Atoi(request_dict["player_status"])
			log.Print(game.Player_status[i])
		}
	}
}

func check_if_player_in_game(game * Game, request_dict map[string]string) bool{
	result := false
	for i := 0 ; i < game.Pi; i++ {
		if game.guids[i] == request_dict["key"]{
			result = true
		}
	}
	return result
}

func save_game_state(game * Game){
	fmt.Println(game)
}

func next_player_turn(game * Game){
	game.Player_status[game.Turn] = 1
	game.Turn = (game.Turn + 1) % game.Pi
	game.Roll_time_start = time.Now() 
	game.Player_status[game.Turn] = ROLLING_DICE
}


func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func log_bytes(bytes []byte){
	sb := string(bytes)
	log.Printf(sb)
}

func load_request(request * http.Request)map[string]string {
	//fmt.Println("New request...")
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	//log_bytes(body)
	request_dict := make(map[string]string)
	json.Unmarshal(body, &request_dict)
	return request_dict
}