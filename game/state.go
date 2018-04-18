package game

import "time"

type State interface {
	Update(g *Game, t time.Duration) error
	HandleInput(g *Game, input InputEvent) error
	Render(g *Game, t time.Duration) error
}
