package Day17

import (
	"fmt"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/math"
	"github.com/bznein/AoC2021/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0
	ints := input.InputToStringSlice(inputF)

	var minX int
	var maxX int
	var minY int
	var maxY int
	fmt.Sscanf(ints[0], "target area: x=%d..%d, y=%d..%d", &minX, &maxX, &maxY, &minY)

	found := false
	for x := 0; x <= maxX; x++ {
		for y := maxY; y <= -maxY; y++ {
			if reaches, highest := getsToTarget(minX, maxX, maxY, minY, x, y); reaches {
				part2++
				if !found || part1 < highest {
					found = true
					part1 = highest
				}
			}
		}
	}

	return part1, part2
}

func getsToTarget(minX int, maxX int, minY int, maxY int, startX, startY int) (bool, int) {
	posX := 0
	posY := 0
	highset := posY
	for posX <= maxX && posY >= minY {
		posX += startX
		posY += startY
		startX = math.Max(0, startX-1)
		startY -= 1
		if posY > highset {
			highset = posY
		}
		if posX >= minX && posX <= maxX && posY <= maxY && posY >= minY {
			return true, highset
		}
	}
	return false, 0
}
