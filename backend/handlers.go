package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/segmentio/ksuid"
)

func greet(w http.ResponseWriter, request *http.Request) {
	setupResponse(&w, request)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	fmt.Print("New request...")

	response := make(map[string]string)
	response["message"] = "Hello dear user!"
	jsonByte, _ := json.Marshal(response)
	w.Write(jsonByte)
}

func regular_fetch(w http.ResponseWriter, request *http.Request){
	setupResponse(&w, request);
	request_dict := load_request(request)
	game_id, _ := strconv.Atoi(request_dict["game_id"])

	if(game_id > last_id || game_id == -1){
		response := make(map[string]string)
		response["success"] = "true"
		jsonByte, _ := json.Marshal(response)
		w.Write(jsonByte)
	}

	game := &games[game_id]

	if !game.Started {
		check_player_ready(game, request_dict)
	}

	for i := 0; i < game.Pi; i++ {
		if game.pawns_in_house[i] == 4{
			game.ended = true;
			game.winner = game.Pi;
		}
		if game.guids[i] == request_dict["key"]{
			game.Last_active[i] = time.Now()
		}
		time_diff := time.Now().Sub(game.Last_active[i])
		if time_diff.Seconds() >= 20{
			game.Player_status[i] = AFK
		}
	}


	time_since_player_turn := time.Now().Sub(game.Roll_time_start).Seconds()


	if time_since_player_turn >= 30 {
		fmt.Println("Czas minal!")
		next_player_turn(game)
	}

	if game.guids[game.Turn] == request_dict["key"]  { 
		player_status, _ := strconv.Atoi(request_dict["player_status"])
		if player_status == DICE_ROLLED {
			fmt.Println("player has rolled the dice")
			seed := rand.NewSource(time.Now().UnixNano())
			random := rand.New(seed)
			dice_roll := random.Intn(5) + 1
			game.Roll = dice_roll
			game.Player_status[game.Turn] = CHOOSING_PAWN
		} else if player_status == PAWN_CHOSEN {
			position_chosen, _ := strconv.Atoi(request_dict["chosen"])
			fmt.Printf("Wybrano %v", position_chosen)
			position := &game.Positions[game.Turn][position_chosen]
			fmt.Println(game.Positions)
			if *position <= 15{
				if game.Roll == 1 || game.Roll == 6 {
					player := game.Turn 
					*position = initial_position_of(player)
					check_pawns(game, *position)
				}
			} else {
				*position = (*position + game.Roll) % (40 + 15);
				if *position <= 15{
					*position += 15
				}
				check_pawns(game, *position)
				player := game.Turn
				game.moves[player][position_chosen] += game.Roll
				if game.moves[player][position_chosen] >= 39{
					*position = first_free_house(game, player)
					game.pawns_in_house[player]++
				}
			}
			fmt.Println(game.Positions)
			next_player_turn(game)
		}
	}
	
	if !game.Started && all_players_READY(game) {
		fmt.Println("game has started!")
		game.Started = true;
		game.Player_status[game.Turn] = ROLLING_DICE
	}

	response, _ := json.Marshal(game)
	w.Write(response)
}


func register_player(w http.ResponseWriter, request *http.Request) {
	setupResponse(&w, request)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	request_dict := load_request(request)

	response := make(map[string]string)

	game := &games[last_id]
	game.Names[game.Pi] = request_dict["name"]
	log.Print(game.Names)
	game.Player_status[game.Pi] = NOT_READY
	game.guids[game.Pi] = ksuid.New().String()
	log.Print(game.Names)
	response["success"] = "true"
	response["player_name"] = request_dict["name"]
	response["player_number"] = fmt.Sprint(game.Pi)
	response["key"] = game.guids[game.Pi]

	game.Pi++

	if(game.Pi == 4){
		log.Print("Creating new room...")
		last_id++
		games = append(games, new_game())
	}

	response["game_id"] = strconv.Itoa(last_id)

	jsonByte, _ := json.Marshal(response);
	w.Write(jsonByte)
}