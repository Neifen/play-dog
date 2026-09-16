package main

import (
	"math/rand/v2"

	"github.com/google/uuid"
)

type Game struct {
	players   []*GamePlayer
	turn      int
	round     *GameRound
	gameBoard *GameBoard

	playersMap map[uuid.UUID]*GamePlayer
}

func NewGame(setups []*PlayerSetup) *Game {
	playerAmount := len(setups)
	players := make([]*GamePlayer, playerAmount)
	playersMap := make(map[uuid.UUID]*GamePlayer)
	for _, setup := range setups {
		player := &GamePlayer{
			id:      setup.ID,
			partner: setup.Partner,
			name:    setup.Name,
		}
		players[setup.Seat] = player
		playersMap[setup.ID] = player
	}

	gb := NewGameBoard(players)
	turn := rand.IntN(playerAmount)
	round := &GameRound{starts: players[turn], deckSize: 6}
	return &Game{
		gameBoard: gb,
		players:   players,
		turn:      turn,
		round:     round,
	}
}

type GameRound struct {
	starts   *GamePlayer
	deckSize int
}

type GamePlayer struct {
	id         uuid.UUID
	partner    uuid.UUID
	name       string
	deck       []*GameCard
	hasSwapped bool
	color      Color

	boardPlayer *Player
}

func (g *Game) Partner(p *GamePlayer) *GamePlayer {
	return g.playersMap[p.id]
}

func (g *Game) hasWon(p *GamePlayer) bool {
	return p.boardPlayer.isDone() && g.Partner(p).boardPlayer.isDone()
}

type GameCard struct {
	card Card
}
