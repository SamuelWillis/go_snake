package game

import(
	"math"
)

type EatFood struct {
	State State
	ValidMoves ValidMoves
}

func (eatfood EatFood) getMoveToFood() string {
	move := ""
	snakeBody := eatfood.State.You.Body
	head := snakeBody[0]
	food := eatfood.State.Board.Food

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
		move = getDirectionToCoord(eatfood.State, minFood, eatfood.ValidMoves)
	}

	return move
}

func getDirectionToCoord(state State, coord Coord, validMoves ValidMoves) (direction string) {
	head := state.You.Body[0]

	if head.X < coord.X && validMoves.right {
		direction = "right"
	}

	if head.X > coord.X && validMoves.left {
		direction = "left"
	}

	if head.Y < coord.Y && validMoves.down {
		direction = "down"
	}

	if head.Y > coord.Y && validMoves.up {
		direction = "up"
	}

	return
}
