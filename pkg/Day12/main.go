package Day12

import (
	"strings"
	"time"
	"unicode"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
)

type cave struct {
	name  string
	small bool
}

type path []cave

type links map[string][]cave

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0

	ints := input.InputToStringSlice(inputF)

	system := links{}
	for _, ss := range ints {
		s := strings.Split(ss, "-")
		if _, ok := system[s[0]]; !ok {
			system[s[0]] = []cave{}
		}
		if _, ok := system[s[1]]; !ok {
			system[s[1]] = []cave{}
		}
		isSmall := unicode.IsLower(rune(s[1][0]))
		isSmallTo := unicode.IsLower(rune(s[0][0]))
		if s[1] != "start" {
			system[s[0]] = append(system[s[0]], cave{s[1], isSmall})
		}
		if s[0] != "start" {
			system[s[1]] = append(system[s[1]], cave{s[0], isSmallTo})
		}
	}

	allPaths := allPathsFrom(cave{"start", true}, system, map[string]bool{}, true)

	for _, p := range allPaths {
		if p[len(p)-1].name == "end" {
			part1++
		}
	}

	allPaths = allPathsFrom(cave{"start", true}, system, map[string]bool{}, false)

	for _, p := range allPaths {
		if p[len(p)-1].name == "end" {
			part2++
		}
	}

	return part1, part2
}

func CopyMap(m map[string]bool) map[string]bool {
	cp := make(map[string]bool)
	for k, v := range m {
		cp[k] = v
	}

	return cp
}

func allPathsFrom(from cave, system links, visited map[string]bool, alreadyDoubleVisited bool) []path {
	if from.name == "end" {
		return []path{{from}}
	}
	if from.small {
		if _, ok := visited[from.name]; ok {
			if !alreadyDoubleVisited {
				alreadyDoubleVisited = true
			} else {
				return []path{}
			}
		}
		visited[from.name] = true
	}

	paths := []path{}
	for _, c := range system[from.name] {
		for _, p := range allPathsFrom(c, system, CopyMap(visited), alreadyDoubleVisited) {
			pp := path{from}
			pp = append(pp, p...)
			paths = append(paths, pp)
		}
	}
	if len(paths) == 0 {
		return []path{{from}}
	}

	return paths
}
