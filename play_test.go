package main

import (
	"fmt"
	"testing"
)

func setupGame(t *testing.T) *GameBoard {
	gb := NewGameBoard()
	gb.Join("a")
	gb.Join("b")
	gb.Join("c")
	gb.Join("d")
	err := gb.Start()
	if err != nil {
		t.Fatalf("failed to create gameboard with error %v", err)
	}

	return gb
}

func Test_goOut(t *testing.T) {
	gb := setupGame(t)
	players := gb.Players

	for pI, player := range players {
		// first Marble should be able to get out
		firstM := player.Marbles[0]
		assertTrue(t, firstM.canOut(), "First M 'canOut'")
		err := firstM.goOut()
		assertNil(t, err, "move marble out for p%d", pI)
		assertTrue(t, firstM.isBlocking(), "marble 'isBlocking'")

		// all subsequencial (and same marble) can't get out
		for mI, marble := range player.Marbles {
			assertFalse(t, marble.canOut(), "Player %d, Marble %d 'canOut'", pI, mI)
			err := firstM.goOut()
			assertNotNil(t, err, "move m%d out for p%d", mI, pI)
		}
	}
}

func Test_block(t *testing.T) {
	gb := setupGame(t)
	players := gb.Players
	var err error

	// first get out
	for i, player := range players {
		// first Marble should be able to get out
		marble := player.Marbles[0]
		assertTrue(t, marble.canOut(), "First M 'canOut'")
		err = marble.goOut()
		assertNil(t, err, "move marble out for p%d", i)

		assertTrue(t, marble.isBlocking(), "marble 'isBlocking'")
		assertEq(t, marble.StartTouches, 1)
		assertEq(t, marble.Position.PositionType, Start)
		assertEq(t, marble.Position.Section, marble.Player.Section)
	}

	// move 8 (y), then 8 (n) blocked
	for i, player := range players {
		marble := player.Marbles[0]
		assertFalse(t, marble.canMove(8, true), "marble for p%d 'canMove(8, true[heaven])", i) //heaven
		err := marble.move(8, true, false)
		assertNotNil(t, err, "move marble 8[heaven] for p%d", i) // because it can't go into heaven there before touching start twice

		assertTrue(t, marble.canMove(8, false), "marble for p%d 'canMove(8, false[heaven])", i)
		err = marble.move(8, false, false)
		assertNil(t, err, "move marble 8 for p%d", i)

		// moving into the next player (blocking)
		if i == 3 {
			continue // last one is not blocked, just skip
		}
		assertFalse(t, marble.canMove(8, false), "marble for p%d 'canMove(8, false[heaven])", i) //blocking
		err = marble.move(8, false, false)
		assertNotNil(t, err, "move marble 8 for p%d", i) // because its blocked
	}

	// move 8 again (y)
	for i, player := range players {
		marble := player.Marbles[0]
		assertTrue(t, marble.canMove(8, false), "marble for p%d 'canMove(8, false[heaven])", i)
		err = marble.move(8, false, false)
		assertNil(t, err, "move marble 8 for p%d", i)

		assertEq(t, marble.StartTouches, 1)
		assertEq(t, marble.Position.PositionType, Start)
		assertNeq(t, marble.Position.Section, marble.Player.Section)
		assertFalse(t, marble.isBlocking(), "'isBlocking'")
	}
}

func Test_quickHome(t *testing.T) {
	gb := setupGame(t)
	players := gb.Players
	var err error

	// first get out
	for i, player := range players {
		marble := player.Marbles[0]
		err = marble.goOut()
		assertNil(t, err, "gm1 o out p%d", i)

		assertFalse(t, marble.canMove(-4, true), "move m1 p%d back 4 [heaven]", i)
		err = marble.move(-4, true, false)
		assertNotNil(t, err, "move m1 p%d 4 back [heaven]", i)

		assertTrue(t, marble.canMove(-4, false), "move m1 p%d back 4", i)
		err = marble.move(-4, false, false)
		assertNil(t, err, "move m1 p%d 4 back", i)

		assertTrue(t, marble.canMove(7, true), "move m1 p%d 7 into heaven", i)
		err = marble.move(7, true, false)
		assertNil(t, err, "move m1 p%d 7 into heaven", i)

		assertEq(t, marble.Position.PositionType, Heaven)
		assertEq(t, marble.StartTouches, 2)
		assertTrue(t, marble.isBlocking(), "marble m1 p%d is blocking [heaven]", i)

		// second marble
		marble = player.Marbles[1]
		err = marble.goOut()
		assertNil(t, err, "m2 go out p%d", i)

		assertTrue(t, marble.canMove(-4, false), "move m2 p%d back 4", i)
		err = marble.move(-4, false, false)
		assertNil(t, err, "move m2 p%d 4 back", i)

		// third marble (block)
		marble = player.Marbles[2]
		err = marble.goOut()
		assertNil(t, err, "m3 go out p%d", i)

		// second marble (blocked)
		marble = player.Marbles[1]
		assertFalse(t, marble.canMove(6, true), "move m2 p%d 6 into heaven [blocked]", i)
		err = marble.move(6, true, false)
		assertNotNil(t, err, "move m2 p%d 6 into heaven [blocked]", i)

		assertFalse(t, marble.canMove(6, false), "move m2 p%d 4 [blocked]", i)
		err = marble.move(6, false, false)
		assertNotNil(t, err, "move m2 p%d 6 [blocked]", i)

		// third marble (move a bit)
		marble = player.Marbles[2]
		assertTrue(t, marble.canMove(2, false), "move m3 p%d 2 [unblocking]", i)
		err = marble.move(2, false, false)
		assertNil(t, err, "move m3 p%d 2 [unblocking]", i)

		// second marble (not blocked anymore, but heaven is partly blocked)
		marble = player.Marbles[1]
		assertFalse(t, marble.canMove(7, true), "move m2 p%d 7 into heaven [heaven-blocked]", i)
		err = marble.move(7, true, false)
		assertNotNil(t, err, "move m2 p%d 7 into heaven [heaven-blocked]", i)

		assertTrue(t, marble.canMove(6, false), "can move m2 p%d 6 [unblocked]", i)
		assertTrue(t, marble.canMove(6, true), "move m2 p%d 6 into heaven [unblocked]", i)
		err = marble.move(6, true, false)
		assertNil(t, err, "move m2 p%d 6 [not blocked]", i)

		assertFalse(t, player.isDone(), "player isDone")
		assertFalse(t, player.hasWon(), "player hasWon")

		// quickly reajust everyone
		if err = player.Marbles[0].move(1, true, false); err != nil {
			fmt.Println(err)
		}
		if err = player.Marbles[1].move(1, true, false); err != nil {
			fmt.Println(err)
		}
		if err = player.Marbles[2].move(-4, false, false); err != nil {
			fmt.Println(err)
		}
		if err = player.Marbles[2].move(4, true, false); err != nil {
			fmt.Println(err)
		}

		marble = player.Marbles[3]
		if err = marble.goOut(); err != nil {
			fmt.Println(err)
		}
		if err = marble.move(-4, false, false); err != nil {
			fmt.Println(err)
		}
		if err = marble.move(4, true, false); err != nil {
			fmt.Println(err)
		}

		assertFalse(t, player.isDone(), "player isDone")
		assertFalse(t, player.hasWon(), "player hasWon")

		assertTrue(t, marble.canMove(1, true), "can last marble go into heaven p%d", i)
		err = marble.move(1, true, false)
		assertNil(t, err, "last marble go into heaven p%d", i)

		assertTrue(t, player.isDone(), "player p%d isDone", i)

		if i < 2 {
			assertFalse(t, player.hasWon(), "player hasWon")
		} else {
			assertTrue(t, player.hasWon(), "player hasWon")
		}
	}
}

func Test_SendHome(t *testing.T) {
	gb := setupGame(t)
	// var err error

	a := gb.Players[0].Marbles[3]
	b := gb.Players[1].Marbles[1]
	c := gb.Players[1].Marbles[2]

	if err := a.goOut(); err != nil {
		fmt.Println(err)
	}
	if err := a.move(12, false, false); err != nil {
		fmt.Println(err)
	}
	if err := b.goOut(); err != nil {
		fmt.Println(err)
	}
	if err := b.move(-4, false, false); err != nil {
		fmt.Println(err)
	}

	assertEq(t, a.Position.PositionType, Home)
	assertEq(t, b.Position.PositionType, Regular)

	if err := a.goOut(); err != nil {
		fmt.Println(err)
	}
	if err := a.move(16, false, false); err != nil {
		fmt.Println(err)
	}
	assertEq(t, a.StartTouches, 1)

	if err := c.goOut(); err != nil {
		fmt.Println(err)
	}
	assertEq(t, a.Position.PositionType, Home)
	assertEq(t, b.Position.PositionType, Regular)
	assertEq(t, c.Position.PositionType, Start)

	if err := a.goOut(); err != nil {
		fmt.Println(err)
	}
	if err := c.move(1, false, false); err != nil {
		fmt.Println(err)
	}
	if err := a.move(15, false, false); err != nil {
		fmt.Println(err)
	}
	if err := b.move(6, false, true); err != nil {
		fmt.Println(err)
	}
	assertEq(t, a.Position.PositionType, Home)
	assertEq(t, b.Position.PositionType, Regular)
	assertEq(t, c.Position.PositionType, Home)
}
