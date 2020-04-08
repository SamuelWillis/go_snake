package game

import (
	"testing"
)

func TestShouldEatFood(t *testing.T) {
	testCases := getShouldEatFoodTestCases()
	for _, testCase := range(testCases) {
		snake := Snake{
			ID: "420",
			Name: "go_snake_420",
			Health: testCase.Health,
			Body: testCase.Body,
		}
		behaviours := Behaviours{
			You: snake,
		}

		result := behaviours.ShouldEatFood()

		if (result != testCase.Expected) {
			t.Errorf(
				"ShouldEatFood failed with:(Health: %d, Body: %#v, result: %t, Expected: %t)",
				testCase.Health,
				testCase.Body,
				result,
				testCase.Expected,
			)
		}
	}
}

func TestShouldChaseTail(t *testing.T) {
	testCases := getShouldChaseTailTestCases()
	for _, testCase := range(testCases) {
		snake := Snake{
			ID: "420",
			Name: "go_snake_420",
			Health: testCase.Health,
			Body: testCase.Body,
		}
		behaviours := Behaviours{
			You: snake,
		}

		result := behaviours.ShouldChaseTail()

		if (result != testCase.Expected) {
			t.Errorf(
				"ShouldChaseTail failed with:(Health: %d, Body: %#v, result: %t, Expected: %t)",
				testCase.Health,
				testCase.Body,
				result,
				testCase.Expected,
			)
		}
	}
}

type BehavioursTestCase struct {
	Health int
	Body []Coord
	Expected bool
}
func getShouldEatFoodTestCases() []BehavioursTestCase {
	return []BehavioursTestCase{
		// Len < 10 -> eat food.
		{Health: 100, Body: []Coord{Coord{1, 2}}, Expected: true},
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

func getShouldChaseTailTestCases() []BehavioursTestCase {
	return []BehavioursTestCase{
		// len >= 10, health >= 50 -> chase tail
		{Health: 50, Body: []Coord{
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
		// health < 50 -> do not chase tail
		{Health: 49, Body: []Coord{Coord{1, 2}}, Expected: false},
	}
}
