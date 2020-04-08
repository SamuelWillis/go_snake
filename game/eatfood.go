package game

import (
	"math"
)

type EatFood struct {
	You  Snake
	Food []Coord
}

func (EatFood EatFood) GetMoveToFood() Coord {
	var minFood Coord

	snakeBody := EatFood.You.Body
	head := snakeBody[0]
	food := EatFood.Food

	// If only one food return it.
	if len(food) == 1 {
		return food[0]
	}

	min := math.Inf(0)

	for i := 0; i < len(food); i++ {
		xDist := math.Abs(float64(food[i].X - head.X))
		yDist := math.Abs(float64(food[i].Y - head.Y))
		dist := math.Abs(xDist + yDist)

		if dist <= min {
			min = dist
			minFood = food[i]
		}
	}

	return minFood
}
