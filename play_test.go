package main

import "testing"

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

func Test_something(t *testing.T) {
	gb := setupGame(t)
	gb.
}
