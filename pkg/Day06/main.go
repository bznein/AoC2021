package Day06

import (
	"strings"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
)

const (
	p1days = 18
	p2days = 256
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1 := 0
	intsStrList := strings.Split(input.InputToStringSlice(inputF)[0], ",")

	fishes := []int{}
	for _, v := range intsStrList {
		fishes = append(fishes, input.AsInt(v))
	}
	timer := make([]int, 9)
	for _, f := range fishes {
		timer[f]++
	}
	for i := 0; i < p2days; i++ {
		firstFish := timer[0]
		for i := 0; i < 8; i++ {
			timer[i] = timer[i+1]
		}
		timer[6] += firstFish
		timer[8] = firstFish

		if i == p1days-1 {
			sum := 0
			for _, f := range timer {
				sum += f
			}
			part1 = sum
		}
	}
	sum := 0
	for _, f := range timer {
		sum += f
	}

	return part1, sum
}
