package game

type Behaviours struct {
	State State
}

func (Behaviours Behaviours) ShouldEatFood() bool {
	health := Behaviours.State.You.Health
	body := Behaviours.State.You.Body
	return len(body) < 10 || health < 50
}

func (Behaviours Behaviours) ShouldChaseTail() bool {
	health := Behaviours.State.You.Health
	return health >= 50
}
