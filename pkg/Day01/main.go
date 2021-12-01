package Day01

import (
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1 := 0
	part2 := 0
	ints := input.InputToIntSlice(inputF)
	for i := 1; i < len(ints); i++ {
		if ints[i] > ints[i-1] {
			part1++
		}
	}

	sums := []int{}
	for i := 2; i < len(ints); i++ {
		sums = append(sums, ints[i]+ints[i-1]+ints[i-2])
	}

	for i := 1; i < len(sums); i++ {
		if sums[i] > sums[i-1] {
			part2++
		}
	}
	return part1, part2
}
