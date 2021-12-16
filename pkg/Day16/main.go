package Day16

import (
	"encoding/hex"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/parser"
	"github.com/bznein/AoC2021/pkg/timing"
)

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	part1, part2 := 0, 0

	ints := input.InputToStringSlice(inputF)

	packet := ints[0]
	packetHex, _ := hex.DecodeString(packet)

	binaryVersion := HexToBin(packetHex)
	p := parser.NewParser(binaryVersion)
	part2 = p.ParsePacket()
	part1 = p.VersionSum()
	return part1, part2
}

func HexToBin(hexBytes []byte) []byte {
	res := make([]byte, 0, len(hexBytes)*4)
	for _, hexByte := range hexBytes {
		for i := 0; i < 8; i++ {
			res = append(res, hexByte&(1<<7)>>7+'0')
			hexByte <<= 1
		}
	}
	return res
}
