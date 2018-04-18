package game

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

// Game is a snake game
type Game struct {
	State      State
	LastUpdate time.Time
	LastRender time.Time
	DebugMode  bool

	done   chan interface{}
	input  chan InputEvent
	update *time.Ticker
	render *time.Ticker
}

// NewGame creates a new game in the playstate
func NewGame(debug bool) *Game {
	return &Game{
		State:     NewPlayState(),
		DebugMode: debug,

		done:   make(chan interface{}),
		input:  make(chan InputEvent),
		update: time.NewTicker(time.Second / 30),
		render: time.NewTicker(time.Second / 30),
	}
}

// Init sets up keyboard and stuff for the game
func (g *Game) Init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	go pollKeyboard(g.input, g.done)

	defer termbox.Close()

	g.LastUpdate = time.Now()
	g.LastRender = time.Now()

mainloop:
	for {
		select {
		case k := <-g.input:
			if k.Input == IEsc {
				break mainloop
			} else if g.State != nil {
				if err := g.State.HandleInput(g, k); err != nil {
					panic(err)
				}
			}
		case u := <-g.update.C:
			if g.State != nil {
				if err := g.State.Update(g, u.Sub(g.LastUpdate)); err != nil {
					panic(err)
				}
			}
			g.LastUpdate = u
		case r := <-g.render.C:
			if g.State != nil {
				if err := g.State.Render(g, r.Sub(g.LastRender)); err != nil {
					panic(err)
				}
			}
			g.LastRender = r
		}
	}
	g.render.Stop()
	g.update.Stop()
	g.done <- struct{}{}
}

// Render clears screen and returns a window object
func (g *Game) Render() Window {
	err := termbox.Clear(CFG, CBG)
	if err != nil {
		panic(err)
	}
	w, h := termbox.Size()
	width := w
	height := h
	if g.DebugMode {
		width = w / 2
	}
	wnd := Window{
		Left:   0,
		Top:    0,
		Width:  width,
		Height: height,
	}

	if g.DebugMode {
		wnd.DLeft = width
		wnd.DTop = 0
		wnd.DWidth = width
		wnd.DHeight = height

		msg := fmt.Sprint("Test\nMultiline\nDebug")
		tbString(wnd.DLeft, wnd.DTop, CFG, CBG, msg)
	}

	return wnd

}
