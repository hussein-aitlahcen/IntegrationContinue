package bowling

import "fmt"

const StrikeScore = 10
const MaxFrame = 10
const ErrNegativeFrameScore = "value must be positive or null"
const ErrFrameSumInvalid = "tuple score limit should be <= %d"
const ErrFrameOverflow = "len must be %d"

type Frame struct {
	firstThrow  int
	secondThrow int
}

func GetSecondThrow(game []Frame, index int) int {
	if index >= len(game) {
		return 0
	}
	return game[index].secondThrow
}

func GetFirstThrow(game []Frame, index int) int {
	if index >= len(game) {
		return 0
	}
	return game[index].firstThrow
}

func GetStrikeScore(game []Frame, index int) int {
	if index >= len(game) {
		return 0
	}
	frame := game[index]
	if frame.firstThrow < StrikeScore {
		return frame.firstThrow + frame.secondThrow
	}
	return frame.firstThrow + GetFirstThrow(game, index+1)
}

func GetSpareScore(game []Frame, index int) int {
	return GetFirstThrow(game, index)
}

func GetFrameScore(game []Frame, index int) (int, error) {

	if index >= len(game) {
		return 0, nil
	}

	var frame = game[index]

	if frame.firstThrow < 0 || frame.secondThrow < 0 {
		return 0, fmt.Errorf(ErrNegativeFrameScore)
	}

	if frame.firstThrow+frame.secondThrow > StrikeScore {
		return 0, fmt.Errorf(ErrFrameSumInvalid, StrikeScore)
	}

	score := 0
	strike := frame.firstThrow == StrikeScore
	if strike {
		score += GetStrikeScore(game, index+1)
	} else {
		spare := frame.firstThrow+frame.secondThrow == StrikeScore
		if spare {
			score += GetSpareScore(game, index+1)
		}
	}

	score += frame.firstThrow + frame.secondThrow

	return score, nil
}

func GetScore(game []Frame) (int, error) {
	if len(game) != MaxFrame {
		return 0, fmt.Errorf(ErrFrameOverflow, MaxFrame)
	}
	score := 0
	for index := range game {
		frameScore, frameErr := GetFrameScore(game, index)
		if frameErr != nil {
			return 0, frameErr
		}
		score += frameScore
	}
	return score, nil
}
