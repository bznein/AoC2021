package Day09

import (
	"sort"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
)

type grid struct {
	data [][]int
	h    int
	w    int
}

func (g *grid) addRow(s string) {
	g.data = append(g.data, []int{})
	i := len(g.data) - 1
	g.h++
	g.w = len(s)
	for _, v := range s {
		g.data[i] = append(g.data[i], input.AsInt(string(v)))
	}
}

func (g grid) getNeighbours(i, j int) []int {
	neighbours := []int{}
	if i > 0 {
		neighbours = append(neighbours, g.data[i-1][j])
	}
	if j > 0 {
		neighbours = append(neighbours, g.data[i][j-1])
	}
	if j < g.w-1 {
		neighbours = append(neighbours, g.data[i][j+1])
	}
	if i < g.h-1 {
		neighbours = append(neighbours, g.data[i+1][j])
	}
	return neighbours
}

func (g grid) getLowPoints() []int {
	lowPoints := []int{}
	for i, row := range g.data {
		for j, val := range row {
			neighbours := g.getNeighbours(i, j)
			low := true
			for _, n := range neighbours {
				if val >= n {
					low = false
					break
				}
			}
			if low {
				lowPoints = append(lowPoints, val)
			}
		}
	}
	return lowPoints
}

func (g grid) riskLevel() int {
	risk := 0
	lowPoints := g.getLowPoints()

	for _, l := range lowPoints {
		risk += 1 + l
	}
	return risk
}

type coord struct {
	i int
	j int
}

func (g grid) expandBasin(i, j int, visited map[coord]bool) int {
	if _, ok := visited[coord{i, j}]; ok {
		return 0
	}
	visited[coord{i, j}] = true
	if i < 0 || i >= g.h || j < 0 || j >= g.w || g.data[i][j] == 9 {
		return 0
	}
	return 1 +
		g.expandBasin(i-1, j, visited) +
		g.expandBasin(i+1, j, visited) +
		g.expandBasin(i, j-1, visited) +
		g.expandBasin(i, j+1, visited)
}

func (g grid) getBasins() []int {
	basins := []int{}

	visited := map[coord]bool{}

	for i, row := range g.data {
		for j, _ := range row {
			basins = append(basins, g.expandBasin(i, j, visited))
		}
	}
	return basins

}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0
	strs := input.InputToStringSlice(inputF)
	g := grid{[][]int{}, 0, 0}

	for _, s := range strs {
		g.addRow(s)
	}

	part1 = g.riskLevel()

	basins := g.getBasins()
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	part2 = basins[0] * basins[1] * basins[2]
	return part1, part2
}
