package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/tomadele/advent-of-code-2024/day-1"
	"github.com/tomadele/advent-of-code-2024/utils"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing day as argument")
	}
	dayFromArgs := os.Args[1]

	day, err := strconv.Atoi(dayFromArgs)
	if err != nil {
		log.Fatalf("strconv.Atoi(day): %v", err)
	}

	partFromArgs := os.Args[2]
	part, err := strconv.Atoi(partFromArgs)
	if err != nil {
		log.Fatalf("strconv.Atoi(part): %v", err)
	}

	input, err := ReadSpecificDayInput(day)
	if err != nil {
		log.Fatalf("ReadSpecificDayInput: %v", err)
	}

	switch day {
	case 1:
		if part == 1 {
			day_1.Part1(input)
		}
		if part == 2 {
			// day_1.Part2(input)
		}
	default:
		log.Fatalf("Day %q unsupported", day)
	}
}

func ReadSpecificDayInput(day int) (string, error) {
	inputPath, err := filepath.Abs(fmt.Sprintf("./day-%d/input.txt", day))
	if err != nil {
		return "", fmt.Errorf("filepath.Abs: %w", err)
	}
	input, err := utils.ReadFile(inputPath)
	if err != nil {
		return "", fmt.Errorf("utils.ReadFile: %w", err)
	}
	return input, nil
}
