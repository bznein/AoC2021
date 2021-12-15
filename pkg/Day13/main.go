package Day13

import (
	"fmt"
	"strings"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
)

type coords struct {
	x int
	y int
}

type fold struct {
	horizontal bool
	point      int
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0

	points := map[coords]bool{}
	folds := []fold{}

	ints := input.InputToStringSlice(inputF)
	for _, i := range ints {
		s := strings.Split(i, ",")
		if len(s) == 2 {
			points[coords{input.AsInt(s[0]), input.AsInt(s[1])}] = true
		} else if i != "" {
			sSpace := strings.Split(i, "=")
			folds = append(folds, fold{sSpace[0] == "fold along y", input.AsInt(sSpace[1])})
		}
	}

	for i, f := range folds {
		performFold(f, points)
		if i == 0 {
			part1 = len(points)
		}
	}

	printPaper(points)
	return part1, part2
}

func performFold(f fold, points map[coords]bool) {
	// First let's do the horizontal
	if f.horizontal {
		for k, _ := range points {
			if k.y > f.point {
				delete(points, k)
				diff := k.y - f.point
				points[coords{k.x, k.y - (2 * diff)}] = true
			}
		}
	} else {
		for k, _ := range points {
			if k.x > f.point {
				delete(points, k)
				diff := k.x - f.point
				points[coords{k.x - (2 * diff), k.y}] = true
			}
		}
	}
}

func mapToSlice(m map[coords]bool) []coords {
	c := make([]coords, len(m))
	i := 0
	for k, _ := range m {
		c[i] = k
		i++
	}
	return c
}

func printPaper(points map[coords]bool) {
	// Find bounding Box
	vals := mapToSlice(points)
	minX := vals[0].x
	maxX := minX
	minY := vals[0].y
	maxY := minY
	for k, _ := range points {
		if k.x < minX {
			minX = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y < minY {
			minY = k.y
		}
		if k.y > maxY {
			maxY = k.y
		}

	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			_, ok := points[coords{x, y}]
			if ok {
				fmt.Printf("█")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

}
