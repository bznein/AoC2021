package Day11

import (
	"fmt"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
)

const steps = 100

type coords struct {
	i int
	j int
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0

	grid := make([][]int, 10)
	for i := range grid {
		grid[i] = make([]int, 10)
	}
	ints := input.InputToStringSlice(inputF)

	for i, row := range ints {
		for j, v := range row {
			grid[i][j] = int(v - '0')
		}
	}

	for i := 0; ; i++ {
		flashed := map[coords]bool{}
		// First step: increase and increase if flashed
		for i, row := range grid {
			for j, _ := range row {
				grid[i][j] += 1
				if grid[i][j] > 9 {
					flash(grid, i, j, flashed)
				}
			}
		}
		for i, row := range grid {
			for j, _ := range row {
				if _, ok := flashed[coords{i, j}]; ok {
					grid[i][j] = 0
				}
			}
		}
		if i < steps {
			part1 += len(flashed)
		}

		if len(flashed) == 100 {
			part2 = i + 1
			break
		}

	}
	return part1, part2
}

func flash(g [][]int, i, j int, flashed map[coords]bool) {
	if _, ok := flashed[coords{i, j}]; ok {
		return
	}
	flashed[coords{i, j}] = true

	for _, n := range neighbours(g, i, j) {
		g[n.i][n.j]++
		if g[n.i][n.j] > 9 {
			flash(g, n.i, n.j, flashed)
		}
	}
}

func copyGrid(g [][]int) [][]int {
	grid := make([][]int, 10)
	for i := range grid {
		grid[i] = make([]int, 10)
	}
	for i, row := range g {
		for j, v := range row {
			grid[i][j] = v
		}
	}

	return grid
}

func printGrid(g [][]int) {
	for _, row := range g {
		for _, v := range row {
			fmt.Print(v)
		}
		fmt.Println()
	}
}

func neighbours(g [][]int, i, j int) []coords {
	n := []coords{}
	if j > 0 {
		n = append(n, coords{i, j - 1})
		if i > 0 {
			n = append(n, coords{i - 1, j - 1})
		}
		if i < len(g)-1 {
			n = append(n, coords{i + 1, j - 1})
		}
	}
	if j < len(g)-1 {
		n = append(n, coords{i, j + 1})
		if i > 0 {
			n = append(n, coords{i - 1, j + 1})
		}
		if i < len(g)-1 {
			n = append(n, coords{i + 1, j + 1})
		}
	}
	if i > 0 {
		n = append(n, coords{i - 1, j})
	}
	if i < len(g)-1 {
		n = append(n, coords{i + 1, j})
	}
	return n
}
