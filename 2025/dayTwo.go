package twenty_five

import (
	"adventOfCode/utils"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	min string
	max string
}

var filePath2 = "2025/files/dayTwo.txt"

var listOfNumb []int

func Day2PartOne() {
	var stringRanges string = utils.FileToString(filePath2)

	var ranges []Range = linesToRanges(strings.Split(stringRanges, ","))
	fmt.Println(ranges)

	for _, r := range ranges {
		listOfNumb = append(listOfNumb, allValidNumb(r)...)
	}
	print(utils.IntSliceAddition(listOfNumb))
}

func allValidNumb(r Range) []int {
	var validNum []int = []int{}
	min, _ := strconv.Atoi(r.min)
	max, _ := strconv.Atoi(r.max)

	if min > max {
		return validNum
	}
	fmt.Print(r)
	fmt.Print("  ----> ")

	for i := min; i <= max; i++ {
		strActual := strconv.Itoa(i)
		if len(strActual)%2 != 0 {
			continue
		}
		mid := len(strActual) / 2
		firsHalf := strActual[:mid]
		secondHalf := strActual[mid:]
		if firsHalf == secondHalf {
			validNum = append(validNum, i)
		}

	}
	fmt.Println(validNum)
	return validNum
}

func linesToRanges(lines []string) []Range {
	var ranges []Range
	for _, l := range lines {
		var splitLine = strings.Split(l, "-")
		ranges = append(ranges, Range{min: splitLine[0], max: splitLine[1]})
	}

	return ranges
}
