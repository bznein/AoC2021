package Day03

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	aocMath "github.com/bznein/AoC2021/pkg/math"
	"github.com/bznein/AoC2021/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())

	strs := input.InputToStringSlice(inputF)
	l := len(strs[0])

	occurrences := make([][]int, l)
	for i := range occurrences {
		occurrences[i] = make([]int, 2)
	}

	for _, s := range strs {
		for i, c := range s {
			val := c - '0'
			occurrences[i][val]++
		}
	}

	gamma := 0
	epsilon := 0
	for i := range occurrences {
		if occurrences[i][0] > occurrences[i][1] {
			epsilon += aocMath.PowInt(2, l-i-1)
		} else {
			gamma += aocMath.PowInt(2, l-i-1)
		}
	}

	oxigenStr := filterValues(strs, 0, '1')
	co2Str := filterValues(strs, 0, '0')

	fmt.Printf("Oxigen: %s, cco2: %s\n", oxigenStr, co2Str)
	oxigen, _ := strconv.ParseInt(oxigenStr, 2, 64)
	co2, _ := strconv.ParseInt(co2Str, 2, 64)
	return gamma * epsilon, int(oxigen) * int(co2)
}

func filterValues(vals []string, index int, tieBreaker byte) string {
	mostCommonBit := mostCommon(vals, index)
	wasTie := false
	if mostCommonBit == '2' {
		mostCommonBit = tieBreaker
		wasTie = true
	}
	keep := []string{}
	for _, v := range vals {
		if tieBreaker == '1' && v[index] == mostCommonBit {
			keep = append(keep, v)
		}
		if tieBreaker == '0' {
			if !wasTie && v[index] != mostCommonBit {
				keep = append(keep, v)
			} else if wasTie && v[index] == tieBreaker {
				keep = append(keep, v)
			}
		}
	}
	if len(keep) == 1 {
		return keep[0]
	}
	return filterValues(keep, index+1, tieBreaker)
}

func mostCommon(vals []string, index int) byte {
	occ := []int{0, 0}
	for _, v := range vals {
		occ[v[index]-'0']++
	}
	if occ[0] > occ[1] {
		return '0'
	} else if occ[1] > occ[0] {
		return '1'
	}
	return '2'
}
