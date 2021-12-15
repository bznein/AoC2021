package Day15

import (
	"strconv"
	"time"

	"github.com/RyanCarrier/dijkstra"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0

	ints := input.InputToStringSlice(inputF)

	part1 = solveForGraph(ints)

	originalL := len(ints)
	// P2
	ints2 := make([]string, len(ints))
	copy(ints2, ints)

	for i, r := range ints2 {
		for j := 1; j < 5; j++ {
			ints2[i] += increaseString(r, j)
		}
	}

	for i := 1; i < 5; i++ {
		for j := 0; j < originalL; j++ {
			ints2 = append(ints2, increaseString(ints2[j], i))
		}
	}

	part2 = solveForGraph(ints2)

	// Uncomment this to print the best path
	// for i, v := range ints2 {
	// 	for j, c := range v {
	// 		if contains(best.Path, getIdx(i, j, w)) {
	// 			fmt.Print(color.Ize(color.Red, string(c)))
	// 		} else {
	// 			fmt.Print(string(c))
	// 		}
	// 	}
	// 	fmt.Println()

	// }

	return part1, part2
}

func solveForGraph(ints []string) int {
	g := dijkstra.NewGraph()
	w := len(ints[0])

	for i, r := range ints {
		for j, _ := range r {
			curIdx := getIdx(i, j, w)
			g.AddVertex(curIdx)
		}
	}
	for i, r := range ints {
		for j, c := range r {
			curIdx := getIdx(i, j, w)
			// Add arc for everyone entering in it
			// Up
			if i > 0 {
				g.AddArc(getIdx(i-1, j, w), curIdx, int64(c-'0'))
			}
			//Down
			if i < len(ints)-1 {
				g.AddArc(getIdx(i+1, j, w), curIdx, int64(c-'0'))
			}
			//Left
			if j > 0 {
				g.AddArc(getIdx(i, j-1, w), curIdx, int64(c-'0'))
			}
			//Right
			if j < w-1 {
				g.AddArc(getIdx(i, j+1, w), curIdx, int64(c-'0'))
			}
		}
	}

	best, _ := g.Shortest(0, getIdx(len(ints)-1, w-1, w))
	return int(best.Distance)
}

func getIdx(i, j int, w int) int {
	return i*w + j
}

func contains(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func increaseString(s string, i int) string {
	ret := ""
	for _, c := range s {
		v := int(c - '0')
		v += i
		if v > 9 {
			v -= 9
		}
		ret += strconv.Itoa(v)
	}
	return ret
}
