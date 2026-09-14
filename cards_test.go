package main

import (
	"fmt"
	"testing"
)

// First Marble is 8 in front, Second is 4 behind, Third is in heaven, Fourth is locking start
func setupMarbles(gb *GameBoard) (*Marble, *Marble, *Marble, *Marble) {
	// can't use names for this as order is random
	first := gb.Players[0]
	second := gb.Players[1]
	third := gb.Players[2]
	fourth := gb.Players[3]

	if err := first.Marbles[0].goOut(); err != nil {
		fmt.Println(err)
	}
	if err := first.Marbles[0].move(8, false, false); err != nil {
		fmt.Println(err)
	}

	if err := second.Marbles[0].goOut(); err != nil {
		fmt.Println(err)
	}
	if err := second.Marbles[0].move(-4, false, false); err != nil {
		fmt.Println(err)
	}

	if err := third.Marbles[0].goOut(); err != nil {
		fmt.Println(err)
	}
	if err := third.Marbles[0].move(-4, false, false); err != nil {
		fmt.Println(err)
	}
	if err := third.Marbles[0].move(5, true, false); err != nil {
		fmt.Println(err)
	}

	if err := fourth.Marbles[0].goOut(); err != nil {
		fmt.Println(err)
	}

	return first.Marbles[0], second.Marbles[0], third.Marbles[0], fourth.Marbles[0]
}

func Test_Two(t *testing.T) {
	gb := setupGame(t)
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	card := Two
	blueMoves := card.LegalMoves(gb, firstAfter)
	greenMoves := card.LegalMoves(gb, secondBefore)
	yellowMoves := card.LegalMoves(gb, thirdHeaven)
	blackMoves := card.LegalMoves(gb, fourthStart)

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	assertEqMsg(t, len(blueMoves), 1, "len blueMoves")
	assertEqMsg(t, len(greenMoves), 1, "len greenMoves")
	assertEqMsg(t, len(yellowMoves), 1, "len yellowMoves")
	assertEqMsg(t, len(blackMoves), 1, "len blackMoves")

	assertNoErr(t, card.ApplyMove(blueMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(greenMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(yellowMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(blackMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -6]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -2]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 2]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR 2]", fourthColor), "position fourthStart")
}

func Test_Three(t *testing.T) {
	gb := setupGame(t)
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	card := Three
	blueMoves := card.LegalMoves(gb, firstAfter)
	greenMoves := card.LegalMoves(gb, secondBefore)
	yellowMoves := card.LegalMoves(gb, thirdHeaven)
	blackMoves := card.LegalMoves(gb, fourthStart)

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	assertEqMsg(t, len(blueMoves), 1, "len blueMoves")
	assertEqMsg(t, len(greenMoves), 1, "len greenMoves")
	assertEqMsg(t, len(yellowMoves), 1, "len yellowMoves")
	assertEqMsg(t, len(blackMoves), 1, "len blackMoves")

	assertNoErr(t, card.ApplyMove(blueMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(greenMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(yellowMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(blackMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -5]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -1]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 3]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR 3]", fourthColor), "position fourthStart")
}

func Test_Five(t *testing.T) {
	gb := setupGame(t)
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	card := Five
	blueMoves := card.LegalMoves(gb, firstAfter)
	greenMoves := card.LegalMoves(gb, secondBefore)
	yellowMoves := card.LegalMoves(gb, thirdHeaven)
	blackMoves := card.LegalMoves(gb, fourthStart)

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	assertEqMsg(t, len(blueMoves), 1, "len blueMoves")
	assertEqMsg(t, len(greenMoves), 2, "len greenMoves")
	assertEqMsg(t, len(yellowMoves), 0, "len yellowMoves")
	assertEqMsg(t, len(blackMoves), 1, "len blackMoves")

	assertNoErr(t, card.ApplyMove(blueMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(greenMoves[0]), "apply Blue move")
	// assertNoErr(t, card.ApplyMove(yellowMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(blackMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -3]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR 1]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR 5]", fourthColor), "position fourthStart")
}
