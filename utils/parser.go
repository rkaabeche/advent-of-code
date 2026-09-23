package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FileToLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func FileToString(filePath string) string {
	b, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}
