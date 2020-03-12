package game

// GetNextMove that our snake will do.
func GetNextMove(state State) string {
	var move string

	if ShouldEatFood() {
		move = getMoveToFood(state)
	}

	if ShouldAttack() {
		move = getMoveToFood(state)
	}

	if move == "" {
		move = "left"
	}

	return move
}
