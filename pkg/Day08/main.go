package Day08

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2021/pkg/input"
	"github.com/bznein/AoC2021/pkg/timing"
)

func stringDiff(s1 string, s2 string) rune {
	for _, c := range s1 {
		if !strings.ContainsRune(s2, c) {
			return c
		}
	}
	return '~'
}

func findStringInSlice(strings []string, target string) int {
	for i, s := range strings {
		if s == target {
			return i
		}
	}
	return -1
}

func findStringByLenAndContent(stringsI []string, l int, subString string) string {

	for _, s := range stringsI {
		if len(s) != l {
			continue
		}
		found := true
		for _, c := range subString {
			if !strings.ContainsRune(s, c) {
				found = false
				break
			}
		}
		if found {
			return s
		}
	}

	return ""
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Dedup(input string) string {
	unique := ""

	for _, word := range input {
		// If we alredy have this word, skip.
		if strings.ContainsRune(unique, word) {
			continue
		}

		unique += string(word)
	}

	return unique
}

func solveSingleLine(line string) map[rune]rune {
	mapping := map[rune]rune{}
	splittedLine := strings.Split(line, " ")
	sortedLine := []string{}
	for _, s := range splittedLine {
		sortedLine = append(sortedLine, SortString(s))
	}

	sort.Slice(sortedLine, func(i, j int) bool {
		return len(sortedLine[i]) < len(sortedLine[j])
	})

	//one := sortedLine[0]
	seven := sortedLine[1]
	four := sortedLine[2]

	mapping[stringDiff(seven, four)] = 'a'
	firstFound := string(stringDiff(seven, four))
	fourSevenSum := SortString(Dedup(four + seven))

	next := findStringByLenAndContent(sortedLine[5:], 6, fourSevenSum)
	mapping[stringDiff(next, fourSevenSum)] = 'g'
	secondFound := string(stringDiff(next, fourSevenSum))

	nextTarget := SortString(Dedup(firstFound + secondFound + sortedLine[0]))
	next = findStringByLenAndContent(sortedLine[3:], 5, nextTarget)
	key := stringDiff(next, nextTarget)
	mapping[key] = 'd'
	thirdFound := string(key)
	nextTarget = sortedLine[0] + thirdFound
	key = stringDiff(sortedLine[2], nextTarget)
	mapping[key] = 'b'
	fourthFound := string(key)
	nextTarget = firstFound + secondFound + thirdFound + fourthFound
	next = findStringByLenAndContent(sortedLine[3:], 5, nextTarget)
	key = stringDiff(next, nextTarget)
	mapping[key] = 'f'
	fifhtFound := string(key)

	key = stringDiff(sortedLine[0], fifhtFound)
	mapping[key] = 'c'
	sixthFound := string(key)
	key = stringDiff(sortedLine[9], firstFound+secondFound+thirdFound+fourthFound+fifhtFound+sixthFound)
	mapping[key] = 'e'

	return mapping
}

func mapsTo(mapping map[rune]rune, s string, target rune) bool {

	for _, c := range s {
		if mapping[c] == target {
			return true
		}
	}

	return false
}

func decode(mapping map[rune]rune, s string) int {
	switch len(s) {
	case 7:
		return 8
	case 3:
		return 7
	case 4:
		return 4
	case 2:
		return 1
	case 5:
		if mapsTo(mapping, s, 'b') {
			return 5
		}
		if mapsTo(mapping, s, 'e') {
			return 2
		}
		return 3
	case 6:
		if mapsTo(mapping, s, 'd') {
			if mapsTo(mapping, s, 'e') {
				return 6
			}
			return 9

		}
		return 0

	default:
		panic("Wrong length!!")
	}

}

func Solve(inputF string) (int, int) {
	defer timing.TimeTrack(time.Now())
	inputS := input.InputToStringSlice(inputF)

	part1 := 0
	part2 := 0
	for _, s := range inputS {
		sSplit := strings.Split(s, " | ")
		digitValue := sSplit[1]
		mapping := solveSingleLine(sSplit[0])
		val := ""
		for _, v := range strings.Split(digitValue, " ") {
			switch len(v) {
			case 2, 3, 4, 7:
				part1++
			}
			val += strconv.Itoa(decode(mapping, v))
		}
		part2 += input.AsInt(val)
	}

	return part1, part2
}
