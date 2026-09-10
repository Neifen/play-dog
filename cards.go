package main

import "fmt"

type Move interface {
	act() error
}

type MoveSwap struct {
	self  *Marble
	other *Marble
}

// act implements [Move].
func (s MoveSwap) act() error {
	return s.self.swap(s.other)
}

type MoveSeven struct {
	marble      *Marble
	heaven      bool
	chosenSteps int
}

// act implements [Move].
func (s MoveSeven) act() error {
	return s.marble.move(s.chosenSteps, s.heaven, true)
}

type MoveSteps struct {
	marble *Marble
	heaven bool
	amount int
}

// act implements [Move].
func (s MoveSteps) act() error {
	return s.marble.move(s.amount, s.heaven, false)
}

type LeaveHome struct {
	marble *Marble
}

// act implements [Move].
func (l LeaveHome) act() error {
	return l.marble.goOut()
}

type CardAction interface {
	LegalMoves(gb *GameBoard, p *Player, m *Marble, c *Card) []Move
	ApplyMove(chosen Move) error //todo: not sure this is needed
}

// var actions = map[CardType]CardAction{
// 	TypeMover:    moveAction{},
// 	TypeSplitter: splitAction{stepsLeft: 7},
// 	TypeJoker:    jokerAction{finishingAllowed: true},
// 	TypeSwap:     swapAction{},
// }

type jokerAction struct {
	finishingAllowed bool
	picked           bool
	cardPicked       Card
}

// ApplyMove implements [CardAction].
func (a jokerAction) ApplyMove(chosen Move) error {
	return a.cardPicked.Action.ApplyMove(chosen)
}

// LegalMoves implements [CardAction].
func (a jokerAction) LegalMoves(gb *GameBoard, p *Player, m *Marble, c *Card) []Move {
	if !a.picked || a.cardPicked.Name == c.Name {
		return []Move{}
	}

	return a.cardPicked.Action.LegalMoves(gb, p, m, c)
}

type swapAction struct {
}

// ApplyMove implements [CardAction].
func (_ swapAction) ApplyMove(chosen Move) error {
	return chosen.act()
}

// LegalMoves implements [CardAction].
func (_ swapAction) LegalMoves(gb *GameBoard, p *Player, m *Marble, c *Card) []Move {
	var moves []Move
	if !m.canSwap() {
		return moves
	}

	for _, otherPlayer := range gb.Players {
		if p == otherPlayer {
			continue
		}

		for _, otherMarble := range otherPlayer.Marbles {
			if m.canSwapWith(otherMarble) {
				moves = append(moves, MoveSwap{self: m, other: otherMarble})
			}
		}
	}

	return moves
}

type splitAction struct {
	stepsLeft int
}

// ApplyMove implements [CardAction].
func (a splitAction) ApplyMove(chosen Move) error {
	moveSeven, ok := chosen.(MoveSeven)
	if !ok {
		return fmt.Errorf("splitAction needs to use a moveSplit")
	}
	a.stepsLeft -= moveSeven.chosenSteps
	return chosen.act()
}

func maxMoves(marbles Marbles, except *Marble, cutoff int) int {
	max := 0
	for _, m := range marbles {
		if m == except {
			continue
		}

		max += m.max(cutoff - max)
		if max >= cutoff {
			return max
		}
	}
	return max
}

// LegalMoves implements [CardAction].
func (a splitAction) LegalMoves(_ *GameBoard, p *Player, m *Marble, c *Card) []Move {
	var moves []Move
	for i := range a.stepsLeft {
		steps := i + 1
		canMoveHeaven := m.canMove(steps, true)
		canMove := m.canMove(steps, true)
		if !canMove && !canMoveHeaven {
			continue // no need to figure out `canFill` -> expensive
		}

		left := a.stepsLeft - steps
		canFill := maxMoves(p.Marbles, m, left) >= left
		if canFill && canMoveHeaven {
			moves = append(moves, MoveSeven{marble: m, chosenSteps: steps, heaven: true})
		}

		if canFill && canMove {
			moves = append(moves, MoveSeven{marble: m, chosenSteps: steps, heaven: false})
		}
	}
	return moves
}

type moveAction struct {
}

// ApplyMove implements [CardAction].
func (_ moveAction) ApplyMove(chosen Move) error {
	return chosen.act()
}

// LegalMoves implements [CardAction].
func (_ moveAction) LegalMoves(_ *GameBoard, p *Player, m *Marble, card *Card) []Move {
	var moves []Move
	if card.leaveHome && m.canOut() {
		moves = append(moves, LeaveHome{marble: m})
		return moves // cannot move more with this
	}

	for _, places := range card.Moves {
		if m.canMove(places, true) {
			moves = append(moves, MoveSteps{marble: m, amount: places, heaven: true})
		}

		if m.canMove(places, false) {
			moves = append(moves, MoveSteps{marble: m, amount: places, heaven: false})
		}
	}
	return moves
}

type Card struct {
	Name      string
	Moves     []int
	Action    CardAction
	leaveHome bool
}

var (
	Joker = Card{
		Name:   "Joker",
		Moves:  []int{},
		Action: jokerAction{finishingAllowed: true}, // static for now, maybe a setting later
	}

	Ace = Card{
		Name:      "Ace",
		Moves:     []int{1, 11},
		leaveHome: true,
		Action:    moveAction{},
	}

	Two = Card{
		Name:   "Two",
		Moves:  []int{2},
		Action: moveAction{},
	}

	Three = Card{
		Name:   "Three",
		Moves:  []int{3},
		Action: moveAction{},
	}

	Four = Card{
		Name:   "Four",
		Moves:  []int{-4, 4},
		Action: moveAction{},
	}

	Five = Card{
		Name:   "Five",
		Moves:  []int{5},
		Action: moveAction{},
	}

	Six = Card{
		Name:   "Six",
		Moves:  []int{6},
		Action: moveAction{},
	}

	Seven = Card{
		Name:   "Seven",
		Moves:  []int{7},
		Action: splitAction{7},
	}

	Eight = Card{
		Name:   "Eight",
		Moves:  []int{8},
		Action: moveAction{},
	}

	Nine = Card{
		Name:   "Nine",
		Moves:  []int{9},
		Action: moveAction{},
	}

	Ten = Card{
		Name:   "Ten",
		Moves:  []int{10},
		Action: moveAction{},
	}

	Jack = Card{
		Name:   "Jack",
		Moves:  []int{},
		Action: swapAction{},
	}

	Queen = Card{
		Name:   "Queen",
		Moves:  []int{12},
		Action: moveAction{},
	}

	King = Card{
		Name:   "King",
		Moves:  []int{13},
		Action: moveAction{},
	}
)
