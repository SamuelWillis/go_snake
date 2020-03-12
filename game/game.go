package game

// GetNextMove that our snake will do.
func GetNextMove(state State) string {
	var move string

	if shouldEatFood() {
		move = getMoveToFood(state)
	}

	if shouldAttack() {
		move = getMoveToFood(state)
	}

	if move == "" {
		move = "left"
	}

	return move
}
