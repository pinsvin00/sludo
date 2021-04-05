package main

import "time"

const (
	red    = iota
	blue   = iota
	yellow = iota
	green  = iota
)
const pawn_destroyed = -1

const (
	NOT_READY     = iota
	READY         = iota
	ROLLING_DICE  = iota
	DICE_ROLLED =   iota
	CHOOSING_PAWN = iota
	PAWN_CHOSEN = iota
	AFK           = iota
)

type Game struct {
	Id             int       `json:"id"`
	Started        bool      `json:"started"`
	Positions      [4][4]int `json:"positions"`
	pawns_in_house [4]int
	guids          [4]string
	moves          [4][4]int
	last_active [4]time.Time
	Names         [4]string `json:"names"`
	Player_status [4]int    `json:"players_status"`
	Pi            int       `json:"player_count"`
	Turn          int       `json:"turn"`
	Roll          int       `json:"roll"`
}

type RegisterRequest struct {
	name string
}

func new_game() Game {
	game := Game{
		Pi:   0,
		Turn: 0,
	}
	pos := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			game.Positions[i][j] = pos
			game.moves[i][j] = 0
			pos++
		}
	}

	return game
}