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
	curWindowSum := ints[2] + ints[1] + ints[0]

	for i := 1; i < len(ints); i++ {
		if ints[i] > ints[i-1] {
			part1++
		}
		if i > 2 {
			tempSum := ints[i] + ints[i-1] + ints[i-2]
			if tempSum > curWindowSum {
				part2++
			}
			curWindowSum = tempSum
		}
	}

	return part1, part2
}
