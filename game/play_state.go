package game

import (
	"time"

	"github.com/nsf/termbox-go"
)

// CSnake is color of the snake
const CSnake = termbox.ColorCyan

// PlayState is the game state where the player controls
type PlayState struct {
	Width  int
	Height int

	Snake *Snake
	Food  Food

	SinceLastMove time.Duration
	MoveThreshold time.Duration

	Direction Direction
}

// NewPlayState Creates a new playstate
func NewPlayState() *PlayState {
	return &PlayState{
		Width:         20,
		Height:        15,
		Snake:         NewSnake(5, Vec2i{10, 10}),
		Food:          NewFood(20, 15),
		MoveThreshold: time.Second / 5,
		Direction:     Left,
	}
}

// InBounds Checks if object is in bounds
func (s *PlayState) InBounds(vec Vec2i) bool {
	return vec.X >= 0 && vec.X < s.Width && vec.Y >= 0 && vec.Y < s.Height
}

// Update moves the snake in current direciton
func (s *PlayState) Update(g *Game, t time.Duration) error {
	s.SinceLastMove += t
	if s.SinceLastMove > s.MoveThreshold {
		s.SinceLastMove = 0
		s.Snake.Move(s.Direction)
	}
	if !s.InBounds(s.Snake.Head()) || s.Snake.KilledSelf() {
		g.State = NewEndState(false)
	}
	if s.Snake.Head() == s.Food.Pos {
		s.Snake.Expand()
		s.Food = NewFood(s.Width, s.Height)
	}
	return nil
}

// HandleInput Moves the nsake in given direction
func (s *PlayState) HandleInput(g *Game, input InputEvent) error {
	switch input.Input {
	case IUp:
		s.Direction = Up
	case IDown:
		s.Direction = Down
	case ILeft:
		s.Direction = Left
	case IRight:
		s.Direction = Right
	}
	return nil
}

// Render renders the game: snake
func (s *PlayState) Render(g *Game, t time.Duration) error {
	wnd := g.Render()
	s.RenderWalls(g, t, wnd)
	// render food
	termbox.SetCell(wnd.Left+1+s.Food.Pos.X, wnd.Top+1+s.Food.Pos.Y, s.Food.Char, CFood, CBG)
	for _, v := range s.Snake.Body {
		termbox.SetCell(wnd.Left+1+v.X, wnd.Top+1+v.Y, 'S', CSnake, CBG)
	}
	return termbox.Flush()
}

// RenderWalls renders the Walls
func (s *PlayState) RenderWalls(g *Game, t time.Duration, wnd Window) {
	for x := 0; x < s.Width+2; x++ {
		termbox.SetCell(wnd.Left+x, wnd.Top, '0', CFG, CBG)
		termbox.SetCell(wnd.Left+x, wnd.Top+s.Height+1, '0', CFG, CBG)
	}

	for y := 0; y < s.Height+1; y++ {
		termbox.SetCell(wnd.Left, wnd.Top+y, '0', CFG, CBG)
		termbox.SetCell(wnd.Left+s.Width+1, wnd.Top+y, '0', CFG, CBG)
	}
}
