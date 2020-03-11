package game

// Coord type
type Coord struct {
	X int
	Y int
}

// Snake type
type Snake struct {
	ID     string
	Name   string
	Health int
	Body   []Coord
}

// Board type
type Board struct {
	Height int
	Width  int
	Food   []Coord
	Snakes []Snake
}

// Game type
type Game struct {
	ID string
}

// State type
type State struct {
	Game  Game
	Turn  int
	Board Board
	You   Snake
}
