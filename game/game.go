package game

import(
	"encoding/json"
	"log"
	"math"
	"github.com/SamuelWillis/go_snake/api"
)

// GetNextMove that our snake will do.
func GetNextMove(decoded api.SnakeRequest) string {
	move := "left"

	snakeBody := decoded.You.Body
	head := snakeBody[0]
	food := decoded.Board.Food

	min := math.Inf(0)
	var minFood api.Coord

	for i := 0; i < len(food); i++ {
		dist :=  math.Abs(float64(food[i].X - head.X)) + math.Abs(float64(food[i].Y - head.Y))

		if dist <= min {
			min = dist
			minFood = food[i]
		}
	}

	if minFood != (api.Coord{}) {
		move = getDirection(snakeBody, minFood)
	}

	return move
}

func getDirection(snakeBody []api.Coord, minFood api.Coord) (direction string) {
	head := snakeBody[0]
	if head.X < minFood.X {
		direction = "right"
	}

	if head.X > minFood.X {
		direction = "left"
	}

	if head.Y < minFood.Y {
		direction = "down"
	}

	if head.Y > minFood.Y {
		direction = "up"
	}

	return
}

func dump(obj interface{}) {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		log.Printf(string(data))
	}
}
