package game

import(
	"math"
)

// GetNextMove that our snake will do.
func GetNextMove(state State) string {
	move := "left"

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
	validMoves := getValidMoves(state.You.Body)

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

func getValidMoves(snake []Coord) ValidMoves {
	head := snake[0]
	body := snake[1:]
	
	return ValidMoves{
		up: isValid(Coord{ X: head.X, Y: head.Y + 1 }, body),
		down: isValid(Coord{ X: head.X, Y: head.Y - 1 }, body),
		left: isValid(Coord{ X: head.X - 1, Y: head.Y }, body),
		right: isValid(Coord{ X: head.X + 1, Y: head.Y }, body),
	}
}

func isValid(test Coord, body []Coord) bool {
	isValid := true
	for i := 0; i < len(body); i++ {
		if (test == body[i]) {
			isValid = false
			break
		}
	}
	return isValid
}
