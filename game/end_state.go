package game

import (
	"time"

	"github.com/nsf/termbox-go"
)

// EndState handles th egame ending
type EndState struct {
	Won bool
}

// NewEndState creates a new state
func NewEndState(won bool) *EndState {
	return &EndState{
		Won: won,
	}
}

// Update does nothing
func (s *EndState) Update(g *Game, t time.Duration) error {
	return nil
}

// HandleInput does nothing
func (s *EndState) HandleInput(g *Game, input InputEvent) error {
	return nil
}

// Render does nothing
func (s *EndState) Render(g *Game, t time.Duration) error {
	g.Render()
	msg := "You Lost!"
	if s.Won {
		msg = "You Won!"
	}
	tbString(0, 0, CFG, CBG, msg)
	return termbox.Flush()
}
