package game

import (
	"testing"
)

func TestEatFood(t *testing.T) {
	testCases := getEatFoodTestCases()

	for _, testCase := range testCases {
		result := EatFood{
			You:  testCase.You,
			Food: testCase.Food,
		}.GetMoveToFood()

		if result != testCase.Expected {
			t.Errorf(
				"TestEatFood failed for case: (You %#v, Food %#v, Expected %#v, Result %#v",
				testCase.You,
				testCase.Food,
				testCase.Expected,
				result,
			)
		}
	}
}

type TestEatFoodCase struct {
	You      Snake
	Food     []Coord
	Expected Coord
}

func getEatFoodTestCases() []TestEatFoodCase {
	you := Snake{
		ID:     "123",
		Name:   "go_snake_420",
		Health: 100,
		Body: []Coord{
			Coord{X: 5, Y: 5},
		},
	}

	return []TestEatFoodCase{
		// no food
		{
			You:      you,
			Food:     []Coord{},
			Expected: (Coord{}),
		},
		// one food
		{
			You:      you,
			Food:     []Coord{Coord{X: 5, Y: 6}},
			Expected: Coord{X: 5, Y: 6},
		},
		// diagonal food closer
		{
			You:      you,
			Food:     []Coord{Coord{X: 4, Y: 13}, Coord{X: 6, Y: 6}, Coord{X: 8, Y: 5}},
			Expected: Coord{X: 6, Y: 6},
		},
		// diagonal food farther
		{
			You:      you,
			Food:     []Coord{Coord{X: 5, Y: 6}, Coord{X: 7, Y: 7}},
			Expected: Coord{X: 5, Y: 6},
		},
	}
}
