package Day01

import (
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0
	ints := input.InputToIntSlice(inputF)

	for i := 1; i < len(ints); i++ {
		if ints[i] > ints[i-1] {
			part1++
		}
		if i > 2 && ints[i] > ints[i-3] {
			part2++
		}
	}

	return part1, part2
}
