package main

import (
	"fmt"
	"math/rand/v2"
	"slices"

	"github.com/google/uuid"
)

func (g *Game) NextRound() {
	deck := shuffle(g.deck)
	g.handoutCards(deck)
	g.turn = g.round.startsTurn
	for _, player := range g.players {
		player.hasSwapped = false
	}

	// prepare for the round after
	g.round.deckSize--
	if g.round.deckSize == 1 {
		g.round.deckSize = 6
	}
	g.round.startsTurn = (g.round.startsTurn + 1) % len(g.players)
}

func (g *Game) NextTurn() {
	done := true
	for _, player := range g.players {
		if len(player.deck) != 0 {
			done = false
			break
		}
	}

	if done {
		g.NextRound()
		return
	}

	g.turn = (g.turn + 1) % len(g.players)
}

func (g *Game) HaveAllSwapped() bool {
	for _, player := range g.players {
		if !player.hasSwapped {
			return false
		}
	}
	return true
}

func (g *Game) Play(player *GamePlayer, card *GameCard, move Move) error {
	turnID := g.players[g.turn].id
	if player.id != turnID {
		return fmt.Errorf("its not %s's turn (%s)", player.name, player.id)
	}

	if !g.HaveAllSwapped() {
		return fmt.Errorf("Not everyone has swapped")
	}

	err := card.card.ApplyMove(move)
	if err != nil {
		return fmt.Errorf("game: error applying move: %w", err)
	}

	if card.isPlayed() {
		player.deck = slices.DeleteFunc(player.deck, func(c *GameCard) bool { return c.id == card.id })
		card.reset()
		g.NextTurn()
	}

	return err
}

func (g *Game) LegalMoves(player *GamePlayer, card *GameCard, m *Marble) ([]Move, error) {
	turnID := g.players[g.turn].id
	if player.id != turnID {
		return nil, fmt.Errorf("its not %s's turn (%s)", player.name, player.id)
	}

	if m.Player != player.boardPlayer ||
		player.boardPlayer.isDone() && m.Player != g.playersMap[player.partner].boardPlayer {
		return nil, fmt.Errorf("Player %s can not move %s's marbles", player.name, m.Player.Name)
	}
	// not checking for swaps, maybe player wants to check legal moves before swap

	return card.card.LegalMoves(g.gameBoard, m), nil
}

func (g *Game) handoutCards(cards []*GameCard) {
	handoutNr := len(g.players) * g.round.deckSize

	player := 0
	for i := range handoutNr {
		g.players[player].deck = append(g.players[0].deck, cards[i])
		player = (player + 1) % len(g.players)
	}
}

func createDeck() []*GameCard {
	deck := make([]*GameCard, 0, 108) // ?
	normalCards := 8
	jokers := 4

	for range normalCards {
		deck = append(deck, &GameCard{uuid.New(), Ace})
		deck = append(deck, &GameCard{uuid.New(), Two})
		deck = append(deck, &GameCard{uuid.New(), Three})
		deck = append(deck, &GameCard{uuid.New(), Four})
		deck = append(deck, &GameCard{uuid.New(), Five})
		deck = append(deck, &GameCard{uuid.New(), Six})
		deck = append(deck, &GameCard{uuid.New(), Seven})
		deck = append(deck, &GameCard{uuid.New(), Eight})
		deck = append(deck, &GameCard{uuid.New(), Nine})
		deck = append(deck, &GameCard{uuid.New(), Ten})
		deck = append(deck, &GameCard{uuid.New(), Jack})
		deck = append(deck, &GameCard{uuid.New(), Queen})
		deck = append(deck, &GameCard{uuid.New(), King})
	}

	for range jokers {
		deck = append(deck, &GameCard{uuid.New(), Joker})
	}

	return deck
}

func shuffle(deck []*GameCard) []*GameCard {
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}

type Game struct {
	deck      []*GameCard
	players   []*GamePlayer
	turn      int
	round     *GameRound
	gameBoard *GameBoard
	swaps     map[uuid.UUID]*GameCard

	playersMap map[uuid.UUID]*GamePlayer
}

func NewGame(setups []*PlayerSetup) *Game {
	playerAmount := len(setups)
	players := make([]*GamePlayer, playerAmount)
	playersMap := make(map[uuid.UUID]*GamePlayer)
	swaps := make(map[uuid.UUID]*GameCard)
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
	round := &GameRound{startsTurn: turn, deckSize: 6}

	deck := createDeck()
	return &Game{
		gameBoard:  gb,
		players:    players,
		playersMap: playersMap,
		turn:       turn,
		round:      round,
		deck:       deck,
		swaps:      swaps,
	}
}

type GameRound struct {
	startsTurn int
	deckSize   int
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
	id   uuid.UUID
	card Card
}

func (c GameCard) isPlayed() bool {
	seven, isSplit := c.card.Action.(splitAction)

	if !isSplit {
		return true
	}

	return seven.stepsLeft == 0
}

func (c *GameCard) reset() {
	seven, isSplit := c.card.Action.(splitAction)

	if !isSplit {
		return
	}

	seven.stepsLeft = 7
	c.card.Action = seven
}
