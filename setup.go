package main

import (
	"fmt"

	"github.com/google/uuid"
)

type GameBoard struct {
	Players  []*Player //todo: find out if needed
	Sections []*Section
	// AvailableColors []Color

	// playersByName map[string]*Player
}

func NewGameBoard(gamePlayers []*GamePlayer) *GameBoard {
	players := make([]*Player, 0, len(gamePlayers))
	sections := make([]*Section, 0, len(gamePlayers))
	for _, gamePlayer := range gamePlayers {
		player := NewPlayer(gamePlayer)
		players = append(players, player)
		gamePlayer.boardPlayer = player

		section := NewSection(player)
		sections = append(sections, section)

		marbles := NewMarbles(player)
		player.Marbles = marbles
	}

	linkSections(sections)

	return &GameBoard{
		Players:  players,
		Sections: sections,
	}
}

func linkSections(sections []*Section) {
	playerAmount := len(sections)

	for i, s := range sections {
		nextSectionIndex := (i + 1) % playerAmount
		linkPositions(s.Last, sections[nextSectionIndex].First)
	}
}

type Player struct {
	Name  string
	ID    uuid.UUID
	Color Color
	// Partner *Player
	Marbles Marbles
	Section *Section
}

func (p *Player) String() string {
	return fmt.Sprintf("%s[%s]", p.Name, p.Color)
}

func NewPlayer(p *GamePlayer) *Player {
	return &Player{
		Name:  p.name,
		ID:    p.id,
		Color: p.color,
	}
}

func (p *Player) setup(section *Section) {
	p.Section = section
}

type Marbles []*Marble

func NewMarbles(player *Player) Marbles {
	var marbles Marbles
	for i := range 4 {
		marble := &Marble{
			StartTouches: 0,
			Player:       player,
			Position:     player.Section.Home[i],
		}
		marbles = append(marbles, marble)
		player.Section.Home[i].Marble = marble
	}

	return marbles
}

type Marble struct {
	Player       *Player
	Position     *Position
	StartTouches int // 0 - 1- 2 then can go to heaven
}

func (m *Marble) String() string {
	return fmt.Sprintf("%s isblocking:[%v]", m.Player.Color, m.isBlocking())
}

type Section struct {
	Color  Color
	First  *Position
	Last   *Position
	Home   [4]*Position
	Heaven [4]*Position
}

func (s *Section) String() string {
	return fmt.Sprintf("%s (%v -> %v)", s.Color, s.First, s.Last)
}

func createHeaven(section *Section) [4]*Position {
	var heaven [4]*Position
	for i := range 4 {
		pos := NewPosition(section, Heaven, i)

		if i > 0 {
			pos.LastPosition = heaven[i-1]
			heaven[i-1].NextPosition = pos
		}

		heaven[i] = pos
	}

	return heaven
}

func createHome(section *Section) [4]*Position {
	var home [4]*Position
	for i := 0; i < 4; i++ {
		pos := NewPosition(section, Home, i)
		home[i] = pos
	}

	return home
}

func createField(section *Section) [16]*Position {
	var field [16]*Position
	for i := 0; i < 16; i++ {
		pos := NewPosition(section, Regular, i)

		field[i] = pos

		if i > 0 {
			pos.LastPosition = field[i-1]
			field[i-1].NextPosition = pos
		}
	}

	return field
}

func NewSection(player *Player) *Section {
	section := &Section{
		Color: player.Color,
	}

	heaven := createHeaven(section)
	home := createHome(section)
	field := createField(section)

	start := field[9]
	start.PositionType = Start
	start.AltNextPosition = heaven[0]
	home[0].NextPosition, home[1].NextPosition = start, start
	home[2].NextPosition, home[3].NextPosition = start, start

	section.Home = home
	section.Heaven = heaven
	section.First = field[0]
	section.Last = field[len(field)-1]
	player.Section = section

	return section
}

type PositionType string

const (
	Regular PositionType = "REGULAR"
	Heaven  PositionType = "HEAVEN"
	Start   PositionType = "START"
	Home    PositionType = "HOME"
	Center  PositionType = "CENTER" //extension
)

type Position struct {
	PositionType PositionType

	NextPosition    *Position
	LastPosition    *Position
	AltNextPosition *Position

	Section     *Section
	NextSection *Section
	PrevSection *Section

	Marble *Marble
	ID     string
}

func (p *Position) String() string {
	return fmt.Sprintf("%s", p.ID)
}

func (p *Position) onBoard() bool {
	return p.PositionType == Regular || p.PositionType == Start || p.PositionType == Center
}

func linkPositions(prev, after *Position) {
	prev.NextPosition = after
	after.LastPosition = prev
}

func NewPosition(section *Section, posType PositionType, index int) *Position {
	if posType == Regular {
		// 0 is start for ID
		index -= 9
		if index == 0 {
			posType = Start
		}
	}

	return &Position{
		PositionType: posType,
		Section:      section,
		ID:           fmt.Sprintf("%s[%s %d]", section.Color, posType, index),
	}
}
