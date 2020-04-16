package game

const snakeBodyMarker = 1

// Get the shortest path from the start to the end on the provided board.
func getShortestPath(start Coord, end Coord, board Board) []Coord {
	createAStarBoard(board)

	return []Coord{}
}

func createAStarBoard(board Board) [][]int {
	// Allocate the board height.
	aStarBoard := make([][]int, board.Height)

	for y := range aStarBoard {
		// Allocate the board width
		aStarBoard[y] = make([]int, board.Width)
	}

	// fill in enemy snakes
	for _, snake := range board.Snakes {
		for _, snakeBody := range snake.Body {
			aStarBoard[snakeBody.Y][snakeBody.X] = snakeBodyMarker;
		}
	}

	return aStarBoard
}
