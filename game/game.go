package game

import(
	"math/rand"
	"time"
	"github.com/SamuelWillis/go_snake/helpers"
)

// GetNextMove that our snake will do.
func GetNextMove(state State) string {
	var coordToMoveTo Coord

	validMoves := getValidMoves(state)
	behaviours := Behaviours{
		You: state.You,
	}

	if behaviours.ShouldEatFood() {
		coordToMoveTo = EatFood{
			You: state.You,
			Food: state.Board.Food,
		}.GetMoveToFood()
		helpers.Dump("Eating Food", coordToMoveTo)
	}

	if behaviours.ShouldChaseTail() {
		coordToMoveTo = ChaseTail{
			You: state.You,
		}.GetMoveToTail()
		helpers.Dump("Moving to Tail", coordToMoveTo)
	}

	// Safety Net. If no coord to move to choose a random move.
	if coordToMoveTo == (Coord{}) {
		return ""
	}

	helpers.Dump("possible validMoves", validMoves)

	return getDirectionToCoord(coordToMoveTo, state, validMoves)
}

func getDirectionToCoord(coord Coord, state State, validMoves ValidMoves) string {
	direction := ""

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

	// if we cannot move in a direction we want to move
	// pick a random one.
	if (direction == "") {
		direction = chooseRandomMove(validMoves)
	}

	return direction
}

func chooseRandomMove(validMoves ValidMoves) string {
	rand.Seed(time.Now().Unix())
	var moves []string

	if validMoves.up {
		moves = append(moves, "up")
	}

	if validMoves.down {
		moves = append(moves, "down")
	}

	if validMoves.left {
		moves = append(moves, "left")
	}

	if validMoves.right {
		moves = append(moves, "right")
	}

	if len(moves) == 0 {
		return ""
	}

	return moves[rand.Intn(len(moves))]
}

func getValidMoves(state State) ValidMoves {
	head := state.You.Body[0]

	validMoves := ValidMoves{
		up: isValidMove(Coord{ X: head.X, Y: head.Y - 1 }, state),
		down: isValidMove(Coord{ X: head.X, Y: head.Y + 1 }, state),
		left: isValidMove(Coord{ X: head.X - 1, Y: head.Y }, state),
		right: isValidMove(Coord{ X: head.X + 1, Y: head.Y }, state),
	}

	return validMoves
}

func isValidMove(test Coord, state State) bool {
	isValid := true
	snakes := state.Board.Snakes

	// out of bounds
	if test.X < 0 || test.X >= state.Board.Width {
		return false
	}

	// out of bounds
	if test.Y < 0 || test.Y >= state.Board.Height {
		return false
	}

	// will hit snake body.
	for i := 0; i < len(snakes); i++ {
		body := snakes[i].Body
		// ignore the snake head.
		// Hitting the head could be a valid move.
		for j := 1; j < len(body); j++ {
			if (test == body[j]) {
				isValid = false
			}
		}
	}
	return isValid
}
