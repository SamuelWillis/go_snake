package game

type Behaviours struct {
	State State
}

func (behaviours Behaviours) ShouldEatFood() bool {
	health := behaviours.State.You.Health
	return health < 50
}

func (behaviours Behaviours) ShouldChaseTail() bool {
	health := behaviours.State.You.Health
	return health >= 50
}
