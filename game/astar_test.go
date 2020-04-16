package game

import (
	"testing"
)

func TestGetShortestPath(t *testing.T) {
	// start := Coord{
	// 	X: 0,
	// 	Y: 0,
	// }

	// end := Coord{
	// 	X: 2,
	// 	Y: 2,
	// }

	board := Board{
		Height: 5,
		Width: 5,
		Food: []Coord{},
		Snakes: []Snake{
			Snake{
				ID: "id",
				Name: "test_snake",
				Health: 400,
				Body: []Coord{
					Coord{1,2},
					Coord{2,2},
				},
			},
		},
	}

	aStarBoard := createAStarBoard(board)

	t.Logf("aStarBoard: %#v", aStarBoard)
}

type ShortestPathTestCase struct {
}

func getShortestPathTestCases() []ShortestPathTestCase {
	// no blockers
	// blockers
	return []AStarTestCase{
	}
}
