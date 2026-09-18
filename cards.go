package main

import "fmt"

type Move interface {
	act() error
	position() *Position
}

type MoveSwap struct {
	self  *Marble
	other *Marble
}

// act implements [Move].
func (s MoveSwap) act() error {
	return s.self.swap(s.other)
}

func (s MoveSwap) position() *Position {
	return s.other.Position
}

type MoveSeven struct {
	marble      *Marble
	heaven      bool
	chosenSteps int
	pos         *Position
}

// act implements [Move].
func (s MoveSeven) act() error {
	return s.marble.move(s.chosenSteps, s.heaven, true)
}

func (s MoveSeven) position() *Position {
	return s.pos
}

type MoveSteps struct {
	marble *Marble
	heaven bool
	amount int
	pos    *Position
}

// act implements [Move].
func (s MoveSteps) act() error {
	return s.marble.move(s.amount, s.heaven, false)
}

func (s MoveSteps) position() *Position {
	return s.pos
}

type LeaveHome struct {
	marble *Marble
}

// act implements [Move].
func (l LeaveHome) act() error {
	return l.marble.goOut()
}

func (l LeaveHome) position() *Position {
	return l.marble.Position.NextPosition
}

type CardAction interface {
	LegalMoves(gb *GameBoard, m *Marble, c Card) []Move
	ApplyMove(chosen Move) error //todo: not sure this is needed
}

type jokerAction struct {
	finishingAllowed bool
	picked           bool
	cardPicked       Card
}

// ApplyMove implements [CardAction].
func (a jokerAction) ApplyMove(chosen Move) error {
	if !a.picked || a.cardPicked.Name == Joker.Name {
		return fmt.Errorf("no card picked")
	}
	return a.cardPicked.Action.ApplyMove(chosen)
}

// LegalMoves implements [CardAction].
func (a jokerAction) LegalMoves(gb *GameBoard, m *Marble, c Card) []Move {
	if !a.picked || a.cardPicked.Name == Joker.Name {
		return []Move{}
	}

	return a.cardPicked.LegalMoves(gb, m)
}

type swapAction struct {
}

// ApplyMove implements [CardAction].
func (_ swapAction) ApplyMove(chosen Move) error {
	return chosen.act()
}

// LegalMoves implements [CardAction].
func (_ swapAction) LegalMoves(gb *GameBoard, m *Marble, c Card) []Move {
	var moves []Move
	if !m.canSwap() {
		return moves
	}

	for _, otherPlayer := range gb.Players {
		if m.Player == otherPlayer {
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
func (a splitAction) LegalMoves(_ *GameBoard, m *Marble, c Card) []Move {
	var moves []Move
	if m.Position.PositionType == Home && !c.leaveHome {
		return moves // can't leave home without a leavehome card
	}

	for i := range a.stepsLeft {
		steps := i + 1
		canMoveHeaven, posHeaven := m.canMove(steps, true)
		canMove, pos := m.canMove(steps, false)
		if !canMove && !canMoveHeaven {
			continue // no need to figure out `canFill` -> expensive
		}

		left := a.stepsLeft - steps
		canFill := maxMoves(m.Player.Marbles, m, left) >= left
		if canFill && canMoveHeaven {
			moves = append(moves, MoveSeven{marble: m, chosenSteps: steps, heaven: true, pos: posHeaven})
		}

		if canFill && canMove {
			moves = append(moves, MoveSeven{marble: m, chosenSteps: steps, heaven: false, pos: pos})
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
func (_ moveAction) LegalMoves(_ *GameBoard, m *Marble, card Card) []Move {
	var moves []Move
	if card.leaveHome && m.canOut() {
		moves = append(moves, LeaveHome{marble: m})
		return moves // cannot move more with this
	} else if m.Position.PositionType == Home && !card.leaveHome {
		return moves // can't leave home without a leavehome card
	}

	for _, places := range card.Moves {
		canMoveHeaven, posHeaven := m.canMove(places, true)
		if canMoveHeaven {
			moves = append(moves, MoveSteps{marble: m, amount: places, heaven: true, pos: posHeaven})
		}

		canMove, pos := m.canMove(places, false)
		if canMove {
			moves = append(moves, MoveSteps{marble: m, amount: places, heaven: false, pos: pos})
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

func (c Card) LegalMoves(gb *GameBoard, m *Marble) []Move {
	return c.Action.LegalMoves(gb, m, c)
}

func (c Card) ApplyMove(move Move) error {
	return c.Action.ApplyMove(move)
}

func (c Card) JockerPickCard(picked Card) Card {
	jokerAction, ok := c.Action.(jokerAction)
	if !ok {
		return c
	}
	jokerAction.picked = true
	jokerAction.cardPicked = picked
	c.Action = jokerAction

	return c
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
