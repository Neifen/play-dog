package main

import (
	"math/rand/v2"

	"github.com/google/uuid"
)

func (g *Game) NextRound() {
	cards := shuffle(g.deck)
	handoutCards(round, players, cards)
}

	handoutCards(round, players, cards)


func createDeck() []*ActualCard {
	deck := make([]*ActualCard, 0, 108) // ?
	normalCards := 8
	jokers := 4

	for range normalCards {
		deck = append(deck, &ActualCard{Ace})
		deck = append(deck, &ActualCard{Two})
		deck = append(deck, &ActualCard{Three})
		deck = append(deck, &ActualCard{Four})
		deck = append(deck, &ActualCard{Five})
		deck = append(deck, &ActualCard{Six})
		deck = append(deck, &ActualCard{Seven})
		deck = append(deck, &ActualCard{Eight})
		deck = append(deck, &ActualCard{Nine})
		deck = append(deck, &ActualCard{Ten})
		deck = append(deck, &ActualCard{Jack})
		deck = append(deck, &ActualCard{Queen})
		deck = append(deck, &ActualCard{King})
	}

	for range jokers {
		deck = append(deck, &ActualCard{Joker})
	}

	return deck
}

func shuffle(deck []*ActualCard) []*ActualCard {
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}

type ActualCard struct {
	Card Card
}

type Game struct {
	deck      []*ActualCard
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
			color:   setup.Color,
		}
		players[setup.Seat] = player
		playersMap[setup.ID] = player
	}

	gb := NewGameBoard(players)
	turn := rand.IntN(playerAmount)
	round := &GameRound{starts: players[turn], deckSize: 6}

	deck := createDeck()
	return &Game{
		gameBoard:  gb,
		players:    players,
		playersMap: playersMap,
		turn:       turn,
		round:      round,
		deck:       deck,
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
