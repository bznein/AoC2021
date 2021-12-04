package Day04

import (
	"strings"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
)

type entry struct {
	num    int
	marked bool
}

type board struct {
	numbers [][]entry
}

func (b board) isCompleted() bool {
	// Check all rows
	for i := range b.numbers {
		if b.isLineCompleted(i) {
			return true
		}
		if b.isColCompleted(i) {
			return true
		}
	}
	return false
}

func (b board) isLineCompleted(i int) bool {
	for _, v := range b.numbers[i] {
		if !v.marked {
			return false
		}
	}
	return true
}

func (b board) isColCompleted(i int) bool {
	for _, v := range b.numbers {
		if !v[i].marked {
			return false
		}
	}
	return true
}

func (b *board) markNumber(n int) {
	for i, row := range b.numbers {
		for j, val := range row {
			if val.num == n {
				b.numbers[i][j].marked = true
				return
			}
		}
	}
}

func (b board) score(winningNumber int) int {
	totUnmarked := 0
	for _, row := range b.numbers {
		for _, val := range row {
			if !val.marked {
				totUnmarked += val.num
			}
		}
	}
	return totUnmarked * winningNumber
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1 := 0
	ints := input.InputToSpaceSplittedStringSlice(inputF)

	boards := []board{}
	boardN := -1

	numsToExtract := ints[0]

	var idx int
	for _, line := range ints {
		if len(line) == 1 {
			boardN++
			idx = 0
			boards = append(boards, board{numbers: [][]entry{}})
			boards[boardN].numbers = [][]entry{}
			continue
		}
		boards[boardN].numbers = append(boards[boardN].numbers, []entry{})
		for _, n := range line {
			n = strings.TrimSpace(n)
			if len(n) == 0 {
				continue
			}
			boards[boardN].numbers[idx] = append(boards[boardN].numbers[idx], entry{input.AsInt(n), false})
		}
		idx++

	}

	alreadyWon := make(map[int]bool)

	var lastWinningBoard board
	var lastWinningNumber int
	for _, n := range strings.Split(numsToExtract[0], ",") {
		v := input.AsInt(n)
		for i, b := range boards {
			if _, ok := alreadyWon[i]; ok {
				continue
			}
			b.markNumber(v)
			if b.isCompleted() {
				if part1 == 0 {
					part1 = b.score(v)
				}
				lastWinningBoard = b
				lastWinningNumber = v
				alreadyWon[i] = true
			}
		}
	}

	return part1, lastWinningBoard.score(lastWinningNumber)
}
