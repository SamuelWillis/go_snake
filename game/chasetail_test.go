package game

import (
	"testing"
)

func TestChaseTail(t *testing.T) {
	body := []Coord{
		Coord{X: 5, Y: 10},
		Coord{X: 6, Y: 10},
		Coord{X: 6, Y: 9},
	}

	you := Snake{
		ID: "123",
		Name: "go_snake_420",
		Health: 95,
		Body: body,
	}

	chasetail := ChaseTail{
		You: you,
	}

	result := chasetail.GetMoveToTail()
	expected := body[len(body) - 1]

	if result != expected {
		t.Errorf(
			"ChaseTail did not return correct coord: (expected: %#v, result %#v)",
			expected,
			result,
		)
	}
}
