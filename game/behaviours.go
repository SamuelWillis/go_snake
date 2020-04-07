package game

type Behaviours struct {
	You Snake
}

func (Behaviours Behaviours) ShouldEatFood() bool {
	health := Behaviours.You.Health
	body := Behaviours.You.Body
	return len(body) < 10 || health < 50
}

func (Behaviours Behaviours) ShouldChaseTail() bool {
	health := Behaviours.You.Health
	return health >= 50
}
