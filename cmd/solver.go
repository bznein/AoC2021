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
	"github.com/bznein/AoC2021/pkg/Day08"
	"github.com/bznein/AoC2021/pkg/Day09"
	"github.com/bznein/AoC2021/pkg/Day10"
	"github.com/bznein/AoC2021/pkg/Day11"
	"github.com/bznein/AoC2021/pkg/Day12"
	"github.com/bznein/AoC2021/pkg/Day13"
	"github.com/bznein/AoC2021/pkg/Day14"
	"github.com/bznein/AoC2021/pkg/Day15"
	"github.com/bznein/AoC2021/pkg/Day16"
	"github.com/bznein/AoC2021/pkg/Day17"
	"github.com/bznein/AoC2021/pkg/timing"

	"github.com/bznein/AoC2021/pkg/input"
)

const days = 17

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
	Day08.Solve(inputs[7])
	Day09.Solve(inputs[8])
	Day10.Solve(inputs[9])
	Day11.Solve(inputs[10])
	Day12.Solve(inputs[11])
	Day13.Solve(inputs[12])
	Day14.Solve(inputs[13])
	Day15.Solve(inputs[14])
	Day16.Solve(inputs[15])
	Day17.Solve(inputs[16])
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
	case 8:
		return Day09.Solve(inputF)
	case 9:
		return Day09.Solve(inputF)
	case 10:
		return Day10.Solve(inputF)
	case 11:
		return Day11.Solve(inputF)
	case 12:
		return Day12.Solve(inputF)
	case 13:
		return Day13.Solve(inputF)
	case 14:
		return Day14.Solve(inputF)
	case 15:
		return Day15.Solve(inputF)
	case 16:
		return Day16.Solve(inputF)
	case 17:
		return Day17.Solve(inputF)
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
