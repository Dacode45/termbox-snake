package game

import (
	"github.com/nsf/termbox-go"
)

// InputType tells what the input should be handled as. Up, down Action
type InputType int

const (
	NoOp InputType = iota
	IUp
	IDown
	ILeft
	IRight
	IConfirm
	IBack
	IEsc
)

// InputEvent maps keyboard to universal event
type InputEvent struct {
	Input InputType
	Event termbox.Event
}

// GetInput maps termbox to universal event InputEvent
func GetInput(ev termbox.Event) InputEvent {
	input := NoOp
	switch ev.Key {
	case termbox.KeyArrowUp:
		input = IUp
	case termbox.KeyArrowDown:
		input = IDown
	case termbox.KeyArrowLeft:
		input = ILeft
	case termbox.KeyArrowRight:
		input = IRight
	case termbox.KeySpace:
		input = IConfirm
	case termbox.KeyBackspace:
		fallthrough
	case termbox.KeyBackspace2:
		input = IBack
	case termbox.KeyEsc:
		input = IEsc
	}
	return InputEvent{
		Event: ev,
		Input: input,
	}
}

func pollKeyboard(send chan<- InputEvent, done <-chan interface{}) {
	termbox.SetInputMode(termbox.InputEsc)
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			input := GetInput(ev)
			select {
			case <-done:
				break mainloop
			case send <- input:
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
