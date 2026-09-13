package main

import (
	"fmt"
	"math"
	"slices"
	"testing"
)

func Test_createBoard_4(t *testing.T) {
	gb := NewGameBoard()

	gb.Join("a")
	assertEq(t, gb.Ready(), false)
	gb.Join("b")
	assertEq(t, gb.Ready(), false)
	gb.Join("c")
	assertEq(t, gb.Ready(), false)
	gb.Join("d")
	assertEq(t, gb.Ready(), true)

	err := gb.Start()
	if err != nil {
		t.Fatalf("failed to create gameboard with error %v", err)
	}
	fmt.Println(gb.Players)
	fmt.Println(gb.Sections)

	assertEq(t, len(gb.Players), 4)
	assertEq(t, len(gb.Sections), 4)

	assertEq(t, gb.Players[0].Section, gb.Sections[0])
	assertEq(t, gb.Players[1].Section, gb.Sections[1])
	assertEq(t, gb.Players[2].Section, gb.Sections[2])
	assertEq(t, gb.Players[3].Section, gb.Sections[3])

	assertEq(t, gb.Players[0].Partner, gb.Players[2])
	assertEq(t, gb.Players[2].Partner, gb.Players[0])
	assertEq(t, gb.Players[1].Partner, gb.Players[3])
	assertEq(t, gb.Players[3].Partner, gb.Players[1])

	start := gb.Sections[0].First
	next := start
	for i := range 64 {
		if (i-9)%16 == 0 {
			assertNotNil(t, next.AltNextPosition, "next alt next Position")
			assertEq(t, next.AltNextPosition.PositionType, Heaven)
		} else {
			assertNil(t, next.AltNextPosition, "next alt next Position")
		}
		next = next.NextPosition
	}

	for i := range 4 {
		for j := range 4 {
			assertEq(t, gb.Players[i].Marbles[j], gb.Players[i].Section.Home[j].Marble)
		}

	}
	assertEq(t, next, start)
}

func Test_PickColorsAndPartners(t *testing.T) {
	gb := NewGameBoard()

	a := gb.Join("a")
	b := gb.Join("b")
	c := gb.Join("c")
	d := gb.Join("d")

	assertEq(t, len(gb.AvailableColors), 6)
	gb.ChooseColor(a, Blue)
	assertEq(t, len(gb.AvailableColors), 5)

	gb.ChooseColor(c, Black)
	assertEq(t, len(gb.AvailableColors), 4)

	err := gb.ChooseColor(c, Black)
	assertNotNil(t, err, "err chooseColor c:Black")

	gb.ChoosePartner(b, d)

	err = gb.ChoosePartner(a, b)
	assertNotNil(t, err, "err choosePartner a,b")

	err = gb.Start()
	if err != nil {
		t.Fatalf("failed to create gameboard with error %v", err)
	}
	fmt.Println(gb.Players)
	fmt.Println(gb.Sections)

	aIdx := slices.Index(gb.Players, a)
	bIdx := slices.Index(gb.Players, b)
	cIdx := slices.Index(gb.Players, c)
	dIdx := slices.Index(gb.Players, d)

	assertEq(t, gb.Players[aIdx].Color, Blue)
	assertEq(t, gb.Players[cIdx].Color, Black)

	assertEq(t, gb.Players[bIdx].Partner.Name, "d")
	assertEq(t, gb.Players[dIdx].Partner.Name, "b")

	assertEq(t, math.Abs(float64(bIdx-dIdx)), 2.0)
	assertEq(t, math.Abs(float64(aIdx-cIdx)), 2.0)

}
