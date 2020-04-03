package game

type ChaseTail struct {
	State State
	ValidMoves ValidMoves
}

func (ChaseTail ChaseTail) getMoveToTail() string {
	body := ChaseTail.State.You.Body
	tailCoord := body[len(body) - 1]

	return getDirectionToCoord(ChaseTail.State, tailCoord, ChaseTail.ValidMoves)
}
