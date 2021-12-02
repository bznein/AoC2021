package main

import (
	"testing"
)

var solveTests = []struct {
	day   int
	part1 int
	part2 int
}{
	{1, 1167, 1130},
	{2, 1882980, 1971232560},
}

func TestSolver(t *testing.T) {
	for _, test := range solveTests {
		p1, p2 := solve(test.day)
		if p1 != test.part1 || p2 != test.part2 {
			t.Errorf("Day %d, expected (%d,%d), got (%d,%d)", test.day, test.part1, test.part2, p1, p2)
		}

	}
}
