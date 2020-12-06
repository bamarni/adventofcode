package main

import "fmt"

var (
	pots  = []byte(".##..#.#..##..##..##...#####.#.....#..#..##.###.#.####......#.......#..###.#.#.##.#.#.###...##.###.#")
	notes = map[string]byte{
		".##.#": '#',
		"##.#.": '#',
		"##...": '#',
		"#....": '.',
		".#..#": '.',
		"#.##.": '.',
		".##..": '.',
		".#.##": '.',
		"###..": '.',
		"..##.": '#',
		"#####": '#',
		"#...#": '#',
		".#...": '#',
		"###.#": '#',
		"#.###": '#',
		"##..#": '.',
		".###.": '#',
		"...##": '.',
		"..#.#": '.',
		"##.##": '#',
		"....#": '.',
		"#.#.#": '#',
		"#.#..": '.',
		".####": '.',
		"...#.": '#',
		"..###": '.',
		"..#..": '#',
		".....": '.',
		"####.": '.',
		"#..##": '#',
		".#.#.": '.',
		"#..#.": '#',
	}
	sum, zero int
)

func main() {
	for gen := 0; gen < 20; gen++ {
		for ; string(pots[:5]) != "....."; zero++ {
			pots = append([]byte("."), pots...)
		}
		for string(pots[len(pots)-5:]) != "....." {
			pots = append(pots, byte('.'))
		}

		transformations := make(map[int]byte)
		for i := 2; i < len(pots)-3; i++ {
			transformations[i] = notes[string(pots[i-2:i+3])]
		}

		for i := range transformations {
			pots[i] = transformations[i]
		}
	}

	for i := 0; i < len(pots); i++ {
		if pots[i] == '#' {
			sum += i - zero
		}
	}
	fmt.Println(sum)
}
