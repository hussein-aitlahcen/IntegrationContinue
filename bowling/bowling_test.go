package bowling

import (
	"fmt"
	"testing"
)

func scoreChecker(t *testing.T, input []Frame, expectedScore int, expectedError error) {
	score, err := GetScore(input)
	if err != expectedError && !(err != nil && expectedError != nil && err.Error() == expectedError.Error()) {
		t.Fatalf("Score error : %+v, expected %+v", err, expectedError)
	}
	if score != expectedScore {
		t.Fatalf("Score : %d, expected %d", score, expectedScore)
	}
}

func TestScore(t *testing.T) {
	input := []Frame{{9, 1}, {0, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}}
	expected := 27
	scoreChecker(t, input, expected, nil)
}

func TestStrikeScore(t *testing.T) {
	input := []Frame{{10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}}
	expected := 270
	scoreChecker(t, input, expected, nil)

	input = []Frame{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {10, 0}, {10, 0}}
	expected = 30
	scoreChecker(t, input, expected, nil)
}

func TestSpareScore(t *testing.T) {
	input := []Frame{{5, 5}, {5, 3}, {5, 3}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expected := 31
	scoreChecker(t, input, expected, nil)

	input = []Frame{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {5, 5}, {5, 5}}
	expected = 25
	scoreChecker(t, input, expected, nil)
}

func TestScoreArraySize(t *testing.T) {
	input := []Frame{{1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}}
	expectedScore := 0
	expectedError := fmt.Errorf("len must be 10")
	scoreChecker(t, input, expectedScore, expectedError)

	input = []Frame{{1, 1}}
	expectedScore = 0
	expectedError = fmt.Errorf("len must be 10")
	scoreChecker(t, input, expectedScore, expectedError)
}

func TestScoreValuesPositiveOrNull(t *testing.T) {
	input := []Frame{{-1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}}
	expectedScore := 0
	expectedError := fmt.Errorf("value must be positive or null")
	scoreChecker(t, input, expectedScore, expectedError)
}

func TestScoreTupleLimit(t *testing.T) {
	input := []Frame{{10, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}}
	expectedScore := 0
	expectedError := fmt.Errorf("tuple score limit should be <= 10")
	scoreChecker(t, input, expectedScore, expectedError)
}
