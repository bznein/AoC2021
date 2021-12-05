package Day05

import (
	"strings"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
)

type point struct {
	i int
	j int
}

type line struct {
	startPoint point
	endPoint   point
}

func (l line) isVertical() bool {
	return l.startPoint.i == l.endPoint.i
}

func (l line) isHorizontal() bool {
	return l.startPoint.j == l.endPoint.j
}

func (l line) isDiagonal() bool {
	return !l.isVertical() && !l.isHorizontal()
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (l line) horizontalPath() []point {
	path := []point{}
	dist := abs(l.endPoint.i - l.startPoint.i)
	if l.startPoint.i < l.endPoint.i {
		for i := 0; i <= dist; i++ {
			path = append(path, point{
				i: l.startPoint.i + i,
				j: l.startPoint.j,
			})
		}
	} else {
		for i := 0; i <= dist; i++ {
			path = append(path, point{
				i: l.startPoint.i - i,
				j: l.startPoint.j,
			})
		}
	}
	return path
}

func (l line) verticalPath() []point {
	path := []point{}
	dist := abs(l.endPoint.j - l.startPoint.j)
	if l.startPoint.j < l.endPoint.j {
		for i := 0; i <= dist; i++ {
			path = append(path, point{
				i: l.startPoint.i,
				j: l.startPoint.j + i,
			})
		}
	} else {
		for i := 0; i <= dist; i++ {
			path = append(path, point{
				i: l.startPoint.i,
				j: l.startPoint.j - i,
			})
		}
	}
	return path
}

func (l line) diagonalPath() []point {
	path := []point{}
	// 4 cases
	if l.startPoint.i <= l.endPoint.i && l.startPoint.j <= l.endPoint.j {
		// 0,0 -> 8,8
		dist := l.endPoint.i - l.startPoint.i
		for i := 0; i <= dist; i++ {
			path = append(path, point{
				i: l.startPoint.i + i,
				j: l.startPoint.j + i,
			})
		}
	}

	if l.startPoint.i > l.endPoint.i && l.startPoint.j > l.endPoint.j {
		// 8,8 -> 0,0
		dist := l.startPoint.i - l.endPoint.i
		for i := 0; i <= dist; i++ {
			path = append(path, point{
				i: l.startPoint.i - i,
				j: l.startPoint.j - i,
			})
		}
	}

	if l.startPoint.i < l.endPoint.i && l.startPoint.j > l.endPoint.j {
		// 0,8 -> 8,0
		dist := l.endPoint.i - l.startPoint.i
		for i := 0; i <= dist; i++ {
			path = append(path, point{
				i: l.startPoint.i + i,
				j: l.startPoint.j - i,
			})
		}
	}
	if l.startPoint.i > l.endPoint.i && l.startPoint.j < l.endPoint.j {
		// 8,0 -> 0,8
		dist := l.startPoint.i - l.endPoint.i
		for i := 0; i <= dist; i++ {
			path = append(path, point{
				i: l.startPoint.i - i,
				j: l.startPoint.j + i,
			})
		}
	}
	return path
}

func (l line) path() []point {
	if l.isHorizontal() {
		return l.horizontalPath()
	}
	if l.isVertical() {
		return l.verticalPath()
	}
	return l.diagonalPath()
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0
	lines := input.InputToStringSlice(inputF)
	coveredp1 := map[point]int{}
	coveredp2 := map[point]int{}
	for _, l := range lines {
		lineSplitted := strings.Split(l, " -> ")
		p1Splitted := strings.Split(lineSplitted[0], ",")
		p2Splitted := strings.Split(lineSplitted[1], ",")
		p1 := point{
			i: input.AsInt(p1Splitted[0]),
			j: input.AsInt(p1Splitted[1]),
		}
		p2 := point{
			i: input.AsInt(p2Splitted[0]),
			j: input.AsInt(p2Splitted[1]),
		}
		curLine := line{
			startPoint: p1,
			endPoint:   p2,
		}

		for _, p := range curLine.path() {
			if !curLine.isDiagonal() {
				coveredp1[p]++
				if coveredp1[p] == 2 {
					part1++
				}
			}

			coveredp2[p]++
			if coveredp2[p] == 2 {
				part2++
			}

		}

	}
	return part1, part2
}
