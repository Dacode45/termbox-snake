package game

// Snake is the protagonixt of the game
type Snake struct {
	Body     []Vec2i
	length   int
	justBorn bool
}

// NewSnake creates a new snake
func NewSnake(length int, start Vec2i) *Snake {
	body := make([]Vec2i, length)
	for i := range body {
		body[i].X = start.X
		body[i].Y = start.Y
	}
	return &Snake{
		Body:     body,
		length:   length,
		justBorn: true,
	}
}

// Head returns the head of the snake a a copy
func (s *Snake) Head() Vec2i {
	return s.Body[len(s.Body)-1]
}

// Move moves snake in direction, maintaining current length
func (s *Snake) Move(dir Direction) {
	if dir == 0 {
		return
	}
	s.justBorn = false
	current := s.Head()
	next := Vec2i{current.X, current.Y}
	switch dir {
	case Up:
		next.Y--
	case Down:
		next.Y++
	case Left:
		next.X--
	case Right:
		next.X++
	}
	s.Body = append(s.Body[1:], next)
}

// KilledSelf returns true if snake eats self
func (s *Snake) KilledSelf() bool {
	if s.justBorn {

		return false
	}
	body := s.Body[:len(s.Body)-1]
	head := s.Head()
	for _, part := range body {
		if part == head {
			return true
		}
	}
	return false
}

// Expand adds on to end of snake
func (s *Snake) Expand() {
	s.Body = append([]Vec2i{s.Body[0]}, s.Body...)
}
