package main

// todo: test card_test when done, being able to play with partners card
type Game struct {
	players []*GamePlayer
	turn    int
	round   *GameRound
}

type GameRound struct {
	starts   *GamePlayer
	deckSize int
}

type GamePlayer struct {
	deck       []*GameCard
	hasSwapped bool
}

type GameCard struct {
	card Card
}
