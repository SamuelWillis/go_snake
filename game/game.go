package game

import(
	"math/rand"
	"time"
	"github.com/SamuelWillis/go_snake/helpers"
)

// GetNextMove that our snake will do.
func GetNextMove(state State) string {
	var move string
	validMoves := getValidMoves(state)
	behaviours := Behaviours{
		State: state,
	}

	if behaviours.ShouldEatFood() {
		move = EatFood{
			State: state,
			ValidMoves: validMoves,
		}.getMoveToFood()
	}

	if !behaviours.ShouldEatFood() && behaviours.ShouldChaseTail() {
		move = ChaseTail{
			State: state,
			ValidMoves: validMoves,
		}.getMoveToTail()
	}

	if move == "" {
		move = chooseRandomMove(validMoves)
		helpers.Dump("choosing random move", move)
	}

	return move
}

func getValidMoves(state State) ValidMoves {
	head := state.You.Body[0]

	validMoves := ValidMoves{
		up: isValidMove(Coord{ X: head.X, Y: head.Y - 1 }, state),
		down: isValidMove(Coord{ X: head.X, Y: head.Y + 1 }, state),
		left: isValidMove(Coord{ X: head.X - 1, Y: head.Y }, state),
		right: isValidMove(Coord{ X: head.X + 1, Y: head.Y }, state),
	}

	helpers.Dump("up", validMoves.up)
	helpers.Dump("down", validMoves.down)
	helpers.Dump("left", validMoves.left)
	helpers.Dump("right", validMoves.right)

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

	for i := 0; i < len(snakes); i++ {
		body := snakes[i].Body
		for j := 0; j < len(body); j++ {
			if (test == body[j]) {
				isValid = false
			}
		}
	}
	return isValid
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

	return moves[rand.Intn(len(moves))]
}
