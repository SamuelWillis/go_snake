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

	food := decoded.Board.Food

	head := decoded.You.Body[0]

	min := math.Inf(0)
	var min_food api.Coord
	for i := 0; i < len(food); i++ {
		dist :=  math.Abs(float64(food[i].X - head.X)) + math.Abs(float64(food[i].Y - head.Y))

		if dist <= min {
			min = dist
			min_food = food[i]
		}
	}

	if min_food != (api.Coord{}) {
		move = getDirection(head, min_food)
	}

	return move
}

func getDirection(head api.Coord, minFood api.Coord) (direction string) {
	if head.X <= minFood.X {
		direction = "right"
	}

	if head.X > minFood.X {
		direction = "left"
	}

	if head.Y <= minFood.Y {
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
