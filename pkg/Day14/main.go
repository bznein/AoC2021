package Day14

import (
	"strings"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
)

type rule map[string]string

const stepsP2 = 40
const stepsP1 = 10

func CopyMap(m map[string]int) map[string]int {
	cp := make(map[string]int)
	for k, v := range m {
		cp[k] = v
	}

	return cp
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0

	ints := input.InputToStringSlice(inputF)

	start := ints[0]

	rules := rule{}
	charOccurrences := map[byte]int{}
	for i := 2; i < len(ints); i++ {
		splitted := strings.Split(ints[i], " -> ")
		rules[splitted[0]] = splitted[1]
	}

	pairs := map[string]int{}

	for j := 0; j < len(start); j++ {
		if j < len(start)-1 {
			pairs[start[j:j+2]]++
		}
		charOccurrences[start[j]]++
	}

	for i := 0; i < stepsP2; i++ {
		tempPairs := CopyMap(pairs)
		for s, v := range tempPairs {
			if r, ok := rules[s]; ok {
				// remove this pair, and add two new
				pairs[s] -= v
				if pairs[s] <= 0 {
					delete(pairs, s)
				}
				pairs[string(s[0])+r] += v
				pairs[r+string(s[1])] += v
				charOccurrences[r[0]] += v
			}
		}
		if i == stepsP1-1 {
			part1 = getAnswer(charOccurrences, start)
		}

	}

	part2 = getAnswer(charOccurrences, start)

	return part1, part2
}

func getAnswer(charOccurrences map[byte]int, start string) int {

	min := charOccurrences[0]
	max := min

	for _, v := range charOccurrences {
		if min == 0 || (v != 0 && v < min) {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return max - min
}
