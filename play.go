package main

import (
	"fmt"
	"math"
)

func (p *Player) isDone() bool {
	for _, m := range p.Marbles {
		if m.Position.PositionType != Heaven {
			return false
		}
	}

	return true
}

func (p *Player) hasWon() bool {
	return p.isDone() && p.Partner.isDone()
}

func (p *Position) moveFrom(m *Marble) error {
	if p.Marble != m {
		return fmt.Errorf("Marble %s was not on position %s", m, p)
	}

	p.Marble = nil
	return nil
}

func (p *Position) moveTo(m *Marble) error {
	if p.Marble != nil && p.Marble.isBlocking() {
		return fmt.Errorf("Marble %s is blocking position %s for Marble %s", p.Marble, p, m)
	} else if p.Marble != nil {
		p.Marble.sendHome()
	}

	p.Marble = m
	m.Position = p
	return nil
}

func (p *Position) moveVia(m *Marble, seven bool) error {
	if p.Marble != nil && p.Marble.isBlocking() {
		return fmt.Errorf("Marble %s is blocking position %s for Marble %s", p.Marble, p, m)
	} else if p.Marble != nil && seven {
		p.Marble.sendHome()
	}

	return nil
}

func (m *Marble) sendHome() {
	m.Position.Marble = nil
	home := m.Player.Section.Home

	for _, h := range home {
		if h.Marble == nil {
			h.Marble = m
			m.Position = h
			return
		}
	}
}

func (m *Marble) canOut() bool {
	if m.Position.PositionType != Home {
		return false
	}

	return m.Position.NextPosition.Marble == nil || !m.Position.NextPosition.Marble.isBlocking()
}

func (m *Marble) goOut() error {
	if m.Position.PositionType != Home {
		return fmt.Errorf("Marble %s was not at home, Marble was on %s", m, m.Position)
	}

	start := m.Position.NextPosition
	if start.PositionType != Start {
		return fmt.Errorf("Next position %s is supposed to be start but is %s instead", start, start.PositionType)
	}

	if start.Marble != nil && start.Marble.isBlocking() {
		return fmt.Errorf("Start is blocked by %s", start.Marble)
	}

	if err := m.Position.moveFrom(m); err != nil {
		return fmt.Errorf("can't go out %w", err)
	}

	if err := start.moveTo(m); err != nil {
		return fmt.Errorf("can't go out %w", err)
	}

	m.StartTouches = 1

	return nil
}

func (m *Marble) isBlocking() bool {
	if m.Position.PositionType == Heaven {
		return true
	}

	if m.Position.Section != m.Player.Section {
		return false
	}

	return m.StartTouches == 1 && m.Position.PositionType == Start
}

func (m *Marble) canMove(places int, heaven bool) bool {
	backwards := places < 0
	if backwards && heaven {
		return false // can walk back into heaven
	}

	next := m.Position
	startTouches := m.StartTouches

	if backwards && !next.onBoard() {
		return false // can walk back when not on the board
	}

	abs := int(math.Abs(float64(places)))
	for range abs {
		if heaven && next.AltNextPosition != nil && next.Section == m.Player.Section {
			if startTouches < 2 {
				return false
			}
			next = next.AltNextPosition
		} else if places > 0 {
			if next.NextPosition == nil {
				return false // end of heaven
			}

			next = next.NextPosition
		} else {
			next = next.LastPosition
		}

		if next.Marble != nil && next.Marble.isBlocking() {
			return false
		}

		if next.PositionType == Start {
			startTouches++
		}
	}

	// needed in order to get right amount of moves in cards.go
	if heaven && next.PositionType != Heaven {
		return false
	} else if !heaven && next.PositionType == Heaven {
		return false
	}

	return true
}

func (m *Marble) swap(other *Marble) error {
	if !m.canSwapWith(other) {
		return fmt.Errorf("can not swap marble %s with marble %s", m, other)
	}
	m.Position, other.Position = other.Position, m.Position
	m.Position.Marble, other.Position.Marble = other.Position.Marble, m.Position.Marble
	return nil
}

// Places a marble in a position, if there is another marble there, returns an error and the marble.
// If the force flag is set, the marble will be send home, returned and no error is thrown
func (marble *Marble) place(position *Position, force bool) (*Marble, error) {
	existing := position.Marble
	if force && existing != nil {
		existing.sendHome()
	} else if existing != nil {
		return existing, fmt.Errorf("position %s already taken by %s", position, existing)
	}

	marble.Position, position.Marble = position, marble
	return existing, nil
}

func (m *Marble) canSwapWith(other *Marble) bool {
	if m.Player == other.Player {
		return false
	}

	return m.canSwap() && other.canSwap()
}

func (m *Marble) canSwap() bool {
	if m.Position.PositionType == Home || m.Position.PositionType == Heaven {
		return false
	}

	return !m.isBlocking()
}

func (m *Marble) max(cutoff int) int {
	if m.Position.PositionType == Home {
		return 0
	}

	next := m.Position
	for i := range cutoff {
		// no need to check Heaven: position 1,2,3,4 are never blocking. So it comes to the same result
		next = next.NextPosition
		if next.Marble != nil && next.Marble.isBlocking() {
			return i
		}
	}

	return cutoff
}

func (m *Marble) move(steps int, heaven, seven bool) error {
	backwards := steps < 0
	if backwards && heaven {
		return fmt.Errorf("can not walk backwards into heaven")
	}

	next := m.Position
	startTouches := m.StartTouches

	if backwards && !next.onBoard() {
		return fmt.Errorf("can not walk backwards when in heaven")
	}

	abs := int(math.Abs(float64(steps)))
	var via []*Position
	for range abs {
		if heaven && next.AltNextPosition != nil && next.Section == m.Player.Section {
			if startTouches < 2 {
				return fmt.Errorf("can not walk %s into heaven before touching start at least twice", m)
			}
			next = next.AltNextPosition
		} else if steps > 0 {
			if next.NextPosition == nil {
				return fmt.Errorf("can't walk %s past top of heaven", m)
			}

			next = next.NextPosition
		} else {
			next = next.LastPosition
		}

		if next.Marble != nil && next.Marble.isBlocking() {
			return fmt.Errorf("Marble %s is blocked by %s", m, next.Marble)
		}

		if next.PositionType == Start && next.Section == m.Player.Section {
			startTouches++
		}
		via = append(via, next)
	}

	if heaven && next.PositionType != Heaven {
		return fmt.Errorf("Heaven flag was set but Marble %s did not land in heaven but is %s", next.Marble, next)
	} else if !heaven && next.PositionType == Heaven {
		return fmt.Errorf("Heaven flag was not set but Marble %s is in heaven at %s", next.Marble, m.Position)
	}

	m.Position.moveFrom(m)
	for _, v := range via {
		v.moveVia(m, seven)
	}
	next.moveTo(m)
	m.StartTouches = startTouches

	return nil
}
