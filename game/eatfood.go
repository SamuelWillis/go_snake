package game

import(
	"math"
)

func getMoveToFood(state State) string {
	move := ""
	snakeBody := state.You.Body
	head := snakeBody[0]
	food := state.Board.Food

	min := math.Inf(0)
	var minFood Coord

	for i := 0; i < len(food); i++ {
		dist :=  math.Abs(float64(food[i].X - head.X)) + math.Abs(float64(food[i].Y - head.Y))

		if dist <= min {
			min = dist
			minFood = food[i]
		}
	}

	if minFood != (Coord{}) {
		move = getDirection(state, minFood)
	}

	return move
}

func getDirection(state State, minFood Coord) (direction string) {
	head := state.You.Body[0]
	snakes := state.Board.Snakes

	validMoves := ValidMoves{
		up: isValidMove(Coord{ X: head.X, Y: head.Y - 1 }, snakes),
		down: isValidMove(Coord{ X: head.X, Y: head.Y + 1 }, snakes),
		left: isValidMove(Coord{ X: head.X - 1, Y: head.Y }, snakes),
		right: isValidMove(Coord{ X: head.X + 1, Y: head.Y }, snakes),
	}

	if head.X < minFood.X && validMoves.right {
		direction = "right"
	}

	if head.X > minFood.X && validMoves.left {
		direction = "left"
	}

	if head.Y < minFood.Y && validMoves.down {
		direction = "down"
	}

	if head.Y > minFood.Y && validMoves.up {
		direction = "up"
	}

	return
}

// ValidMoves type
type ValidMoves struct {
	up bool
	down bool
	left bool
	right bool
}

func isValidMove(test Coord, snakes []Snake) bool {
	isValid := true
	for i := 0; i < len(snakes); i++ {
		body := snakes[0].Body
		for j := 0; j < len(body); j++ {
			if (test == body[i]) {
				isValid = false
				break
			}
		}
	}
	return isValid
}
