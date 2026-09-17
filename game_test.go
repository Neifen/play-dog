package main

import (
	"testing"

	"github.com/google/uuid"
)

func Test_GameSetup(t *testing.T) {
	aID := uuid.New()
	bID := uuid.New()
	cID := uuid.New()
	dID := uuid.New()

	a := &PlayerSetup{
		ID:      aID,
		Name:    "a",
		Color:   White,
		Seat:    0,
		Partner: cID,
	}

	b := &PlayerSetup{
		ID:      bID,
		Name:    "b",
		Color:   Black,
		Seat:    1,
		Partner: dID,
	}

	c := &PlayerSetup{
		ID:      cID,
		Name:    "a",
		Color:   Yellow,
		Seat:    3,
		Partner: aID,
	}

	d := &PlayerSetup{
		ID:      dID,
		Name:    "d",
		Color:   Red,
		Seat:    2,
		Partner: bID,
	}
	game := NewGame([]*PlayerSetup{a, b, c, d})
	assertNotNil(t, game, "new game")
	assertEq(t, len(game.players), 4)
	assertEq(t, len(game.playersMap), 4)
	assertEq(t, game.round.starts.id, game.players[game.turn].id)
	assertEq(t, game.round.deckSize, 6)

	assertEq(t, game.players[0].id, aID)
	assertEq(t, game.players[1].id, bID)
	assertEq(t, game.players[2].id, dID)
	assertEq(t, game.players[3].id, cID)
}
