package Day07

import (
	"math"
	"sort"
	"strings"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
)

const (
	p1days = 18
	p2days = 256
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func distsFromTarget(target int, inputs []int) (int, int) {
	dist1, dist2 := 0, 0
	for _, v := range inputs {
		absDist := abs(v - target)
		dist1 += absDist
		dist2 += (absDist * (absDist + 1)) / 2
	}
	return dist1, dist2
}

func dist2FromTarget(target int, inputs []int) int {
	dist := 0
	for _, v := range inputs {
		dist += abs(v - target)
	}
	return dist
}
func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	intsStrList := strings.Split(input.InputToStringSlice(inputF)[0], ",")

	positions := []int{}
	for _, v := range intsStrList {
		positions = append(positions, input.AsInt(v))
	}

	minDist1, minDist2 := math.MaxInt64, math.MaxInt64
	sort.Ints(positions)
	for i := 0; i < positions[len(positions)-1]; i++ {
		dist1, dist2 := distsFromTarget(i, positions)
		if dist1 < minDist1 {
			minDist1 = dist1
		}
		if dist2 < minDist2 {
			minDist2 = dist2
		}
	}

	return minDist1, minDist2
}
