package main

import (
	"math"
	"testing"

	"github.com/google/uuid"
)

func Test_Join(t *testing.T) {
	lobby := NewLobby()

	a, err := lobby.Join("a")
	assertNoErr(t, err, "join a")
	b, err := lobby.Join("b")
	assertNoErr(t, err, "join b")
	c, err := lobby.Join("c")
	assertNoErr(t, err, "join c")
	d, err := lobby.Join("d")
	assertNoErr(t, err, "join d")

	assertEq(t, len(lobby.players), 4)
	assertEq(t, len(lobby.AvailableColors), 6)

	game, err := lobby.Start()
	assertNoErr(t, err, "lobby start")

	assertEq(t, len(game.playersMap), 4)
	assertEq(t, len(game.players), 4)

	assertEq(t, game.playersMap[a].id, a)
	assertEq(t, game.playersMap[b].id, b)
	assertEq(t, game.playersMap[c].id, c)
	assertEq(t, game.playersMap[d].id, d)

	assertEqMsg(t, game.playersMap[a].name, "a", "check player name a")
	assertEqMsg(t, game.playersMap[b].name, "b", "check player name b")
	assertEqMsg(t, game.playersMap[c].name, "c", "check player name c")
	assertEqMsg(t, game.playersMap[d].name, "d", "check player name d")

	for _, player := range game.players {
		assertEqMsg(t, game.playersMap[player.id], player, "testing player %s", player.name)
		assertNotEmpty(t, game.playersMap[player.id].color, "color %s", player.name)
		partner := game.playersMap[player.id].partner
		assertNeq(t, partner, uuid.Nil)
		assertNotNil(t, game.playersMap[player.id].boardPlayer, "boardPlayer %s", player.name)
		assertEqMsg(t, game.playersMap[player.id].boardPlayer.Color, game.playersMap[player.id].color, "boardPlayer color %s", player.name)
		assertEqMsg(t, game.playersMap[player.id].boardPlayer.Name, player.name, "boardPlayer name %s", player.name)

		assertEqMsg(t, math.Abs(float64(int(*lobby.players[player.id].seat)-int(*lobby.players[partner].seat))), 2.0, "partner have oposite seats: %s-%s", player.id, partner)
	}
}

func Test_ChooseColorNonsense(t *testing.T) {
	lobby := NewLobby()
	_, err := lobby.Join("a")
	assertNoErr(t, err, "join a")
	err = lobby.ChooseColor(uuid.New(), White)
	assertErr(t, err, "nonsense uid")
}

func Test_ChooseColorTaken(t *testing.T) {
	lobby := NewLobby()
	a, err := lobby.Join("a")
	assertNoErr(t, err, "join a")
	err = lobby.ChooseColor(a, White)
	assertNoErr(t, err, "choose white for a")

	b, err := lobby.Join("b")
	assertNoErr(t, err, "join b")
	err = lobby.ChooseColor(b, White)
	assertErr(t, err, "choose white again for b")
}

func Test_BadAmountOfPlayers(t *testing.T) {
	lobby := NewLobby()
	_, err := lobby.Join("a")
	assertNoErr(t, err, "join a")
	_, err = lobby.Start()
	assertErr(t, err, "start after 1 join")

	_, err = lobby.Join("b")
	assertNoErr(t, err, "join b")
	_, err = lobby.Start()
	assertErr(t, err, "start after 2 joins")

	_, err = lobby.Join("b")
	assertNoErr(t, err, "join b2")
	_, err = lobby.Start()
	assertErr(t, err, "start after 3 joins")

	_, err = lobby.Join("c")
	assertNoErr(t, err, "join c")
	_, err = lobby.Join("d")
	assertNoErr(t, err, "join d")

	_, err = lobby.Start()
	assertErr(t, err, "start after 5 joins")

}

func Test_ChooseColor(t *testing.T) {
	lobby := NewLobby()

	a, err := lobby.Join("a")
	assertNoErr(t, err, "join a")
	err = lobby.ChooseColor(a, White)
	assertNoErr(t, err, "choose white for a")

	b, err := lobby.Join("b")
	assertNoErr(t, err, "join b")
	err = lobby.ChooseColor(b, Red)
	assertNoErr(t, err, "choose red for b")

	c, err := lobby.Join("c")
	assertNoErr(t, err, "join c")

	d, err := lobby.Join("d")
	assertNoErr(t, err, "join d")

	err = lobby.ChooseColor(c, Yellow)
	assertNoErr(t, err, "choose yellow for c")
	err = lobby.ChooseColor(d, Blue)
	assertNoErr(t, err, "choose blue for d")

	assertEq(t, len(lobby.players), 4)
	assertEq(t, len(lobby.AvailableColors), 2)

	game, err := lobby.Start()
	assertNoErr(t, err, "lobby start")

	assertEq(t, len(game.playersMap), 4)
	assertEq(t, len(game.players), 4)

	assertEq(t, game.playersMap[a].id, a)
	assertEq(t, game.playersMap[b].id, b)
	assertEq(t, game.playersMap[c].id, c)
	assertEq(t, game.playersMap[d].id, d)

	assertEqMsg(t, game.playersMap[a].color, White, "check player a color")
	assertEqMsg(t, game.playersMap[b].color, Red, "check player b color")
	assertEqMsg(t, game.playersMap[c].color, Yellow, "check player c color")
	assertEqMsg(t, game.playersMap[d].color, Blue, "check player d color")

	for _, player := range game.players {
		assertEqMsg(t, game.playersMap[player.id], player, "testing player %s", player.name)
		assertNotEmpty(t, game.playersMap[player.id].color, "color %s", player.name)
		assertEqMsg(t, game.playersMap[player.id].boardPlayer.Color, game.playersMap[player.id].color, "boardPlayer color %s", player.name)
	}
}

func Test_ChoosePartnerOwn(t *testing.T) {
	lobby := NewLobby()
	a, err := lobby.Join("a")
	assertNoErr(t, err, "join a")
	err = lobby.ChoosePartner(a, a)
	assertErr(t, err, "choose a for a")
}

func Test_ChoosePartnerNonSense(t *testing.T) {
	lobby := NewLobby()
	a, err := lobby.Join("a")
	assertNoErr(t, err, "join a")
	err = lobby.ChoosePartner(a, uuid.New())
	assertErr(t, err, "choose a for unknown")
	err = lobby.ChoosePartner(uuid.New(), a)
	assertErr(t, err, "choose a for unknown")
}

func Test_ChoosePartner(t *testing.T) {
	lobby := NewLobby()

	a, err := lobby.Join("a")
	assertNoErr(t, err, "join a")

	b, err := lobby.Join("b")
	assertNoErr(t, err, "join b")

	c, err := lobby.Join("c")
	assertNoErr(t, err, "join c")

	d, err := lobby.Join("d")
	assertNoErr(t, err, "join d")

	err = lobby.ChoosePartner(c, a)
	assertNoErr(t, err, "choose c for a")
	assertNotNil(t, lobby.players[c].partner, "c partner")
	assertNotNil(t, lobby.players[a].partner, "a partner")

	// now its c-b and a-d
	err = lobby.ChoosePartner(c, b)
	assertNoErr(t, err, "choose c for b")
	assertNotNil(t, lobby.players[c].partner, "c partner")
	assertNotNil(t, lobby.players[b].partner, "b partner")
	assertNil(t, lobby.players[a].partner, "a partner")

	// taken
	err = lobby.ChoosePartner(a, c)
	assertErr(t, err, "choose a for c")

	assertEq(t, len(lobby.players), 4)

	game, err := lobby.Start()
	assertNoErr(t, err, "lobby start")

	assertEq(t, len(game.playersMap), 4)
	assertEq(t, len(game.players), 4)

	assertNotNil(t, lobby.players[a].seat, "a seat not nil")
	assertNotNil(t, lobby.players[b].seat, "b seat not nil")
	assertNotNil(t, lobby.players[c].seat, "c seat not nil")
	assertNotNil(t, lobby.players[d].seat, "d seat not nil")

	assertEqMsg(t, math.Abs(float64(int(*lobby.players[a].seat)-int(*lobby.players[d].seat))), 2.0, "partner have oposite seats: a-d")
	assertEqMsg(t, math.Abs(float64(int(*lobby.players[c].seat)-int(*lobby.players[b].seat))), 2.0, "partner have oposite seats: b-c")

	assertEq(t, game.playersMap[a].id, a)
	assertEq(t, game.playersMap[b].id, b)
	assertEq(t, game.playersMap[c].id, c)
	assertEq(t, game.playersMap[d].id, d)

	assertEqMsg(t, game.playersMap[a].partner, d, "a-d partnership")
	assertEqMsg(t, game.playersMap[d].partner, a, "d-a partnership")
	assertEqMsg(t, game.playersMap[b].partner, c, "b-c partnership")
	assertEqMsg(t, game.playersMap[c].partner, b, "c-b partnership")

	for _, player := range game.players {
		assertEqMsg(t, game.playersMap[player.id], player, "testing player %s", player.name)
		assertNotEmpty(t, game.playersMap[player.id].color, "color %s", player.name)
		assertEqMsg(t, game.playersMap[player.id].partner, player.partner, "playersMap partner %s", player.name)
	}
}
