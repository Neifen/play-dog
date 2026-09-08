package main

import (
	"fmt"
	"reflect"
	"testing"
)

func assertEq(t *testing.T, value, expected any) {
	if value != expected {
		t.Errorf("expected %+v but got %+v", expected, value)
	}
}

func assertNeq(t *testing.T, value, expected any) {
	if value == expected {
		t.Errorf("expected %+v to be different from %+v", expected, value)
	}
}

func assertNil(t *testing.T, value any) {
	v := reflect.ValueOf(value)
	if !v.IsNil() {
		t.Errorf("expected %+v to be nil", value)
	}
}

func assertNotNil(t *testing.T, value any) {
	v := reflect.ValueOf(value)
	if v.IsNil() {
		t.Errorf("expected value to be not be nil")
	}
}

func Test_createBoard_4(t *testing.T) {
	gb, err := NewGameBoard(4)

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
			assertNotNil(t, next.AltNextPosition)
			assertEq(t, next.AltNextPosition.PositionType, Heaven)
		} else {
			assertNil(t, next.AltNextPosition)
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
