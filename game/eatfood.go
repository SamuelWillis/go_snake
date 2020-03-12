package game

import(
	"math"
)

func getMoveToFood(state State, validMoves ValidMoves) string {
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
		move = getDirection(state, minFood, validMoves)
	}

	return move
}

func getDirection(state State, minFood Coord, validMoves ValidMoves) (direction string) {
	head := state.You.Body[0]

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
