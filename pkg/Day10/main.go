package Day10

import (
	"sort"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
	"github.com/golang-collections/collections/stack"
)

var mappingP = map[rune]rune{
	')': '(',
	']': '[',
	'>': '<',
	'}': '{',
}

var scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionScores = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func illegalScore(in string) (int, *stack.Stack) {
	s := stack.New()
	for _, c := range in {
		switch c {
		case '(', '[', '{', '<':
			s.Push(c)
		case ')', ']', '}', '>':
			t := s.Pop().(rune)
			if mappingP[c] != t {
				return scores[c], s
			}
		}
	}
	return 0, s
}

func completionScore(in string, s *stack.Stack) int {
	score := 0
	for s.Len() > 0 {
		c := s.Pop().(rune)
		score *= 5
		score += completionScores[c]
	}
	return score
}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1 := 0
	str := input.InputToStringSlice(inputF)

	completionScores := make([]int, 0, len(str))
	for _, s := range str {
		score, remainder := illegalScore(s)
		part1 += score
		if score == 0 {
			completionScores = append(completionScores, completionScore(s, remainder))
		}
	}

	sort.Ints(completionScores)
	return part1, completionScores[len(completionScores)/2]
}
