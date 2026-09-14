package main

import (
	"fmt"
	"math/rand/v2"
	"play-dog/util"
	"slices"
)

type GameBoard struct {
	Players         []*Player
	Sections        []*Section
	AvailableColors []Color

	playersByName map[string]*Player
}

func NewGameBoard() *GameBoard {
	return &GameBoard{
		playersByName:   make(map[string]*Player),
		AvailableColors: []Color{Blue, Red, Yellow, Green, White, Black},
	}
}

func (gb *GameBoard) Join(name string) (*Player, error) {
	player := NewPlayer(name)
	if _, exists := gb.playersByName[name]; exists {
		return nil, fmt.Errorf("player with name %s already exists", name)
	}
	gb.Players = append(gb.Players, player)
	gb.playersByName[name] = player
	return player, nil
}

func (gb *GameBoard) ChooseColor(player *Player, color Color) error {

	idx := slices.Index(gb.AvailableColors, color)
	if idx == -1 {
		return fmt.Errorf("color %s is already picked", color)
	}

	gb.AvailableColors = slices.Delete(gb.AvailableColors, idx, idx+1)
	player.Color = color
	return nil
}

func (gb *GameBoard) ChoosePartner(one, two *Player) error {
	if one.Partner != nil {
		return fmt.Errorf("player %s, already has a partner", one)
	}

	if two.Partner != nil {
		return fmt.Errorf("player %s, already has a partner", two)
	}

	one.Partner = two
	two.Partner = one
	return nil
}

func (gb *GameBoard) Ready() bool {
	nr := len(gb.Players)

	return nr == 4 || nr == 6
}

func (gb *GameBoard) Start() error {
	numberOfPlayers := len(gb.Players)

	if numberOfPlayers != 4 && numberOfPlayers != 6 {
		return fmt.Errorf("Number of players must be 4 or 6")
	}

	sections := make([]*Section, numberOfPlayers)
	partnerless := []*Player{}

	for i, player := range gb.Players {
		if player.Color == "" {
			randomIndex := rand.IntN(len(gb.AvailableColors))
			player.Color = gb.AvailableColors[randomIndex]
			gb.AvailableColors = slices.Delete(gb.AvailableColors, randomIndex, randomIndex+1)
		}

		if player.Partner == nil {
			partnerless = append(partnerless, player)
		}

		player.Section = NewSection(player)
		player.Marbles = NewMarbles(player)
		sections[i] = player.Section
	}

	// create random partners for the reminding players.
	for len(partnerless) != 0 {
		if len(partnerless)%2 != 0 {
			return fmt.Errorf("Amount of players without a partner is not even: %d", len(partnerless))
		}
		one := partnerless[0]
		two, index := util.RandItemSkipFirst(partnerless)

		one.Partner = two
		two.Partner = one

		partnerless = slices.Delete(partnerless, index, index+1)
		partnerless = slices.Delete(partnerless, 0, 1)
	}

	gb.orderAndLink()
	return nil
}

func (gb *GameBoard) orderAndLink() {
	playerAmount := len(gb.Players)
	unsortedPlayers := gb.Players
	sortedPlayers := make([]*Player, playerAmount)
	sortedSections := make([]*Section, playerAmount)

	index := 0
	for len(unsortedPlayers) > 0 {
		if index >= 3 {
			panic(fmt.Sprintf("something went wrong with ordering players, half of players should not excede 3, was %d", index))
		}
		// Player
		player, playerIndex := util.RandItem(unsortedPlayers)
		sortedPlayers[index] = player
		sortedSections[index] = player.Section
		unsortedPlayers = slices.Delete(unsortedPlayers, playerIndex, playerIndex+1)

		// Partner
		opositeIndex := (playerAmount / 2) + index
		partner := player.Partner
		sortedPlayers[opositeIndex] = partner
		sortedSections[opositeIndex] = partner.Section
		unsortedPlayers = slices.DeleteFunc(unsortedPlayers, func(p *Player) bool {
			if p == nil {
				return false
			}
			return p.ID == partner.ID
		})

		// 0, 1, 2 (max)
		index++
	}

	for i, s := range sortedSections {
		nextSectionIndex := (i + 1) % playerAmount
		linkPositions(s.Last, sortedSections[nextSectionIndex].First)
	}

	gb.Players = sortedPlayers
	gb.Sections = sortedSections
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
	Name    string
	ID      int8
	Color   Color
	Partner *Player
	Marbles Marbles
	Section *Section
}

func (p *Player) String() string {
	return fmt.Sprintf("%s - %s + [%s-%s]", p.Name, p.Color, p.Partner.Name, p.Partner.Color)
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
		ID:   int8(rand.IntN(64)),
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
