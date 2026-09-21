package twenty_five

import (
	"adventOfCode/utils"
	"fmt"
	"strconv"
)

var dialPosition int = 50

var lZeroCounter int = 0

var filePath = "2025/files/dayOne.txt"

type Rotation struct {
	direction string
	numbers   int
}

func PartOne() int {
	var lines []string = utils.FileToLines(filePath)

	var rotations []Rotation = linesToRotations(lines)

	for _, r := range rotations {

		r.numbers = r.numbers % 100

		calculateRotation(r, r.numbers)

		if dialPosition == 0 {
			lZeroCounter++
		}
	}

	fmt.Println(lZeroCounter)

	return lZeroCounter
}

func PartTwo() int {
	var lines []string = utils.FileToLines(filePath)

	var rotations []Rotation = linesToRotations(lines)

	for _, r := range rotations {
		if r.numbers > 99 {
			lZeroCounter += r.numbers / 100
		}

		r.numbers = r.numbers % 100

		if calculateRotation(r, r.numbers) {
			lZeroCounter++
		}
	}

	fmt.Println(lZeroCounter)

	return lZeroCounter
}

func calculateRotation(r Rotation, actual int) bool {
	if r.direction == "R" {
		dialPosition += actual
		if dialPosition > 99 {
			dialPosition = dialPosition - 100
			return true
		}
	} else {
		dialPosition -= actual
		if dialPosition < 0 {
			dialPosition = 100 + dialPosition
			return true
		}
	}
	return false
}

func linesToRotations(lines []string) []Rotation {
	var rotations []Rotation

	for _, l := range lines {
		d := l[:1]
		n, _ := strconv.Atoi(l[1:])

		rotations = append(rotations, Rotation{direction: d, numbers: n})
	}

	return rotations
}
