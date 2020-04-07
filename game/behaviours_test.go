package game

import "testing"

func TestShouldEatFood(t *testing.T) {
	tables := getShouldEatFoodTestCases()
	for _, table := range(tables) {
		snake := Snake{
			ID: "420",
			Name: "go_snake_420",
			Health: table.Health,
			Body: table.Body,
		}
		behaviours := Behaviours{
			You: snake,
		}

		result := behaviours.ShouldEatFood()

		if (result != table.Expected) {
			t.Errorf(
				"ShouldEatFood failed with:(Health: %d, Body: %#v, result: %t, Expected: %t",
				table.Health,
				table.Body,
				result,
				table.Expected,
			)
		}
	}
}

type ShouldEatFoodCase struct {
	Health int
	Body []Coord
	Expected bool
}

func getShouldEatFoodTestCases() []ShouldEatFoodCase {
	return []ShouldEatFoodCase{
		// Len < 10 -> eat food.
		{Health: 100, Body: []Coord{ Coord{1, 2} }, Expected: true},
		// len >= 10, Health < 50 -> eat food
		{Health: 49, Body: []Coord{
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
		}, Expected: true},
		// len >= 10, Health > 50 -> eat food
		{Health: 66, Body: []Coord{
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
			Coord{1, 2},
		}, Expected: false},
	}
}
