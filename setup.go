package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

type GameBoard struct {
	Players  []*Player
	Sections []*Section
}

func NewGameBoard() *GameBoard {
	return &GameBoard{}
}

func Join(name string) {
	
}

func (g *GameBoard) Start(numberOfPlayers int) (*GameBoard, error) {
	if numberOfPlayers != 4 && numberOfPlayers != 6 {
		return nil, fmt.Errorf("Number of players must be 4 or 6")
	}

	availableColors := []Color{Blue, Red, Yellow, Green}
	if numberOfPlayers == 6 {
		availableColors = append(availableColors, White, Black)
	}

	players := make([]*Player, numberOfPlayers)
	sections := make([]*Section, numberOfPlayers)

	for i := range numberOfPlayers {
		randomIndex := rand.IntN(len(availableColors))
		color := availableColors[randomIndex]
		availableColors = slices.Delete(availableColors, randomIndex, randomIndex+1)

		section := NewSection(color)
		player := NewPlayer(color, section)
		player.Marbles = NewMarbles(player)

		sections[i] = section
		players[i] = player
	}

	for i := 0; i < numberOfPlayers; i++ {
		partnerIndex := (i + numberOfPlayers/2) % numberOfPlayers

		players[i].Partner, players[partnerIndex].Partner = players[partnerIndex], players[i]

		nextSectionIndex := (i + 1) % numberOfPlayers
		linkPositions(sections[i].Last, sections[nextSectionIndex].First)
	}

	return &GameBoard{
		Players:  players,
		Sections: sections,
	}, nil
}

type Color string

const (
	Blue   Color = "BLUE"
	Red    Color = "RED"
	Yellow Color = "YELLOW"
	Green  Color = "GREEN"
	Black  Color = "BLACK"
	White  Color = "WHITE"
)

type Player struct {
	Color   Color
	Partner *Player
	Marbles Marbles
	Section *Section
}

func (p *Player) String() string {
	return fmt.Sprintf("%s (+ %s)", p.Color, p.Partner.Color)
}

func NewPlayer(color Color, section *Section) *Player {
	return &Player{
		Color:   color,
		Section: section,
	}
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
	StartTouches uint8 // 0 - 1- 2 then can go to heaven
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

func NewSection(color Color) *Section {
	section := &Section{
		Color: color,
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

func linkPositions(prev, after *Position) {
	prev.NextPosition = after
	after.LastPosition = prev
}

func NewPosition(section *Section, posType PositionType, index int) *Position {
	return &Position{
		PositionType: posType,
		Section:      section,
		ID:           fmt.Sprintf("%s[%s-%d]", section.Color, posType, index),
	}
}
