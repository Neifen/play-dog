package main

import (
	"fmt"
	"math/rand/v2"
	"slices"

	"github.com/google/uuid"
)

var (
	ErrUUIDTaken      = fmt.Errorf("lobby: player with uuid already exists")
	ErrPlayerNotFound = fmt.Errorf("lobby: player with uuid not found")
	ErrColorTaken     = fmt.Errorf("lobby: color is already taken")
	ErrPartnerTaken   = fmt.Errorf("lobby: player already has a partner")
	ErrNotReady       = fmt.Errorf("lobby: lobby is not ready")
)

type Lobby struct {
	players         map[uuid.UUID]*LobbyPlayer
	AvailableColors []Color
}

func NewLobby() *Lobby {
	return &Lobby{
		players:         make(map[uuid.UUID]*LobbyPlayer),
		AvailableColors: []Color{Blue, Red, Yellow, Green, White, Black},
	}
}

type Seat int
type PlayerSetup struct {
	ID      uuid.UUID
	Name    string
	Color   Color
	Seat    Seat
	Partner uuid.UUID
}

func (l *Lobby) Start() (*Game, error) {
	if !l.ready() {
		return nil, ErrNotReady
	}
	l.assignMissingParters()
	l.assignMissingColors()
	l.assignSeats()
	gameSetup, err := l.setupPlayers()
	if err != nil {
		return nil, fmt.Errorf("lobby: start failed: %w", err)
	}

	return NewGame(gameSetup), nil
}

func (l *Lobby) ready() bool {
	nr := len(l.players)

	return nr == 4 || nr == 6
}

func (l *Lobby) setupPlayers() ([]*PlayerSetup, error) {
	players := make([]*PlayerSetup, 0, len(l.players))

	for id, player := range l.players {
		if player.partner == nil {
			return nil, fmt.Errorf("lobby: player %s doesn't have a partner during setup", id)
		}
		if player.color == nil {
			return nil, fmt.Errorf("lobby: player %s doesn't have a color during setup", id)
		}
		if player.seat == nil {
			return nil, fmt.Errorf("lobby: player %s doesn't have a seat during setup", id)
		}

		player := &PlayerSetup{
			ID:      player.id,
			Name:    player.name,
			Color:   *player.color,
			Partner: *player.partner,
			Seat:    *player.seat,
		}
		players = append(players, player)
	}

	return players, nil
}

func (l *Lobby) assignSeats() error {
	playerAmount := len(l.players)
	index := 0
	for id, player := range l.players {
		if player.partner == nil {
			return fmt.Errorf("lobby: player %s doesn't have a partner during seat assign", id)
		}

		if player.seat != nil {
			continue
		}

		partner, ok := l.players[*player.partner]
		if !ok {
			return fmt.Errorf("%w: partner %s for player %s", ErrPlayerNotFound, *player.partner, id)
		}

		seat := Seat(index)
		player.seat = &seat
		opositeSeat := Seat((playerAmount / 2) + index)
		partner.seat = &opositeSeat
		index++
	}
	return nil
}

func (l *Lobby) assignMissingParters() {
	unasigned := []*LobbyPlayer{}
	for _, player := range l.players {
		if player.partner == nil {
			unasigned = append(unasigned, player)
		}
	}

	rand.Shuffle(len(unasigned), func(i, j int) {
		unasigned[i], unasigned[j] = unasigned[j], unasigned[i]
	})

	for i := 0; i < len(unasigned); i++ {
		unasigned[i].partner = &unasigned[i+1].id
		unasigned[i+1].partner = &unasigned[i].id

		i++
	}
}

func (l *Lobby) assignMissingColors() {
	for _, player := range l.players {
		if player.color == nil {
			i := rand.IntN(len(l.AvailableColors))
			c := l.AvailableColors[i]
			player.color = &c
			l.AvailableColors = slices.Delete(l.AvailableColors, i, i+1)
		}
	}
}

func (l *Lobby) Join(name string) (uuid.UUID, error) {
	player := newLobbyPlayer(name)
	if _, exists := l.players[player.id]; exists {
		return uuid.Nil, ErrUUIDTaken
	}
	l.players[player.id] = player
	return player.id, nil
}

func (l *Lobby) ChooseColor(id uuid.UUID, color Color) error {
	idx := slices.Index(l.AvailableColors, color)
	if idx == -1 {
		return ErrColorTaken
	}
	player, ok := l.players[id]
	if !ok {
		return fmt.Errorf("%w: player %s", ErrPlayerNotFound, id)
	}

	prevColor := player.color

	l.AvailableColors = slices.Delete(l.AvailableColors, idx, idx+1)
	player.color = &color

	if prevColor != nil {
		l.AvailableColors = append(l.AvailableColors, *prevColor)
	}

	return nil
}

func (l *Lobby) ChoosePartner(one, two uuid.UUID) error {
	if one == two {
		return fmt.Errorf("player %s can not be it's own partner", one)
	}
	p1, ok1 := l.players[one]
	p2, ok2 := l.players[two]
	if !ok1 {
		return fmt.Errorf("%w: player %s", ErrPlayerNotFound, one)
	}
	if !ok2 {
		return fmt.Errorf("%w: player %s", ErrPlayerNotFound, two)
	}

	if p2.partner != nil {
		return ErrPartnerTaken
	}

	// free previous partner
	if p1.partner != nil {
		prev, ok := l.players[*p1.partner]
		if !ok {
			return fmt.Errorf("%w: partner %s for player %s", ErrPlayerNotFound, *p1.partner, one)
		}
		prev.partner = nil
	}

	p1.partner = &two
	p2.partner = &one
	return nil
}

func (l *Lobby) LoosePartner(id uuid.UUID) error {
	player, ok := l.players[id]
	if !ok {
		return fmt.Errorf("%w: player %s", ErrPlayerNotFound, id)
	}

	if player.partner == nil {
		return nil
	}

	prevPartner, ok := l.players[*player.partner]
	if !ok {
		return fmt.Errorf("%w: partner %s for player %s", ErrPlayerNotFound, *player.partner, id)
	}
	player.partner = nil
	prevPartner.partner = nil

	return nil
}

type LobbyPlayer struct {
	id      uuid.UUID
	name    string
	partner *uuid.UUID
	color   *Color
	seat    *Seat
}

func newLobbyPlayer(name string) *LobbyPlayer {
	return &LobbyPlayer{
		name: name,
		id:   uuid.New(),
	}
}
