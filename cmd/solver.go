package main

import (
	"fmt"
	"time"

	"github.com/bznein/AoC2021/pkg/Day01"
	"github.com/bznein/AoC2021/pkg/Day02"
	"github.com/bznein/AoC2021/pkg/Day03"
	"github.com/bznein/AoC2021/pkg/Day04"
	"github.com/bznein/AoC2021/pkg/Day05"
	"github.com/bznein/AoC2021/pkg/Day06"
	"github.com/bznein/AoC2021/pkg/Day07"
	"github.com/bznein/AoC2021/pkg/timing"

	"github.com/bznein/AoC2021/pkg/input"
)

const days = 7

func readAllInputs() []string {
	defer timing.TimeTrack(time.Now())
	inputs := make([]string, days)
	for i := 0; i < days; i++ {
		inputs[i] = input.ReadInput(fmt.Sprintf("../inputs/%d.txt", i+1))
	}
	return inputs
}

func timeSolves() {
	defer timing.TimeTrack(time.Now())
	inputs := readAllInputs()
	Day01.Solve(inputs[0])
	Day02.Solve(inputs[1])
	Day03.Solve(inputs[2])
	Day04.Solve(inputs[3])
	Day05.Solve(inputs[4])
	Day06.Solve(inputs[5])
	Day07.Solve(inputs[6])
}

func solve(day int) (int, int) {

	inputF := input.ReadInput(fmt.Sprintf("../inputs/%d.txt", day))
	switch day {
	case 1:
		return Day01.Solve(inputF)
	case 2:
		return Day02.Solve(inputF)
	case 3:
		return Day03.Solve(inputF)
	case 4:
		return Day04.Solve(inputF)
	case 5:
		return Day05.Solve(inputF)
	case 6:
		return Day06.Solve(inputF)
	case 7:
		return Day07.Solve(inputF)
	}
	return -1, -1
}

func main() {
	input.ParseFlags()

	if input.Timing {
		timeSolves()
		return
	}

	part1, part2 := solve(input.Day)
	fmt.Printf("Part 1: %d, Part 2:%d\n", part1, part2)

}
