package main

import (
	"fmt"
	"log"
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
	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	card := Two
	firstMoves := card.LegalMoves(gb, firstAfter)
	secondMoves := card.LegalMoves(gb, secondBefore)
	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	fourthMoves := card.LegalMoves(gb, fourthStart)

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertEqMsg(t, len(secondMoves), 1, "len secondMoves")
	assertEqMsg(t, len(thirdMoves), 1, "len thirdMoves")
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")

	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -6]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -2]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 2]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR 2]", fourthColor), "position fourthStart")
}

func Test_Three(t *testing.T) {
	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	card := Three
	firstMoves := card.LegalMoves(gb, firstAfter)
	secondMoves := card.LegalMoves(gb, secondBefore)
	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	fourthMoves := card.LegalMoves(gb, fourthStart)

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertEqMsg(t, len(secondMoves), 1, "len secondMoves")
	assertEqMsg(t, len(thirdMoves), 1, "len thirdMoves")
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")

	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -5]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -1]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 3]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR 3]", fourthColor), "position fourthStart")
}

func Test_FiveA(t *testing.T) {
	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	card := Five
	firstMoves := card.LegalMoves(gb, firstAfter)
	secondMoves := card.LegalMoves(gb, secondBefore)
	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	fourthMoves := card.LegalMoves(gb, fourthStart)

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertEqMsg(t, len(secondMoves), 2, "len secondMoves")
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")

	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -3]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR 5]", fourthColor), "position fourthStart")
}

func Test_FiveB(t *testing.T) {
	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	card := Five
	firstMoves := card.LegalMoves(gb, firstAfter)
	secondMoves := card.LegalMoves(gb, secondBefore)
	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	fourthMoves := card.LegalMoves(gb, fourthStart)

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertEqMsg(t, len(secondMoves), 2, "len secondMoves")
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")

	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(secondMoves[1]), "apply Blue move")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -3]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR 1]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR 5]", fourthColor), "position fourthStart")
}

func Test_SixA(t *testing.T) {
	card := Six

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")

	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 2, "len secondMoves")
	assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")

	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -2]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[HEAVEN 1]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR 6]", fourthColor), "position fourthStart")
}

func Test_SixB(t *testing.T) {
	card := Six

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")

	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 2, "len secondMoves")
	assertNoErr(t, card.ApplyMove(secondMoves[1]), "apply Blue move")

	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -2]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR 2]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR 6]", fourthColor), "position fourthStart")
}

func Test_EightA(t *testing.T) {
	card := Eight

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")

	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 2, "len secondMoves")
	assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")

	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[START 0]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[HEAVEN 3]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR -8]", firstColor), "position fourthStart")
}

func Test_EightB(t *testing.T) {
	card := Eight

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")

	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 2, "len secondMoves")
	assertNoErr(t, card.ApplyMove(secondMoves[1]), "apply Blue move")

	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[START 0]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR 4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR -8]", firstColor), "position fourthStart")
}

func Test_Nine(t *testing.T) {
	card := Nine

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")

	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 1, "len secondMoves")
	assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")

	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR 1]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR 5]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR -7]", firstColor), "position fourthStart")
}

func Test_Ten(t *testing.T) {
	card := Ten

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")

	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 1, "len secondMoves")
	assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")

	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR 2]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR 6]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR -6]", firstColor), "position fourthStart")
}

func Test_Queen(t *testing.T) { //12
	card := Queen

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")

	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 1, "len secondMoves")
	assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")

	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR 4]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -8]", thirdColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR -4]", firstColor), "position fourthStart")
}

func Test_King(t *testing.T) { //12
	card := King

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")

	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 1, "len secondMoves")
	assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")

	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR 5]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -7]", thirdColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR -3]", firstColor), "position fourthStart")
}

func Test_AceOne(t *testing.T) {
	card := Ace

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 2, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")

	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 2, "len secondMoves")
	assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")

	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 1, "len thirdMoves")
	assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 2, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -7]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -3]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 1]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR 1]", fourthColor), "position fourthStart")
}

func Test_AceEleven(t *testing.T) {
	card := Ace

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 2, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[1]), "apply Blue move")

	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 2, "len secondMoves")
	assertNoErr(t, card.ApplyMove(secondMoves[1]), "apply Blue move")

	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 1, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 2, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[1]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR 3]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -9]", thirdColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR -5]", firstColor), "position fourthStart")
}

func Test_FourBack(t *testing.T) {
	card := Four

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 2, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")

	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 2, "len secondMoves")
	assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")

	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 2, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR 4]", firstColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR -4]", fourthColor), "position fourthStart")
}

func Test_FourForward(t *testing.T) {
	card := Four

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 2, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[1]), "apply Blue move")

	// second got eaten
	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 0, "len secondMoves")
	// assertNoErr(t, card.ApplyMove(secondMoves[1]), "apply Blue move")

	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 2, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[1]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[HOME 0]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR 4]", fourthColor), "position fourthStart")
}

func Test_Seven(t *testing.T) {
	card := Seven

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	// first eats second
	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")

	// second got eaten
	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 0, "len secondMoves")
	// assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[HOME 0]", secondColor), "position secondBefore")

	//third, cant move, only marble
	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	// third, move other marble out
	if err := gb.Players[2].Marbles[1].goOut(); err != nil {
		fmt.Println(err)
	}
	thirdMoves = card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 3, "len thirdMoves")

	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 1, "len fourthMoves")
	assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -1]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[HOME 0]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[REGULAR -9]", firstColor), "position fourthStart")
}

func Test_Jack(t *testing.T) {
	card := Jack

	gb := setupGame()
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

	// first, add another marble
	if err := gb.Players[0].Marbles[2].goOut(); err != nil {
		log.Fatal(err)
	}
	if err := gb.Players[0].Marbles[2].move(2, false, false); err != nil {
		log.Fatal(err)
	}

	// first, can only swap with second
	firstMoves := card.LegalMoves(gb, firstAfter)
	assertEqMsg(t, len(firstMoves), 1, "len firstMoves")
	assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply Blue move")

	// second, has two marbles from player 1 to choose from
	secondMoves := card.LegalMoves(gb, secondBefore)
	assertEqMsg(t, len(secondMoves), 2, "len secondMoves")
	assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Blue move")

	// third, is in heaven, no bueno
	thirdMoves := card.LegalMoves(gb, thirdHeaven)
	assertEqMsg(t, len(thirdMoves), 0, "len thirdMoves")
	// assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Blue move")

	// forth, blocked
	fourthMoves := card.LegalMoves(gb, fourthStart)
	assertEqMsg(t, len(fourthMoves), 0, "len fourthMoves")
	// assertNoErr(t, card.ApplyMove(fourthMoves[0]), "apply Blue move")

	assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
	assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
	assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
	assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")
}

func Test_GoOut(t *testing.T) {

	cards := []Card{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Ace, Queen, King, Jack}

	for _, card := range cards {
		gb := setupGame()
		firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
		firstColor := firstAfter.Player.Color
		secondColor := secondBefore.Player.Color
		thirdColor := thirdHeaven.Player.Color
		fourthColor := fourthStart.Player.Color
		_ = firstColor

		assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
		assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
		assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
		assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")

		for i := range 3 {
			if i == 0 {
				continue // don't test first marble
			}
			result := 0
			if card.leaveHome {
				result = 1
			}

			// first player
			firstPMarble := gb.Players[0].Marbles[i]
			firstMoves := card.LegalMoves(gb, firstPMarble)
			assertEqMsg(t, len(firstMoves), result, "len firstMoves")
			if i == 3 {
				assertNoErr(t, card.ApplyMove(firstMoves[0]), "apply First player move")
				assertEqMsg(t, firstPMarble.Position.String(), fmt.Sprintf("%s[START 0]", firstColor), "position firstPMarble")
			}

			// second player
			secondPMarble := gb.Players[1].Marbles[i]
			secondMoves := card.LegalMoves(gb, secondPMarble)
			assertEqMsg(t, len(secondMoves), result, "len secondMoves")
			if i == 3 {
				assertNoErr(t, card.ApplyMove(secondMoves[0]), "apply Second player move")
				assertEqMsg(t, secondPMarble.Position.String(), fmt.Sprintf("%s[START 0]", secondColor), "position secondPMarble")
			}

			thirdPMarble := gb.Players[2].Marbles[i]
			thirdMoves := card.LegalMoves(gb, thirdPMarble)
			assertEqMsg(t, len(thirdMoves), result, "len thirdMoves")
			if i == 3 {
				assertNoErr(t, card.ApplyMove(thirdMoves[0]), "apply Third player move")
				assertEqMsg(t, thirdPMarble.Position.String(), fmt.Sprintf("%s[START 0]", thirdColor), "position thirdPMarble")
			}

			// fourth player is blocked
			fourthPMarble := gb.Players[3].Marbles[i]
			fourthMoves := card.LegalMoves(gb, fourthPMarble)
			assertEqMsg(t, len(fourthMoves), 0, "len fourthMoves")

			if i == 3 {
				assertEqMsg(t, firstAfter.Position.String(), fmt.Sprintf("%s[REGULAR -8]", secondColor), "position firstAfter")
				assertEqMsg(t, secondBefore.Position.String(), fmt.Sprintf("%s[REGULAR -4]", secondColor), "position secondBefore")
				assertEqMsg(t, thirdHeaven.Position.String(), fmt.Sprintf("%s[HEAVEN 0]", thirdColor), "position thirdHeaven")
				assertEqMsg(t, fourthStart.Position.String(), fmt.Sprintf("%s[START 0]", fourthColor), "position fourthStart")
			}
		}

	}
}
