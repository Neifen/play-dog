package main

import (
	"fmt"
	"log"
	"testing"
)

func Test_Two_Joker(t *testing.T) {
	gb := setupGame(t)
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	card := Joker
	card = card.JockerPickCard(Two)

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

func Test_Three_Joker(t *testing.T) {
	gb := setupGame(t)
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	card := Joker
	card = card.JockerPickCard(Three)

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

func Test_FiveA_Joker(t *testing.T) {
	gb := setupGame(t)
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	card := Joker
	card = card.JockerPickCard(Five)
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

func Test_FiveB_Joker(t *testing.T) {
	gb := setupGame(t)
	firstAfter, secondBefore, thirdHeaven, fourthStart := setupMarbles(gb)
	firstColor := firstAfter.Player.Color
	secondColor := secondBefore.Player.Color
	thirdColor := thirdHeaven.Player.Color
	fourthColor := fourthStart.Player.Color
	_ = firstColor

	card := Joker
	card = card.JockerPickCard(Five)

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

func Test_SixA_Joker(t *testing.T) {
	card := Joker
	card = card.JockerPickCard(Six)

	gb := setupGame(t)
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

func Test_SixB_Joker(t *testing.T) {
	card := Joker
	card = card.JockerPickCard(Six)

	gb := setupGame(t)
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

func Test_EightA_Joker(t *testing.T) {
	card := Joker
	card = card.JockerPickCard(Eight)

	gb := setupGame(t)
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

func Test_EightB_Joker(t *testing.T) {
	card := Joker
	card = card.JockerPickCard(Eight)

	gb := setupGame(t)
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

func Test_Nine_Joker(t *testing.T) {
	card := Joker
	card = card.JockerPickCard(Nine)

	gb := setupGame(t)
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

func Test_Ten_Joker(t *testing.T) {
	card := Joker
	card = card.JockerPickCard(Ten)

	gb := setupGame(t)
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

func Test_Queen_Joker(t *testing.T) { //12
	card := Joker
	card = card.JockerPickCard(Queen)

	gb := setupGame(t)
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

func Test_King_Joker(t *testing.T) { //13
	card := Joker
	card = card.JockerPickCard(King)

	gb := setupGame(t)
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

func Test_AceOne_Joker(t *testing.T) {
	card := Joker
	card = card.JockerPickCard(Ace)

	gb := setupGame(t)
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

func Test_AceEleven_Joker(t *testing.T) {
	card := Joker
	card = card.JockerPickCard(Ace)

	gb := setupGame(t)
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

func Test_FourBack_Joker(t *testing.T) {
	card := Joker
	card = card.JockerPickCard(Four)

	gb := setupGame(t)
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

func Test_FourForward_Joker(t *testing.T) {
	card := Joker
	card = card.JockerPickCard(Four)

	gb := setupGame(t)
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

func Test_Seven_Joker(t *testing.T) {
	card := Joker
	card = card.JockerPickCard(Seven)

	gb := setupGame(t)
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

func Test_Jack_Joker(t *testing.T) {
	card := Joker
	card = card.JockerPickCard(Jack)

	gb := setupGame(t)
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
