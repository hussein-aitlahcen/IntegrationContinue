package bowling

import (
	"fmt"
	"testing"
)

func ensureValidScore(t *testing.T, input []Frame, expectedScore int, expectedError error) {
	score, err := GetScore(input)
	if err != expectedError && !(err != nil && expectedError != nil && err.Error() == expectedError.Error()) {
		t.Fatalf("Score error : %+v, expected %+v", err, expectedError)
	}
	if score != expectedScore {
		t.Fatalf("Score : %d, expected %d", score, expectedScore)
	}
}

func TestScore(t *testing.T) {
	input := []Frame{{9, 1}, {0, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {0, 0}, {0, 0}}
	expected := 27
	ensureValidScore(t, input, expected, nil)
}

func TestStrikeScore(t *testing.T) {
	input := []Frame{{10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}}
	expected := 300
	ensureValidScore(t, input, expected, nil)

	input = []Frame{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {10, 0}, {10, 0}, {0, 0}, {0, 0}}
	expected = 30
	ensureValidScore(t, input, expected, nil)
}

func TestSpareScore(t *testing.T) {
	input := []Frame{{5, 5}, {5, 3}, {5, 3}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expected := 31
	ensureValidScore(t, input, expected, nil)

	input = []Frame{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {5, 5}, {5, 5}, {0, 0}, {0, 0}}
	expected = 25
	ensureValidScore(t, input, expected, nil)
}

func TestScoreArraySize(t *testing.T) {
	input := []Frame{{1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {0, 0}, {0, 0}, {0, 0}}
	expectedScore := 0
	expectedError := fmt.Errorf(ErrFrameOverflow, MaxFrame)
	ensureValidScore(t, input, expectedScore, expectedError)

	input = []Frame{{1, 1}}
	expectedScore = 0
	expectedError = fmt.Errorf(ErrFrameOverflow, MaxFrame)
	ensureValidScore(t, input, expectedScore, expectedError)
}

func TestNegativeScoreValue(t *testing.T) {
	input := []Frame{{-1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {0, 0}, {0, 0}}
	expectedScore := 0
	expectedError := fmt.Errorf(ErrNegativeFrameScore)
	ensureValidScore(t, input, expectedScore, expectedError)
}

func TestScoreTupleLimit(t *testing.T) {
	input := []Frame{{10, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {0, 0}, {0, 0}}
	expectedScore := 0
	expectedError := fmt.Errorf(ErrFrameSumInvalid, StrikeScore)
	ensureValidScore(t, input, expectedScore, expectedError)
}
