package Day02

import (
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/input"
	"github.com/bznein/AoC2020/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0
	strs := input.InputToStringSlice(inputF)

	depth := []int{0, 0}
	pos := 0
	aim := 0
	for _, s := range strs {
		splittedS := strings.Split(s, " ")
		command := splittedS[0][0]

		l, _ := strconv.Atoi(splittedS[1])
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
