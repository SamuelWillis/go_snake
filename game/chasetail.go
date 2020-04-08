package game

type ChaseTail struct {
	You Snake
}

func (ChaseTail ChaseTail) GetMoveToTail() Coord {
	body := ChaseTail.You.Body
	tailCoord := body[len(body) - 1]

	return tailCoord
}
