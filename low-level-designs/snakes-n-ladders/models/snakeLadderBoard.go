package models

type SnakeLadderBoard struct {
	Size            int
	Ladders         []Ladders
	Snakes          []Snakes
	PlayerPositions map[string]int
}

func NewSnakeLadderBoard(size int) SnakeLadderBoard {
	return SnakeLadderBoard{
		Size:            size,
		Ladders:         make([]Ladders, 0),
		Snakes:          make([]Snakes, 0),
		PlayerPositions: make(map[string]int),
	}
}
