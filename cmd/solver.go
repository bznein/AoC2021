package main

import (
	"fmt"
	"time"

	"github.com/bznein/AoC2020/pkg/Day01"
	"github.com/bznein/AoC2020/pkg/timing"

	"github.com/bznein/AoC2020/pkg/input"
)

func readAllInputs() []string {
	defer timing.TimeTrack(time.Now())
	inputs := make([]string, 25)
	for i := 0; i < 25; i++ {
		inputs[i] = input.ReadInput(fmt.Sprintf("../inputs/%d.txt", i+1))
	}
	return inputs
}

func timeSolves() {
	defer timing.TimeTrack(time.Now())
	inputs := readAllInputs()
	Day01.Solve(inputs[0])
}

func solve(day int) (int, int) {

	inputF := input.ReadInput(fmt.Sprintf("../inputs/%d.txt", day))
	switch day {
	case 1:
		return Day01.Solve(inputF)
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
