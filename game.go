package main

import (
	"fmt"
	"math/rand/v2"
	"play-dog/util"
	"slices"
)

type Game struct {
	players   []*GamePlayer
	turn      int
	round     *GameRound
	gameBoard *GameBoard
}

func NewGame() *Game {
	gb := NewGameBoard()
	return &Game{gameBoard: gb}
}

func (g *Game) Ready() bool {
	//todo: unsure if gb.Ready is really needed
	nr := len(g.players)

	return nr == 4 || nr == 6
}

func (g *Game) Start() error {
	g.turn = rand.IntN(len(g.players))
	g.round = &GameRound{
		starts:   g.players[g.turn],
		deckSize: 6,
	}

	err := g.gameBoard.Start()
	if err != nil {
		return fmt.Errorf("start game: %w", err)
	}

	partnerless := []*GamePlayer{}
	for _, player := range g.players {
		if player.partner == nil {
			partnerless = append(partnerless, player)
		}
	}

	return nil
}

func (g *Game) ChooseColor(player *GamePlayer, color Color) error {
	return g.gameBoard.ChooseColor(&player.boardPlayer, color)
}

func (g *Game) ChoosePartner(one, two *GamePlayer) error {
	if one.partner != nil {
		return fmt.Errorf("player %s, already has a partner", one)
	}

	if two.partner != nil {
		return fmt.Errorf("player %s, already has a partner", two)
	}

	one.partner = two
	two.partner = one

	//todo: potentially better fit in game than in boardgame
	{
		one.boardPlayer.Partner = &two.boardPlayer
		two.boardPlayer.Partner = &one.boardPlayer
	}
	return nil
}

func (g *Game) orderAndLink() {
	playerAmount := len(g.players)
	unsortedPlayers := g.players
	sortedPlayers := make([]*GamePlayer, playerAmount)

	index := 0
	for len(unsortedPlayers) > 0 {
		if index >= 3 {
			panic(fmt.Sprintf("something went wrong with ordering players, half of players should not excede 3, was %d", index))
		}
		// Player
		player, playerIndex := util.RandItem(unsortedPlayers)
		sortedPlayers[index] = player
		unsortedPlayers = slices.Delete(unsortedPlayers, playerIndex, playerIndex+1)

		// Partner
		opositeIndex := (playerAmount / 2) + index
		partner := player.partner
		sortedPlayers[opositeIndex] = partner
		unsortedPlayers = slices.DeleteFunc(unsortedPlayers, func(p *GamePlayer) bool {
			if p == nil {
				return false
			}
			return p.ID == partner.ID
		})

		// 0, 1, 2 (max)
		index++
	}

	for i, s := range sortedSections {
		nextSectionIndex := (i + 1) % playerAmount
		linkPositions(s.Last, sortedSections[nextSectionIndex].First)
	}

	gb.Players = sortedPlayers
	gb.Sections = sortedSections
}

func (g *Game) Join(name string) (*GamePlayer, error) {
	boardPlayer, err := g.gameBoard.Join(name)
	if err != nil {
		return nil, fmt.Errorf("joining game: %w", err)
	}

	gamePlayer := &GamePlayer{
		name:        name,
		boardPlayer: *boardPlayer,
	}

	g.players = append(g.players, gamePlayer)
	return gamePlayer, nil
}

type GameRound struct {
	starts   *GamePlayer
	deckSize int
}

type GamePlayer struct {
	partner     *GamePlayer
	name        string
	deck        []*GameCard
	hasSwapped  bool
	boardPlayer Player
}

func NewGamePlayer()

type GameCard struct {
	card Card
}
