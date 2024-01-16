package day1

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed part1_test.txt
var part1Text string

var numbersmap map[byte]int

func init() {
	numbersmap = map[byte]int{
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
}

func getFirstDigit(line string) int {
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			return int(line[i] - '0')
		}
	}
	// This should not happen
	panic(fmt.Sprintf("First digit was not found in %s", line))
}

func getLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if val, ok := numbersmap[line[i]]; ok {
			return val
		}
	}

	panic(fmt.Sprintf("Last digit was not found in %s", line))
}

func calculateTotal(text string) int {
	firstDigitSum := 0
	lastDigitSum := 0

	for _, line := range strings.Split(text, "\n") {
		firstDigitSum += getFirstDigit(line)
		lastDigitSum += getLastDigit(line)
	}

	return firstDigitSum*10 + lastDigitSum

}

func RunPart1() {
	fmt.Println("Result: ", calculateTotal(part1Text))
}
