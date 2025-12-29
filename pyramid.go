package pyramid

import "strings"

func TowerBuilder(nFloors int) []string {
	width := nFloors*2 - 1
	result := []string{}

	for i := 0; i < nFloors; i++ {
		stars := 2*i + 1
		spaces := (width - stars) / 2

		line := strings.Repeat(" ", spaces) +
			strings.Repeat("*", stars) +
			strings.Repeat(" ", spaces)

		result = append(result, line)
	}

	return result
}
