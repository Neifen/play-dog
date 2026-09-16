package main

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func setupGame() *GameBoard {
	aID := uuid.New()
	bID := uuid.New()
	cID := uuid.New()
	dID := uuid.New()

	a := &GamePlayer{
		id:      aID,
		partner: cID,
		name:    "a",
		color:   Blue,
	}
	b := &GamePlayer{
		id:      bID,
		partner: dID,
		name:    "b",
		color:   Green,
	}
	c := &GamePlayer{
		id:      cID,
		partner: aID,
		name:    "c",
		color:   Yellow,
	}
	d := &GamePlayer{
		id:      dID,
		partner: bID,
		name:    "d",
		color:   Black,
	}

	gb := NewGameBoard([]*GamePlayer{a, b, c, d})
	return gb
}

func Test_goOut(t *testing.T) {
	gb := setupGame()
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
	gb := setupGame()
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
	gb := setupGame()
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

		// trying to move marble into heaven but not actually
		assertFalse(t, marble.canMove(2, true), "cant move m1 p%d 2 into heaven, its not heaven there", i)
		assertNotNil(t, marble.move(2, true, false), "cant move m1 p%d 2 into heaven, its not heaven there", i)

		// trying to move marble past heaven
		assertFalse(t, marble.canMove(9, true), "cant move m1 p%d 9 into heaven", i)
		assertNotNil(t, marble.move(9, true, false), "cant move m1 p%d 9 into heaven", i)

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
		if err = marble.move(4, false, false); err != nil {
			fmt.Println(err)
		}

		assertFalse(t, player.isDone(), "player isDone")
		assertTrue(t, marble.canMove(1, true), "can last marble go into heaven p%d", i)
		err = marble.move(1, true, false)
		assertNil(t, err, "last marble go into heaven p%d", i)

		assertTrue(t, player.isDone(), "player p%d isDone", i)
	}
}

func Test_SendHome(t *testing.T) {
	gb := setupGame()
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

func Test_swap(t *testing.T) {
	gb := setupGame()
	// var err error

	selfA := gb.Players[1].Marbles[1]
	selfB := gb.Players[1].Marbles[2]
	selfColor := gb.Players[1].Color
	lockedA := gb.Players[0].Marbles[3]
	lockedB := gb.Players[0].Marbles[1]
	lockedC := gb.Players[0].Marbles[2]
	otherA := gb.Players[2].Marbles[3]
	otherB := gb.Players[3].Marbles[3]
	otherAColor := gb.Players[2].Color
	otherAPlayer := gb.Players[2]
	otherBColor := gb.Players[3].Color
	otherBPlayer := gb.Players[3]

	// prepare all marbles
	if err := selfA.goOut(); err != nil {
		fmt.Println(err)
	}
	if err := selfA.move(4, false, false); err != nil {
		fmt.Println(err)
	}

	if err := selfB.goOut(); err != nil {
		fmt.Println(err)
	}
	if err := selfB.move(2, false, false); err != nil {
		fmt.Println(err)
	}

	if err := lockedA.goOut(); err != nil {
		fmt.Println(err)
	}
	if err := lockedA.move(-4, false, false); err != nil {
		fmt.Println(err)
	}
	if err := lockedA.move(5, true, false); err != nil {
		fmt.Println(err)
	}
	if err := lockedB.goOut(); err != nil {
		fmt.Println(err)
	}

	if err := otherA.goOut(); err != nil {
		fmt.Println(err)
	}
	if err := otherA.move(1, false, false); err != nil {
		fmt.Println(err)
	}

	if err := otherB.goOut(); err != nil {
		fmt.Println(err)
	}
	if err := otherB.move(1, false, false); err != nil {
		fmt.Println(err)
	}

	// swap with self -> nope
	assertTrue(t, selfA.canSwap(), "can selfA swap")
	assertTrue(t, selfB.canSwap(), "can selfB swap")
	assertFalse(t, selfA.canSwapWith(selfA), "can selfA swap with selfA")
	assertFalse(t, selfB.canSwapWith(selfA), "can selfB swap with selfA")
	assertFalse(t, selfA.canSwapWith(selfB), "can selfA swap with selfB")
	assertNotNil(t, selfA.swap(selfA), "err for swapping selfA with selfA")
	assertNotNil(t, selfB.swap(selfA), "err for swapping selfB with selfA")
	assertNotNil(t, selfA.swap(selfB), "err for swapping selfA with selfB")

	// swap with locked (heaven, home, start) -> nope
	assertFalse(t, lockedA.canSwap(), "can lockedA swap")
	assertFalse(t, lockedB.canSwap(), "can lockedB swap")
	assertFalse(t, lockedC.canSwap(), "can lockedC swap")
	assertFalse(t, selfA.canSwapWith(lockedA), "can selfA swap with lockedA")
	assertFalse(t, selfB.canSwapWith(lockedB), "can selfB swap with lockedB")
	assertFalse(t, selfA.canSwapWith(lockedC), "can selfA swap with lockedC")
	assertFalse(t, lockedA.canSwapWith(selfB), "can selfA swap with lockedA")
	assertFalse(t, lockedB.canSwapWith(selfA), "can selfB swap with lockedB")
	assertFalse(t, lockedC.canSwapWith(selfB), "can selfA swap with lockedC")
	assertNotNil(t, selfA.swap(lockedA), "err for swapping selfA with lockedA")
	assertNotNil(t, selfB.swap(lockedB), "err for swapping selfB with lockedB")
	assertNotNil(t, selfA.swap(lockedC), "err for swapping selfA with lockedC")
	assertNotNil(t, lockedA.swap(selfB), "err for swapping lockedA with selfB")
	assertNotNil(t, lockedB.swap(selfA), "err for swapping lockedB with selfA")
	assertNotNil(t, lockedC.swap(selfB), "err for swapping lockedC with selfB")

	// swap with another
	assertTrue(t, otherA.canSwap(), "can otherA swap")
	assertTrue(t, otherB.canSwap(), "can otherB swap")
	assertTrue(t, otherA.canSwapWith(selfA), "can otherA swap with selfA")
	assertTrue(t, otherB.canSwapWith(selfB), "can otherB swap with selfB")
	assertTrue(t, selfA.canSwapWith(otherA), "can selfA swap with otherA")
	assertTrue(t, selfB.canSwapWith(otherB), "can selfB swap with otherB")
	assertTrue(t, otherA.canSwapWith(otherB), "can otherA swap with otherB")

	assertEq(t, selfA.Position.Section.Color, selfColor)
	assertEq(t, selfB.Position.Section.Color, selfColor)
	assertEq(t, otherA.Position.Section.Color, otherAColor)
	assertEq(t, otherB.Position.Section.Color, otherBColor)

	assertNil(t, selfA.swap(otherA), "no err for swapping selfA with otherA")
	assertNil(t, selfB.swap(otherB), "no err for swapping selfB with otherB")
	assertEq(t, selfA.Position.Section.Color, otherAColor)
	assertEq(t, selfB.Position.Section.Color, otherBColor)
	assertEq(t, otherA.Position.Section.Color, selfColor)
	assertEq(t, otherB.Position.Section.Color, selfColor)
	assertEq(t, otherAPlayer.Section.Home[0].NextPosition.NextPosition.Marble, selfA)
	assertEq(t, otherBPlayer.Section.Home[0].NextPosition.NextPosition.Marble, selfB)

	assertNil(t, otherA.swap(selfA), "no err for swapping otherA with selfA")
	assertNil(t, otherB.swap(selfB), "no err for swapping otherB with selfB")
	assertEq(t, selfA.Position.Section.Color, selfColor)
	assertEq(t, selfB.Position.Section.Color, selfColor)
	assertEq(t, otherA.Position.Section.Color, otherAColor)
	assertEq(t, otherB.Position.Section.Color, otherBColor)
	assertEq(t, otherAPlayer.Section.Home[0].NextPosition.NextPosition.Marble, otherA)
	assertEq(t, otherBPlayer.Section.Home[0].NextPosition.NextPosition.Marble, otherB)

	assertNil(t, otherA.swap(otherB), "no err for swapping otherA with otherB")
	assertEq(t, otherA.Position.Section.Color, otherBColor)
	assertEq(t, otherB.Position.Section.Color, otherAColor)
	assertEq(t, otherAPlayer.Section.Home[0].NextPosition.NextPosition.Marble, otherB)
	assertEq(t, otherBPlayer.Section.Home[0].NextPosition.NextPosition.Marble, otherA)

	assertNil(t, otherB.swap(otherA), "no err for swapping otherB with otherA")
	assertEq(t, otherA.Position.Section.Color, otherAColor)
	assertEq(t, otherB.Position.Section.Color, otherBColor)
	assertEq(t, otherAPlayer.Section.Home[0].NextPosition.NextPosition.Marble, otherA)
	assertEq(t, otherBPlayer.Section.Home[0].NextPosition.NextPosition.Marble, otherB)
}

func Test_max(t *testing.T) {
	gb := setupGame()

	marble1 := gb.Players[0].Marbles[3]
	marble2 := gb.Players[1].Marbles[3]

	if err := marble1.goOut(); err != nil {
		fmt.Println(err)
	}

	if err := marble1.move(12, false, false); err != nil {
		fmt.Println(err)
	}

	if err := marble2.goOut(); err != nil {
		fmt.Println(err)
	}

	assertEq(t, marble1.max(1), 1)
	assertEq(t, marble1.max(2), 2)
	assertEq(t, marble1.max(3), 3)
	assertEq(t, marble1.max(4), 3)
	assertEq(t, marble1.max(5), 3)
	assertEq(t, marble1.max(6), 3)

	assertEq(t, marble2.max(6), 6)
}
