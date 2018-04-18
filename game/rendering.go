package game

import (
	"github.com/nsf/termbox-go"
)

// CFG foreground color
const CFG = termbox.ColorDefault

// CBG backgorund color
const CBG = termbox.ColorDefault

func tbString(x, y int, fg, bg termbox.Attribute, msg string) {
	originalX := x
	for _, ch := range msg {
		if ch == '\n' || ch == '\r' {
			y++
			x = originalX
			continue
		}
		termbox.SetCell(x, y, ch, fg, bg)
		x++
	}
}
