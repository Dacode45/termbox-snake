package game

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

//CFOOD Color of food
const CFood = termbox.ColorRed

// Food object in game
type Food struct {
	Char rune
	Pos  Vec2i
}

// Creates a food randomly in bounds
func NewFood(width, height int) Food {
	return Food{
		Char: '@',
		Pos: Vec2i{
			X: rand.Intn(width),
			Y: rand.Intn(height),
		},
	}
}
