package Day02

import (
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0

	depth := []int{0, 0}
	pos := 0
	aim := 0
	for _, s := range input.InputToSpaceSplittedStringSlice(inputF) {
		command := s[0][0]

		l := input.AsInt(s[1])
		switch command {
		case 'f':
			pos += l
			depth[1] += (aim * l)
		case 'd':
			depth[0] += l
			aim += l
		case 'u':
			depth[0] -= l
			aim -= l
		}
	}

	part1 = depth[0] * pos
	part2 = depth[1] * pos
	return part1, part2
}
